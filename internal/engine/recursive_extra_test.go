// internal/engine/recursive_extra_test.go v2.3.2

package engine

import (
	"path/filepath"
	"strings"
	"testing"

	"repo-manager/internal/i18n"
)

// TestRecursiveClean V1：Recursive 开启时清理孙目录项目（深度 3 层），
// 关闭时仅处理直接子目录。
func TestRecursiveClean(t *testing.T) {
	s := newStub(t)
	root := s.DefaultRepoRoot()
	direct := newProject(t, root, "app-direct", map[string][]byte{
		"package.json":                     []byte("{}"),
		"node_modules/l.js":                []byte("x"),
		"packages/inner/package.json":      []byte("{}"),
		"packages/inner/node_modules/l.js": []byte("x"),
	})
	inner := filepath.Join(direct, "packages", "inner")

	var log captureLogger
	if err := CleanAllCaches(s, root, i18n.Zh, Options{}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(direct, "node_modules")) {
		t.Fatal("直接子项目 node_modules 应被删除")
	}
	if !exists(t, filepath.Join(inner, "node_modules")) {
		t.Fatal("未开启递归时孙目录项目不应被处理")
	}

	var log2 captureLogger
	if err := Run(s, 6, root, i18n.Zh, Options{Recursive: true}, &log2); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(inner, "node_modules")) {
		t.Fatal("开启递归后孙目录项目 node_modules 应被删除")
	}
	if !strings.Contains(strings.Join(log2.lines, "\n"), "[递归]") {
		t.Fatalf("递归开启时应输出提示日志，得到 %v", log2.lines)
	}
}

// TestExtraCleanDirs V3：自定义清理目录并入功能 7 白名单，仅限项目目录内部匹配。
func TestExtraCleanDirs(t *testing.T) {
	s := newStub(t)
	root := s.DefaultRepoRoot()
	proj := newProject(t, root, "spa-app", map[string][]byte{
		"package.json":        []byte("{}"),
		"custom-build/out.js": []byte("x"),
		"src/main.js":         []byte("keep"),
	})

	var log captureLogger
	if err := CleanViteCache(s, root, i18n.Zh, Options{ExtraCleanDirs: []string{"custom-build"}}, &log); err != nil {
		t.Fatal(err)
	}
	if exists(t, filepath.Join(proj, "custom-build")) {
		t.Fatal("自定义清理目录应被删除")
	}
	if !exists(t, filepath.Join(proj, "src", "main.js")) {
		t.Fatal("白名单之外的文件必须保留")
	}
}

// TestExtraCleanDirsRejectsTraversal engine 层第二道防线：直接构造 Options 也不允许删除 project/..。
func TestExtraCleanDirsRejectsTraversal(t *testing.T) {
	s := newStub(t)
	root := s.DefaultRepoRoot()
	proj := newProject(t, root, "app", map[string][]byte{
		"package.json":    []byte("{}"),
		"keep/marker.txt": []byte("keep"),
		"custom/out.js":   []byte("x"),
	})
	sibling := newProject(t, root, "sibling", map[string][]byte{"package.json": []byte("{}")})

	var log captureLogger
	opts := Options{ExtraCleanDirs: []string{"..", "custom"}}
	if err := CleanAllCaches(s, root, i18n.Zh, opts, &log); err != nil {
		t.Fatal(err)
	}
	if !exists(t, root) || !exists(t, sibling) || !exists(t, filepath.Join(proj, "keep", "marker.txt")) {
		t.Fatal("发生路径穿越或越界删除：仓库根目录 / 兄弟项目 / 白名单外文件被误删")
	}
	if exists(t, filepath.Join(proj, "custom")) {
		t.Fatal("合法的自定义目录仍应被清理")
	}
	if !strings.Contains(strings.Join(log.lines, "\n"), "不安全") {
		t.Fatalf("拒绝非法目录名时应输出警告，得到 %v", log.lines)
	}
}
