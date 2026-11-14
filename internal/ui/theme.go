//go:build !nogui

// internal/ui/theme.go v2.3.2

package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"repo-manager/internal/buildinfo"
)

// 设计令牌（对齐 docs/DESIGN-SYSTEM.md §2.1 暗色主题）。
var (
	colorBackground  = color.NRGBA{R: 0x0a, G: 0x0a, B: 0x0b, A: 0xff}
	colorForeground  = color.NRGBA{R: 0xed, G: 0xed, B: 0xed, A: 0xff}
	colorPrimary     = color.NRGBA{R: 0x10, G: 0xb9, B: 0x81, A: 0xff}
	colorDestructive = color.NRGBA{R: 0xef, G: 0x44, B: 0x44, A: 0xff}
	colorMuted       = color.NRGBA{R: 0x1c, G: 0x1c, B: 0x1f, A: 0xff}
	colorBorder      = color.NRGBA{R: 0x27, G: 0x27, B: 0x2a, A: 0xff}
)

// darkTheme 基于 Fyne 默认暗色主题覆盖品牌配色（对齐 DESIGN-SYSTEM §2.1）。
type darkTheme struct{ fyne.Theme }

// newTheme 构建应用主题。
func newTheme() fyne.Theme { return &darkTheme{Theme: theme.DefaultTheme()} }

// Color 覆盖品牌令牌，其余回落 Fyne 默认。
// Fyne v2.8.2 的 Theme 接口为 Color(ThemeColorName, ThemeVariant) color.Color。
func (t *darkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return colorBackground
	case theme.ColorNameForeground:
		return colorForeground
	case theme.ColorNamePrimary:
		return colorPrimary
	case theme.ColorNameFocus:
		return colorPrimary
	case theme.ColorNameError:
		return colorDestructive
	case theme.ColorNameInputBorder:
		return colorBorder
	case theme.ColorNameSeparator:
		return colorBorder
	default:
		return t.Theme.Color(name, variant)
	}
}

// footerText 返回应用页脚展示的版本与 build 编号（格式见 CONTRIBUTING.md §8.2）。
func footerText() string { return buildinfo.Footer() }
