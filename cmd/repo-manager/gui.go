//go:build !nogui

// cmd/repo-manager/gui.go v2.0.0

package main

import (
	"repo-manager/internal/config"
	"repo-manager/internal/ui"
)

// runGUI 启动 Fyne 桌面 GUI（默认构建包含 GUI）。
func runGUI(cfg config.Config) { ui.Run(cfg) }
