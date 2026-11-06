//go:build nogui

// cmd/repo-cleaner/headless.go v2.0.0

package main

import (
	"fmt"

	"repo-cleaner/internal/config"
)

// runGUI 在 nogui 构建下提示使用命令行模式（用于无 C 工具链环境的校验构建）。
func runGUI(cfg config.Config) {
	fmt.Println("本构建未包含 GUI（构建标签 nogui）。")
	fmt.Println("请使用命令行模式：repo-cleaner <1-7> [目录路径]，或 repo-cleaner -h 查看帮助。")
}
