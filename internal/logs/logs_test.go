// internal/logs/logs_test.go v2.3.3

package logs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFileLogger V4：日志应同时落盘（带时间戳）并转发下游。
func TestFileLogger(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "repo-manager.log")
	var forwarded int
	l := NewFileLogger(path, &countLogger{&forwarded})
	l.Logf("step %d", 1)
	l.Logf("clean %s", "proj-a")
	l.Close()
	l.Logf("after close") // 关闭后不写文件，仍应转发

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if strings.Count(content, "\n") != 2 {
		t.Fatalf("文件应恰有 2 条日志，得到 %q", content)
	}
	if !strings.Contains(content, "step 1") || !strings.Contains(content, "clean proj-a") {
		t.Fatalf("文件应包含日志内容，得到 %q", content)
	}
	if strings.Contains(content, "after close") {
		t.Fatal("关闭后不应再写文件")
	}
	if forwarded != 3 {
		t.Fatalf("下游应收到全部 3 条日志，得到 %d", forwarded)
	}
}

// TestNewFileLoggerInvalidPath 打开失败时应降级为仅下游输出，不 panic。
func TestNewFileLoggerInvalidPath(t *testing.T) {
	n := 0
	l := NewFileLogger(filepath.Join(t.TempDir(), "no-such-dir", "x.log"), &countLogger{&n})
	l.Logf("hello")
	l.Close()
	if n != 1 {
		t.Fatalf("降级时下游仍应收到日志，得到 %d", n)
	}
}

type countLogger struct{ n *int }

func (c *countLogger) Logf(format string, args ...any) { *c.n++ }
