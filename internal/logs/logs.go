// internal/logs/logs.go v2.3.3

// Package logs 提供执行日志的文件落盘能力（V4）：在下游 Logger 输出之上
// 追加写入本地日志文件；文件打开 / 写入失败时仅保留下游输出，不中断功能。
package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"repo-manager/internal/engine"
)

// DefaultPath 返回与可执行文件同目录的日志文件路径（repo-manager.log）。
func DefaultPath() string {
	exe, err := os.Executable()
	if err != nil {
		return "repo-manager.log"
	}
	return filepath.Join(filepath.Dir(exe), "repo-manager.log")
}

// FileLogger 将日志同时写入文件与下游 Logger（engine.Logger）。
// 每次 Logf 带时间戳前缀；线程安全。
type FileLogger struct {
	mu   sync.Mutex
	f    *os.File
	next engine.Logger
}

// NewFileLogger 以追加模式打开日志文件；打开失败时返回仅含下游输出的 logger。
func NewFileLogger(path string, next engine.Logger) *FileLogger {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		f = nil
	}
	return &FileLogger{f: f, next: next}
}

// Logf 实现 engine.Logger：写文件（带时间戳）并转发下游。
func (l *FileLogger) Logf(format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.f != nil {
		fmt.Fprintf(l.f, "%s %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
	if l.next != nil {
		l.next.Logf(format, args...)
	}
}

// Close 关闭日志文件（幂等）。
func (l *FileLogger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.f != nil {
		_ = l.f.Close()
		l.f = nil
	}
}
