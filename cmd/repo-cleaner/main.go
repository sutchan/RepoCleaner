// cmd/repo-cleaner/main.go v2.3.2

// 命令 repo-cleaner 是产品入口：解析参数后分流到命令行模式或 Fyne GUI。
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"repo-cleaner/internal/buildinfo"
	"repo-cleaner/internal/config"
	"repo-cleaner/internal/engine"
	"repo-cleaner/internal/i18n"
	"repo-cleaner/internal/logs"
	"repo-cleaner/internal/platform"
)

func main() {
	langFlag := flag.String("lang", "", "界面语言 en|zh（写入 config.ini）")
	dryRunFlag := flag.Bool("dry-run", false, "预演模式：仅显示将要执行的操作，不修改任何文件")
	recursiveFlag := flag.Bool("recursive", false, "递归扫描 monorepo 深层项目（深度 3 层）")
	showHelp := flag.Bool("h", false, "显示帮助")
	flag.BoolVar(showHelp, "help", false, "显示帮助")
	flag.Parse()

	cfgPath := config.Path()
	cfg, err := config.Load(cfgPath)
	if *showHelp {
		printHelp()
		return
	}
	if *langFlag != "" {
		cfg.Lang = string(i18n.Normalize(*langFlag))
		if saveErr := config.Save(cfgPath, cfg); saveErr != nil {
			fmt.Fprintf(os.Stderr, i18n.Logs(i18n.Normalize(cfg.Lang)).SaveCfgErr, saveErr)
		}
	}
	// 选项优先级：命令行参数 > config.ini > 默认关闭（对齐 OPENSPEC §4.5 / §5.2）。
	opts := engine.Options{
		DryRun:         *dryRunFlag,
		Recursive:      *recursiveFlag || cfg.Recursive,
		ExtraCleanDirs: cfg.ExtraCleanDirs,
	}

	lang := i18n.Normalize(cfg.Lang)
	d := i18n.Logs(lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, d.LoadCfgErr, err)
	}
	p := platform.New()
	fn, pathArg := parseArgs(flag.Args())
	root := resolveRoot(p, cfg, pathArg)

	if fn > 0 {
		runCLI(p, cfgPath, &cfg, fn, root, lang, opts)
		return
	}
	runGUI(cfg)
}

// parseArgs 解析位置参数：[功能号 1-7] [路径]。
func parseArgs(args []string) (fn int, path string) {
	if len(args) == 0 {
		return 0, ""
	}
	if n, err := strconv.Atoi(args[0]); err == nil && n >= 1 && n <= 7 {
		fn = n
		args = args[1:]
	}
	if len(args) > 0 {
		path = args[0]
	}
	return fn, path
}

// resolveRoot 确定仓库根目录：命令行路径 → config.ini → 平台默认值（对齐 OPENSPEC §4.2）。
func resolveRoot(p platform.Platform, cfg config.Config, pathArg string) string {
	if pathArg != "" {
		return platform.NormalizeRoot(pathArg)
	}
	if cfg.RepoRoot != "" {
		return platform.NormalizeRoot(cfg.RepoRoot)
	}
	return p.DefaultRepoRoot()
}

// runCLI 命令行模式：危险操作保留二次确认，预演模式跳过（对齐 OPENSPEC §4.5 / §4.7）。
func runCLI(p platform.Platform, cfgPath string, cfg *config.Config, fn int, root string, lang i18n.Lang, opts engine.Options) {
	d := i18n.Logs(lang)
	if err := platform.ValidateRoot(root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if engine.IsDangerous(fn) && !opts.DryRun {
		fmt.Print(d.ConfirmCLI)
		reader := bufio.NewReader(os.Stdin)
		answer, readErr := reader.ReadString('\n')
		if !strings.EqualFold(strings.TrimSpace(answer), "Y") {
			// 非交互环境（stdin 关闭 / EOF）不能静默当作「用户取消」，否则脚本会误判成功。
			if readErr != nil {
				fmt.Fprintln(os.Stderr, d.ConfirmUnavailable)
				os.Exit(1)
			}
			fmt.Println(d.Cancelled)
			return
		}
	}
	// V4：执行日志同时追加写入可执行文件同目录的 repo-cleaner.log。
	fileLog := logs.NewFileLogger(logs.DefaultPath(), engine.StdLogger{})
	defer fileLog.Close()
	if err := engine.Run(p, fn, root, lang, opts, fileLog); err != nil {
		fmt.Fprintf(os.Stderr, d.ErrorPrefix+"\n", err)
		os.Exit(1)
	}
	// 命令行模式记住本次使用的目录
	cfg.RepoRoot = root
	if err := config.Save(cfgPath, *cfg); err != nil {
		fmt.Fprintf(os.Stderr, d.SaveCfgErr, err)
	}
}

// printHelp 输出帮助信息；CLI 帮助固定英文，界面文案取自 internal/i18n。
func printHelp() {
	fmt.Println("RepoCleaner - Project Cleaner")
	fmt.Println()
	fmt.Printf("Version: %s\n\n", buildinfo.BuildInfo())
	fmt.Println("Usage: repo-cleaner [options] [function] [path]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  1-7       Run corresponding function directly (CLI mode)")
	fmt.Println("  -lang en  Set language to English")
	fmt.Println("  -lang zh  Set language to Chinese")
	fmt.Println("  -dry-run  Preview actions without modifying anything")
	fmt.Println("  -recursive  Recursive scan (monorepo, depth 3)")
	fmt.Println("  -h        Show this help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  repo-cleaner                    # GUI mode")
	fmt.Println("  repo-cleaner 1                  # CLI: enable global node_modules")
	fmt.Println("  repo-cleaner 5 E:\\Projects      # CLI with path")
}
