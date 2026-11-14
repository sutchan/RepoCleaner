//go:build nogui

// cmd/repo-manager/headless.go v2.0.2

package main

import (
	"fmt"

	"repo-manager/internal/config"
	"repo-manager/internal/i18n"
)

// runGUI 在 nogui 构建下提示使用命令行模式（用于无 C 工具链环境的校验构建）。
func runGUI(cfg config.Config) {
	d := i18n.Logs(i18n.Normalize(cfg.Lang))
	fmt.Println(d.NoGUI[0])
	fmt.Println(d.NoGUI[1])
}
