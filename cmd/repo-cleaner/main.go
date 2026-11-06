// cmd/repo-cleaner/main.go v2.0.0

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
	"repo-cleaner/internal/platform"
)

func main() {
	langFlag := flag.String("lang", "", "界面语言 en|zh（写入 config.ini）")
	showHelp := flag.Bool("h", false, "显示帮助")
	flag.BoolVar(showHelp, "help", false, "显示帮助")
	flag.Parse()

	cfgPath := config.Path()
	cfg, err := config.Load(cfgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取配置失败: %v\n", err)
	}
	p := platform.New()

	if *showHelp {
		printHelp()
		return
	}
	if *langFlag != "" {
		cfg.Lang = string(i18n.Normalize(*langFlag))
		if err := config.Save(cfgPath, cfg); err != nil {
			fmt.Fprintf(os.Stderr, "保存配置失败: %v\n", err)
		}
	}

	fn, pathArg := parseArgs(flag.Args())
	root := resolveRoot(p, cfg, pathArg)

	if fn > 0 {
		runCLI(p, cfgPath, &cfg, fn, root)
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

// runCLI 命令行模式：危险操作保留二次确认（对齐 OPENSPEC §4.5）。
func runCLI(p platform.Platform, cfgPath string, cfg *config.Config, fn int, root string) {
	if err := platform.ValidateRoot(root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if engine.IsDangerous(fn) {
		fmt.Print("是否继续? [Y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		if !strings.EqualFold(strings.TrimSpace(answer), "Y") {
			fmt.Println("已取消")
			return
		}
	}
	if err := engine.Run(p, fn, root, engine.StdLogger{}); err != nil {
		fmt.Fprintf(os.Stderr, "[错误] %v\n", err)
		os.Exit(1)
	}
	// 命令行模式记住本次使用的目录
	cfg.RepoRoot = root
	_ = config.Save(cfgPath, *cfg)
}

// printHelp 输出帮助信息（对齐 v1.x 行为）。
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
	fmt.Println("  -h        Show this help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  repo-cleaner                    # GUI mode")
	fmt.Println("  repo-cleaner 1                  # CLI: enable global node_modules")
	fmt.Println("  repo-cleaner 5 E:\\Projects      # CLI with path")
}
