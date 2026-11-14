//go:build !nogui

// internal/ui/ui.go v2.3.3

// Package ui 实现 Fyne 桌面 GUI（对齐 docs/DESIGN-SYSTEM.md §3.0）。
// 本包需要 C 工具链编译（Fyne 依赖 OpenGL）；无 C 工具链环境用 `go build -tags nogui` 校验。
package ui

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"repo-manager/internal/config"
	"repo-manager/internal/engine"
	"repo-manager/internal/i18n"
	"repo-manager/internal/logs"
	"repo-manager/internal/platform"
)

// langOptions 语言切换顺序：简中 → 繁中 → 英文（覆盖全部支持语言）。
var langOptions = []struct {
	label string
	lang  i18n.Lang
}{
	{"简体中文", i18n.Zh},
	{"繁體中文", i18n.Hant},
	{"English", i18n.En},
}

// Run 启动 Fyne GUI（阻塞直到窗口关闭）。
func Run(cfg config.Config) {
	a := app.NewWithID("repo-manager")
	a.Settings().SetTheme(newTheme())
	w := a.NewWindow("Repo Manager")
	w.Resize(fyne.NewSize(760, 600))

	p := platform.New()
	cfgPath := config.Path()
	if cfg.RepoRoot == "" {
		cfg.RepoRoot = p.DefaultRepoRoot()
	}

	// render 按当前语言重建窗口内容（语言切换即整体重绘）。
	var render func()
	render = func() {
		lang := i18n.Normalize(cfg.Lang)
		dict := i18n.T(lang)
		d := i18n.Logs(lang)
		w.SetTitle(dict.AppTitle)

		output := widget.NewMultiLineEntry()
		output.Wrapping = fyne.TextWrapWord
		output.Disable()
		progress := newProgressView()
		logger := &guiLogger{out: output, mu: &sync.Mutex{}, progress: progress}

		dirEntry := widget.NewEntry()
		dirEntry.SetText(cfg.RepoRoot)

		labels := make([]string, len(langOptions))
		for i, option := range langOptions {
			labels[i] = option.label
		}
		langSelect := widget.NewSelect(labels, func(s string) {
			for _, option := range langOptions {
				if option.label != s {
					continue
				}
				if option.lang == i18n.Normalize(cfg.Lang) {
					return
				}
				cfg.Lang = string(option.lang)
				if err := config.Save(cfgPath, cfg); err != nil {
					dialog.ShowError(err, w)
					return
				}
				render()
				return
			}
		})
		for _, option := range langOptions {
			if option.lang == lang {
				langSelect.SetSelected(option.label)
				break
			}
		}
		langSelect.PlaceHolder = dict.LangLabel

		browse := widget.NewButton(dict.ChooseDir, func() {
			dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
				if err != nil || uri == nil {
					return
				}
				dirEntry.SetText(uri.Path())
			}, w)
		})

		saveBtn := widget.NewButton(dict.Save, func() {
			root := platform.NormalizeRoot(dirEntry.Text)
			if err := platform.ValidateRoot(root); err != nil {
				dialog.ShowError(err, w)
				return
			}
			cfg.RepoRoot = root
			if err := config.Save(cfgPath, cfg); err != nil {
				dialog.ShowError(err, w)
				return
			}
			dialog.ShowInformation(dict.Ok, dict.Saved, w)
		})

		aboutBtn := widget.NewButton(dict.About, func() {
			dialog.ShowInformation(dict.About, dict.AboutText, w)
		})

		// 预演模式复选框：会话级选择，不写入 config.ini（对齐 OPENSPEC §4.7）。
		dryRunCheck := widget.NewCheck(dict.DryRunLabel, nil)
		dryRunCheck.SetChecked(false)

		// 递归扫描复选框：会话级选择（V1 monorepo，深度 3 层）。
		recursiveCheck := widget.NewCheck(dict.RecursiveLabel, nil)
		recursiveCheck.SetChecked(cfg.Recursive)

		// runFunction 执行功能；危险操作先弹二次确认，预演模式跳过确认（对齐 OPENSPEC §4.5 / §4.7）。
		runFunction := func(fn int) {
			root := platform.NormalizeRoot(dirEntry.Text)
			if err := platform.ValidateRoot(root); err != nil {
				dialog.ShowError(err, w)
				return
			}
			cfg.RepoRoot = root
			opts := engine.Options{
				DryRun:         dryRunCheck.Checked,
				Recursive:      recursiveCheck.Checked,
				ExtraCleanDirs: cfg.ExtraCleanDirs,
			}
			do := func(confirmed bool) {
				if !confirmed {
					return
				}
				progress.Reset()
				go func() {
					defer progress.Finish()
					// V4：执行日志同时追加写入 repo-manager.log。
					fileLog := logs.NewFileLogger(logs.DefaultPath(), logger)
					defer fileLog.Close()
					if err := engine.Run(p, fn, root, lang, opts, fileLog); err != nil {
						fileLog.Logf(d.ErrorPrefix, err)
					}
				}()
			}
			if engine.IsDangerous(fn) && !opts.DryRun {
				warning := dict.WarningF1
				if fn == 6 {
					warning = dict.WarningAll
				} else if fn == 7 {
					warning = dict.WarningVite
				}
				dialog.ShowConfirm(warning, dict.Confirm, do, w)
				return
			}
			do(true)
		}

		fnButtons := container.NewGridWithColumns(2)
		for i, name := range dict.FuncNames {
			fn := i + 1
			fnButtons.Add(widget.NewButton(fmt.Sprintf("[%d] %s", fn, name), func() { runFunction(fn) }))
		}

		top := container.NewVBox(
			container.NewHBox(widget.NewLabel(dict.TargetDir), dirEntry, browse, saveBtn),
			container.NewHBox(widget.NewLabel(dict.LangLabel), langSelect, aboutBtn),
			container.NewHBox(widget.NewLabel(dict.RunHint), dryRunCheck, recursiveCheck),
			// V5：进度条与当前步骤文案，按执行日志的 [n/m] 标记推进
			container.NewVBox(progress.label, progress.bar),
		)
		// 页脚区（Border bottom）：版本号 + build 编号的唯一展示位置，页眉不放版本号
		w.SetContent(container.NewBorder(top,
			widget.NewLabel(footerText()), nil, nil,
			container.NewVScroll(container.NewVBox(fnButtons,
				widget.NewLabel(dict.OutputTitle), output))),
		)
	}
	render()
	w.ShowAndRun()
}
