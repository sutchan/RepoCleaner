//go:build !nogui

// internal/ui/output.go v2.3.3

package ui

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// progressView 把执行日志的 [n/m] 分步标记映射为进度条与当前步骤文案。
type progressView struct {
	bar   *widget.ProgressBar
	label *widget.Label
}

func newProgressView() *progressView {
	return &progressView{
		bar:   widget.NewProgressBar(),
		label: widget.NewLabel(""),
	}
}

// Reset 在每次功能执行前清空进度。
func (p *progressView) Reset() {
	fyne.Do(func() {
		p.bar.SetValue(0)
		p.label.SetText("")
	})
}

// Observe 读取一条日志，命中 [n/m] 时更新进度与步骤文案。
func (p *progressView) Observe(line string) {
	done, total, ok := parseStep(line)
	if !ok {
		return
	}
	fyne.Do(func() {
		p.bar.SetValue(float64(done) / float64(total))
		p.label.SetText(line)
	})
}

// Finish 在功能结束后把进度收口为 100%。
func (p *progressView) Finish() {
	fyne.Do(func() { p.bar.SetValue(1) })
}

// guiLogger 把 engine 日志安全地投递回 UI 线程，并同步进度视图。
type guiLogger struct {
	out      *widget.Entry
	mu       *sync.Mutex
	progress *progressView
}

// Logf 实现 engine.Logger。
func (g *guiLogger) Logf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	g.mu.Lock()
	defer g.mu.Unlock()
	fyne.Do(func() {
		g.out.SetText(g.out.Text + msg + "\n")
	})
	if g.progress != nil {
		g.progress.Observe(msg)
	}
}
