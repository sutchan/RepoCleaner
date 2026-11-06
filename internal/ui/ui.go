//go:build !nogui

// internal/ui/ui.go v2.0.1

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

	"repo-cleaner/internal/config"
	"repo-cleaner/internal/engine"
	"repo-cleaner/internal/i18n"
	"repo-cleaner/internal/platform"
)

// Run 启动 Fyne GUI（阻塞直到窗口关闭）。
func Run(cfg config.Config) {
	a := app.NewWithID("repo-cleaner")
	a.Settings().SetTheme(newTheme())
	w := a.NewWindow("RepoCleaner")
	w.Resize(fyne.NewSize(760, 600))

	p := platform.New()
	cfgPath := config.Path()
	if cfg.RepoRoot == "" {
		cfg.RepoRoot = p.DefaultRepoRoot()
	}

	// render 按当前语言重建窗口内容（语言切换即整体重绘）。
	var render func()
	render = func() {
		dict := i18n.T(i18n.Lang(cfg.Lang))
		w.SetTitle(dict.AppTitle)

		output := widget.NewMultiLineEntry()
		output.Wrapping = fyne.TextWrapWord
		output.Disable()
		logger := &guiLogger{out: output, mu: &sync.Mutex{}}

		dirEntry := widget.NewEntry()
		dirEntry.SetText(cfg.RepoRoot)

		langSelect := widget.NewSelect([]string{"中文", "English"}, func(s string) {
			lang := i18n.Zh
			if s == "English" {
				lang = i18n.En
			}
			if lang == i18n.Lang(cfg.Lang) {
				return
			}
			cfg.Lang = string(lang)
			_ = config.Save(cfgPath, cfg)
			render()
		})
		langSelect.PlaceHolder(dict.LangLabel)

		browse := widget.NewButton(dict.ChooseDir, func() {
			dialog.ShowFolderOpen(func(uri fyne.URIReadCloser, err error) {
				if err != nil || uri == nil {
					return
				}
				dirEntry.SetText(uri.URI().Path())
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

		// runFunction 执行功能；危险操作先弹二次确认（对齐 OPENSPEC §4.5）。
		runFunction := func(fn int) {
			root := platform.NormalizeRoot(dirEntry.Text)
			if err := platform.ValidateRoot(root); err != nil {
				dialog.ShowError(err, w)
				return
			}
			cfg.RepoRoot = root
			do := func(confirmed bool) {
				if !confirmed {
					return
				}
				go func() {
					if err := engine.Run(p, fn, root, logger); err != nil {
						logger.Logf("[错误] %v", err)
					}
				}()
			}
			if engine.IsDangerous(fn) {
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
			container.NewHBox(widget.NewLabel(dict.LangLabel), langSelect),
			widget.NewLabel(dict.RunHint),
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

// guiLogger 把 engine 日志安全地投递回 UI 线程。
type guiLogger struct {
	out *widget.Entry
	mu  *sync.Mutex
}

// Logf 实现 engine.Logger。
func (g *guiLogger) Logf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	g.mu.Lock()
	defer g.mu.Unlock()
	fyne.Do(func() {
		g.out.SetText(g.out.Text + msg + "\n")
	})
}
