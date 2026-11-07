// internal/engine/testkit_test.go v2.0.2

package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"repo-cleaner/internal/platform"
)

// captureLogger 收集日志便于断言。
type captureLogger struct{ lines []string }

func (c *captureLogger) Logf(format string, args ...any) {
	c.lines = append(c.lines, strings.TrimSpace(fmt.Sprintf(format, args...)))
}

// stubPlatform 委托真实平台实现（目录枚举 / 删除 / 隐藏属性），
// 仅替换外部副作用（npm 查询、环境变量持久化、默认根目录），
// 确保删除类测试只作用于 t.TempDir()，绝不触碰真实仓库根目录。
type stubPlatform struct {
	platform.Platform
	root   string
	global string
	envs   map[string]string
}

func newStub(t *testing.T) *stubPlatform {
	t.Helper()
	return &stubPlatform{
		Platform: platform.New(),
		root:     t.TempDir(),
		global:   t.TempDir(),
		envs:     map[string]string{},
	}
}

func (s *stubPlatform) DefaultRepoRoot() string                { return s.root }
func (s *stubPlatform) GlobalNodeModulesPath() (string, error) { return s.global, nil }
func (s *stubPlatform) SetUserEnv(k, v string) error           { s.envs[k] = v; return nil }
func (s *stubPlatform) RemoveUserEnv(k string) error           { delete(s.envs, k); return nil }

// newProject 在临时仓库根下创建一个假项目。
func newProject(t *testing.T, root, name string, files map[string][]byte) string {
	t.Helper()
	dir := filepath.Join(root, name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	for rel, content := range files {
		full := filepath.Join(dir, rel)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, content, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func exists(t *testing.T, path string) bool {
	t.Helper()
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	t.Fatal(err)
	return false
}
