// internal/engine/engine_test.go v2.1.0

package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"repo-cleaner/internal/i18n"
)

func TestCleanAllCaches(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "proj-a", map[string][]byte{
		"package.json":        []byte("{}"),
		"node_modules/lib.js": []byte("x"),
		"dist/index.js":       []byte("x"),
		"package-lock.json":   []byte("{}"),
		"src/main.js":         []byte("keep"),
	})
	var log captureLogger
	if err := CleanAllCaches(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(proj, "node_modules")) || exists(t, filepath.Join(proj, "dist")) {
		t.Fatal("白名单目录应被删除")
	}
	if exists(t, filepath.Join(proj, "package-lock.json")) {
		t.Fatal("锁文件应被删除")
	}
	if !exists(t, filepath.Join(proj, "src", "main.js")) {
		t.Fatal("白名单之外的文件必须保留")
	}
}

func TestCleanViteCache(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "vite-app", map[string][]byte{
		"package.json":        []byte("{}"),
		".vite/deps.json":     []byte("x"),
		"node_modules/lib.js": []byte("x"),
	})
	var log captureLogger
	if err := CleanViteCache(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(proj, ".vite")) {
		t.Fatal(".vite 应被删除")
	}
	if !exists(t, filepath.Join(proj, "node_modules")) {
		t.Fatal("功能 7 不应删除 node_modules")
	}
}

func TestCleanThenRestoreNextCache(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "next-app", map[string][]byte{
		"package.json":   []byte("{}"),
		"next.config.ts": []byte("export default {};"),
	})
	var log captureLogger

	if err := CleanNextCache(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	next := filepath.Join(proj, ".next")
	info, err := os.Stat(next)
	if err != nil {
		t.Fatalf("功能 3 后应存在 .next: %v", err)
	}
	if info.IsDir() {
		t.Fatal(".next 应被替换为文件而非目录")
	}

	if err := RestoreNextCache(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, next) {
		t.Fatal("功能 4 后 .next 文件应被删除")
	}
}

func TestEnableAndResetNodeConfig(t *testing.T) {
	s := newStub(t)
	proj := newProject(t, s.DefaultRepoRoot(), "proj-a", map[string][]byte{
		"package.json":        []byte("{}"),
		"node_modules/lib.js": []byte("x"),
	})
	var log captureLogger

	if err := EnableGlobalNodeModules(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	npmrc := filepath.Join(proj, ".npmrc")
	if !exists(t, npmrc) {
		t.Fatal("功能 1 应写入 .npmrc")
	}
	data, _ := os.ReadFile(npmrc)
	content := string(data)
	if !strings.Contains(content, "global=true") {
		t.Fatal(".npmrc 应包含 global=true")
	}
	if !strings.Contains(content, `tmp="`) {
		t.Fatalf(".npmrc 的 tmp= 值必须加引号（L1 修复），实际:\n%s", content)
	}
	if s.envs["NODE_PATH"] != s.global {
		t.Fatal("功能 1 应设置 NODE_PATH")
	}
	if exists(t, filepath.Join(proj, "node_modules")) {
		t.Fatal("功能 1 应清理本地 node_modules")
	}

	if err := ResetNodeConfig(s, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, npmrc) {
		t.Fatal("功能 2 应删除 .npmrc")
	}
	if _, ok := s.envs["NODE_PATH"]; ok {
		t.Fatal("功能 2 应删除 NODE_PATH")
	}
}

func TestRunDispatcher(t *testing.T) {
	s := newStub(t)
	newProject(t, s.DefaultRepoRoot(), "proj-a", map[string][]byte{
		"package.json": []byte("{}"),
	})
	var log captureLogger
	if err := Run(s, 7, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if err := Run(s, 9, s.DefaultRepoRoot(), i18n.Zh, Options{}, &log); err == nil {
		t.Fatal("未知功能号应返回错误")
	}
	if !IsDangerous(1) || !IsDangerous(6) || !IsDangerous(7) {
		t.Fatal("功能 1/6/7 应为危险操作")
	}
	if IsDangerous(3) || IsDangerous(4) || IsDangerous(5) {
		t.Fatal("功能 3/4/5 不应为危险操作")
	}
}
