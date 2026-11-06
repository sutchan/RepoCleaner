//go:build !nogui

// cmd/repo-cleaner/gui.go v2.0.0

package main

import (
	"repo-cleaner/internal/config"
	"repo-cleaner/internal/ui"
)

// runGUI 启动 Fyne 桌面 GUI（默认构建包含 GUI）。
func runGUI(cfg config.Config) { ui.Run(cfg) }
