// internal/config/config_test.go v2.0.0

package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadMissingFile(t *testing.T) {
	cfg, err := Load(filepath.Join(t.TempDir(), "config.ini"))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Lang != "zh" {
		t.Fatalf("默认语言应为 zh，得到 %q", cfg.Lang)
	}
	if cfg.RepoRoot != "" {
		t.Fatalf("默认 REPO_ROOT 应为空，得到 %q", cfg.RepoRoot)
	}
}

func TestSaveLoadRoundtrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.ini")
	old := "REPO_ROOT=C:\\old\nLANG=en\nEXTRA=keep\n; comment line\nREPO_ROOT=dup\n"
	if err := os.WriteFile(path, []byte(old), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Save(path, Config{RepoRoot: `D:\new`, Lang: "zh"}); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.RepoRoot != `D:\new` {
		t.Fatalf("REPO_ROOT 应更新为 D:\\new，得到 %q", cfg.RepoRoot)
	}
	if cfg.Lang != "zh" {
		t.Fatalf("LANG 应更新为 zh，得到 %q", cfg.Lang)
	}

	data, _ := os.ReadFile(path)
	content := string(data)
	if !strings.Contains(content, "EXTRA=keep") {
		t.Fatal("未知键必须保留")
	}
	if !strings.Contains(content, "; comment line") {
		t.Fatal("注释必须保留")
	}
	if strings.Count(content, "REPO_ROOT=") != 1 {
		t.Fatal("REPO_ROOT 应去重且只出现一次")
	}
	if strings.Count(content, "LANG=") != 1 {
		t.Fatal("LANG 应去重且只出现一次")
	}
	if strings.Contains(content, ".tmp") {
		t.Fatal("不应残留临时文件内容")
	}
}

func TestLangWhitelist(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.ini")
	if err := os.WriteFile(path, []byte("LANG=fr\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Lang != "zh" {
		t.Fatalf("非法语言应回落 zh，得到 %q", cfg.Lang)
	}
}
