// internal/engine/dryrun_test.go v2.1.0

package engine

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"repo-manager/internal/i18n"
)

// TestDryRunCleanAllCaches 校验预演模式不删除任何文件且输出预演日志。
func TestDryRunCleanAllCaches(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "proj-a", map[string][]byte{
		"package.json":      []byte("{}"),
		"node_modules/l.js": []byte("x"),
		"dist/i.js":         []byte("x"),
		"package-lock.json": []byte("{}"),
	})
	var log captureLogger
	if err := Run(s, 6, s.DefaultRepoRoot(), i18n.Zh, Options{DryRun: true}, &log); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"node_modules", "dist", "package-lock.json"} {
		if !exists(t, filepath.Join(proj, filepath.FromSlash(name))) {
			t.Fatalf("dry-run 不得删除 %s", name)
		}
	}
	joined := strings.Join(log.lines, "\n")
	if !strings.Contains(joined, "[预演] 将删除:") || !strings.Contains(joined, "node_modules") {
		t.Fatalf("应输出删除预演日志，实际:\n%s", joined)
	}
}

// TestDryRunEnableGlobalNodeModules 校验功能 1 预演不写环境变量 / .npmrc / 删目录。
func TestDryRunEnableGlobalNodeModules(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "proj-a", map[string][]byte{
		"package.json":        []byte("{}"),
		"node_modules/lib.js": []byte("x"),
	})
	var log captureLogger
	if err := EnableGlobalNodeModules(s, s.DefaultRepoRoot(), i18n.Zh, Options{DryRun: true}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(proj, ".npmrc")) {
		t.Fatal("dry-run 不得写入 .npmrc")
	}
	if !exists(t, filepath.Join(proj, "node_modules")) {
		t.Fatal("dry-run 不得删除 node_modules")
	}
	if _, ok := s.envs["NODE_PATH"]; ok {
		t.Fatal("dry-run 不得设置 NODE_PATH")
	}
	joined := strings.Join(log.lines, "\n")
	for _, want := range []string{"[预演] 将写入:", "[预演] 将删除:", "[预演] 将设置环境变量:"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("预演日志缺少 %q，实际:\n%s", want, joined)
		}
	}
}

// TestDryRunInstallGlobalDeps 校验功能 5 预演只列包不执行 npm。
func TestDryRunInstallGlobalDeps(t *testing.T) {
	called := false
	saved := npmRunner
	npmRunner = func([]string) (string, error) {
		called = true
		return "", nil
	}
	t.Cleanup(func() { npmRunner = saved })

	var log captureLogger
	if err := InstallGlobalDeps(i18n.Zh, Options{DryRun: true}, &log); err != nil {
		t.Fatal(err)
	}
	if called {
		t.Fatal("dry-run 不得执行 npm install")
	}
	joined := strings.Join(log.lines, "\n")
	if !strings.Contains(joined, "[预演] 将安装: react react-dom next") {
		t.Fatalf("应输出安装预演日志，实际:\n%s", joined)
	}
}

// TestInstallGlobalDepsSummary 校验功能 5 汇总失败组并返回错误（L4 修复）。
func TestInstallGlobalDepsSummary(t *testing.T) {
	saved := npmRunner
	npmRunner = func([]string) (string, error) { return "", errors.New("network down") }
	t.Cleanup(func() { npmRunner = saved })

	var log captureLogger
	err := InstallGlobalDeps(i18n.Zh, Options{}, &log)
	if err == nil {
		t.Fatal("存在失败组时应返回错误")
	}
	joined := strings.Join(log.lines, "\n")
	if !strings.Contains(joined, "以下分组安装失败: 1, 2, 3, 4, 5") {
		t.Fatalf("应汇总全部失败组，实际:\n%s", joined)
	}
	if !strings.Contains(err.Error(), "以下分组安装失败") {
		t.Fatalf("错误信息应包含失败组汇总，实际: %v", err)
	}
}

// TestInstallGlobalDepsAllSuccess 校验全部成功时返回 nil 且无汇总。
func TestInstallGlobalDepsAllSuccess(t *testing.T) {
	saved := npmRunner
	npmRunner = func([]string) (string, error) { return "", nil }
	t.Cleanup(func() { npmRunner = saved })

	var log captureLogger
	if err := InstallGlobalDeps(i18n.Zh, Options{}, &log); err != nil {
		t.Fatalf("全部安装成功不应返回错误: %v", err)
	}
	if strings.Contains(strings.Join(log.lines, "\n"), "以下分组安装失败") {
		t.Fatal("全部成功时不应输出失败汇总")
	}
}
