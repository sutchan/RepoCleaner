// internal/config/keys_test.go v2.3.0

package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadNewKeys(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.ini")
	content := "REPO_ROOT=E:\\Github\nLANG=zh\nRECURSIVE=true\nEXTRA_CLEAN_DIRS=.cache, custom-out ,\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.Recursive {
		t.Fatal("RECURSIVE=true should parse to true")
	}
	if len(cfg.ExtraCleanDirs) != 2 || cfg.ExtraCleanDirs[0] != ".cache" || cfg.ExtraCleanDirs[1] != "custom-out" {
		t.Fatalf("EXTRA_CLEAN_DIRS should be trimmed and split, got %v", cfg.ExtraCleanDirs)
	}
}

func TestLoadDefaultRecursive(t *testing.T) {
	cfg, err := Load(filepath.Join(t.TempDir(), "missing.ini"))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Recursive {
		t.Fatal("Recursive should default to false")
	}
}

func TestSaveRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.ini")
	if err := os.WriteFile(path, []byte("; comment kept\nREPO_ROOT=D:\\Old\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cfg := Config{RepoRoot: "D:\\New", Lang: "en", Recursive: true, ExtraCleanDirs: []string{"a", "b"}}
	if err := Save(path, cfg); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.RepoRoot != "D:\\New" || got.Lang != "en" || !got.Recursive {
		t.Fatalf("round-trip mismatch: %+v", got)
	}
	if len(got.ExtraCleanDirs) != 2 {
		t.Fatalf("EXTRA_CLEAN_DIRS round-trip mismatch: %v", got.ExtraCleanDirs)
	}
	data, _ := os.ReadFile(path)
	if !strings.Contains(string(data), "; comment kept") {
		t.Fatal("comments and unknown lines should be preserved")
	}
}
