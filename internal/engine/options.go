// internal/engine/options.go v2.1.0

package engine

import (
	"fmt"

	"repo-cleaner/internal/i18n"
	"repo-cleaner/internal/platform"
)

// Options 功能执行选项。
type Options struct {
	// DryRun 预演模式：仅输出将要执行的操作，不修改任何文件与环境变量
	// （对齐 OPENSPEC §4.7）；危险操作在预演下跳过二次确认。
	DryRun bool
}

// Run 按功能编号（1-7）执行对应功能，lang 决定执行日志语言。
func Run(p platform.Platform, fn int, repoRoot string, lang i18n.Lang, opts Options, log Logger) error {
	if opts.DryRun {
		log.Logf("%s", i18n.Logs(lang).DryRunOn)
	}
	switch fn {
	case 1:
		return EnableGlobalNodeModules(p, repoRoot, lang, opts, log)
	case 2:
		return ResetNodeConfig(p, repoRoot, lang, opts, log)
	case 3:
		return CleanNextCache(p, repoRoot, lang, opts, log)
	case 4:
		return RestoreNextCache(p, repoRoot, lang, opts, log)
	case 5:
		return InstallGlobalDeps(lang, opts, log)
	case 6:
		return CleanAllCaches(p, repoRoot, lang, opts, log)
	case 7:
		return CleanViteCache(p, repoRoot, lang, opts, log)
	default:
		return fmt.Errorf("unknown function: %d", fn)
	}
}

// IsDangerous 返回功能是否为危险操作（需要二次确认）：1 / 6 / 7。
func IsDangerous(fn int) bool { return fn == 1 || fn == 6 || fn == 7 }
