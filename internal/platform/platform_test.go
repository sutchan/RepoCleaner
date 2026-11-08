// internal/platform/platform_test.go v2.3.0

package platform

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestListProjectDirs(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"proj-a", "proj-b"} {
		if err := os.Mkdir(filepath.Join(root, name), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "file.txt"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	dirs, err := New().ListProjectDirs(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(dirs) != 2 {
		t.Fatalf("期望 2 个子目录，得到 %d", len(dirs))
	}
}

func TestNormalizeRoot(t *testing.T) {
	if got := NormalizeRoot(" /tmp/x/ "); filepath.Base(got) != "x" {
		t.Fatalf("末尾分隔符与空白应被去除，得到 %q", got)
	}
	if got := NormalizeRoot("relative/path"); filepath.IsAbs(got) != filepath.IsAbs(mustAbs(t)) {
		t.Fatalf("相对路径应转为绝对路径（随进程工作目录），得到 %q", got)
	}
}

func mustAbs(t *testing.T) string {
	t.Helper()
	abs, err := filepath.Abs(".")
	if err != nil {
		t.Fatal(err)
	}
	return abs
}

func TestValidateRoot(t *testing.T) {
	dir := t.TempDir()
	if err := ValidateRoot(dir); err != nil {
		t.Fatalf("目录应通过校验: %v", err)
	}
	file := filepath.Join(dir, "f.txt")
	if err := os.WriteFile(file, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ValidateRoot(file); err == nil {
		t.Fatal("文件不应通过目录校验")
	}
	if err := ValidateRoot(filepath.Join(dir, "missing")); err == nil {
		t.Fatal("不存在的路径不应通过校验")
	}
}

func TestProjectDetect(t *testing.T) {
	dir := t.TempDir()
	if IsNodeProject(dir) {
		t.Fatal("空目录不应识别为 Node 项目")
	}
	if err := os.WriteFile(filepath.Join(dir, "package.json"), []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if !IsNodeProject(dir) {
		t.Fatal("应识别为 Node 项目")
	}
	if IsNextProject(dir) {
		t.Fatal("未含 next.config.* 不应识别为 Next.js 项目")
	}
	if err := os.WriteFile(filepath.Join(dir, "next.config.ts"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if !IsNextProject(dir) {
		t.Fatal("应识别为 Next.js 项目")
	}
}

func TestListProjectDirsDeep(t *testing.T) {
	root := t.TempDir()
	// 结构：root/proj-a（直接子目录）、root/proj-a/inner（孙目录）、
	// root/proj-a/inner/node_modules 与 .git（应跳过）、root/proj-b（直接子目录）。
	for _, rel := range []string{
		"proj-a",
		"proj-a/inner",
		"proj-a/inner/node_modules/dep",
		"proj-a/inner/.git/objects",
		"proj-b",
	} {
		if err := os.MkdirAll(filepath.Join(root, rel), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	dirs, err := ListProjectDirsDeep(root, 2)
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(dirs, ";")
	for _, want := range []string{"proj-a", "proj-b", filepath.Join("proj-a", "inner")} {
		if !strings.Contains(joined, want) {
			t.Fatalf("递归枚举应包含 %s，得到 %v", want, dirs)
		}
	}
	if strings.Contains(joined, "node_modules") || strings.Contains(joined, ".git") {
		t.Fatalf("依赖 / VCS 目录应被跳过，得到 %v", dirs)
	}
	if len(dirs) != 3 {
		t.Fatalf("期望 3 个目录（proj-a、inner、proj-b），得到 %d: %v", len(dirs), dirs)
	}
}
