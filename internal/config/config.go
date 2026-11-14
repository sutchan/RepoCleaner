// internal/config/config.go v2.3.2

// Package config 负责 config.ini 的读写：读旧 → 过滤目标键 → 补写 → 回写，
// 保留注释与未知行（对齐 OPENSPEC §5.3 与 CONTRIBUTING.md §4.5 约定）。
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"repo-cleaner/internal/platform"
)

// Config 产品运行配置。
type Config struct {
	RepoRoot       string   // 仓库根目录
	Lang           string   // 界面语言：zh / en
	Recursive      bool     // 递归扫描 monorepo 深层项目（V1，默认关闭）
	ExtraCleanDirs []string // 自定义清理目录名（V3，逗号分隔，仅限项目目录内部）
}

// Default 返回默认配置（REPO_ROOT 留空，由调用方按平台取默认目录）。
func Default() Config { return Config{RepoRoot: "", Lang: "zh"} }

// Path 返回与可执行文件同目录的 config.ini 路径（对齐 v1.x 行为）。
func Path() string {
	exe, err := os.Executable()
	if err != nil {
		return "config.ini"
	}
	return filepath.Join(filepath.Dir(exe), "config.ini")
}

// Load 读取配置；文件不存在时返回默认值。
func Load(path string) (Config, error) {
	cfg := Default()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		switch strings.ToUpper(strings.TrimSpace(key)) {
		case "REPO_ROOT":
			cfg.RepoRoot = strings.TrimSpace(value)
		case "LANG":
			cfg.Lang = strings.TrimSpace(value)
		case "RECURSIVE":
			v := strings.TrimSpace(value)
			cfg.Recursive = v == "true" || v == "1" || strings.EqualFold(v, "yes")
		case "EXTRA_CLEAN_DIRS":
			cfg.ExtraCleanDirs = parseDirList(value)
		}
	}
	if cfg.Lang != "en" {
		cfg.Lang = "zh" // 白名单校验，非法值回落中文
	}
	return cfg, nil
}

// parseDirList 解析逗号分隔的目录名列表（去空白、去空项），
// 并丢弃会造成路径穿越的非法目录名，避免清理目标指向项目目录之外。
func parseDirList(value string) []string {
	var names []string
	for _, part := range strings.Split(value, ",") {
		name := strings.TrimSpace(part)
		if !platform.IsSafeCleanDirName(name) {
			continue
		}
		names = append(names, name)
	}
	return names
}

// Save 以「读旧 → 过滤目标键 → 补写 → 临时文件 → 回写」方式保存，去重同名键。
func Save(path string, cfg Config) error {
	var lines []string
	if data, err := os.ReadFile(path); err == nil {
		lines = strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	}
	repoWritten, langWritten, recWritten, extraWritten := false, false, false, false
	extraLine := "EXTRA_CLEAN_DIRS=" + strings.Join(cfg.ExtraCleanDirs, ",")
	recLine := "RECURSIVE=false"
	if cfg.Recursive {
		recLine = "RECURSIVE=true"
	}
	for i, line := range lines {
		key, _, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		switch strings.ToUpper(strings.TrimSpace(key)) {
		case "REPO_ROOT":
			if !repoWritten {
				lines[i] = "REPO_ROOT=" + cfg.RepoRoot
				repoWritten = true
			} else {
				lines[i] = ""
			}
		case "LANG":
			if !langWritten {
				lines[i] = "LANG=" + cfg.Lang
				langWritten = true
			} else {
				lines[i] = ""
			}
		case "RECURSIVE":
			if !recWritten {
				lines[i] = recLine
				recWritten = true
			} else {
				lines[i] = ""
			}
		case "EXTRA_CLEAN_DIRS":
			if !extraWritten {
				lines[i] = extraLine
				extraWritten = true
			} else {
				lines[i] = ""
			}
		}
	}
	if !repoWritten {
		lines = append(lines, "REPO_ROOT="+cfg.RepoRoot)
	}
	if !langWritten {
		lines = append(lines, "LANG="+cfg.Lang)
	}
	if !recWritten {
		lines = append(lines, recLine)
	}
	if !extraWritten {
		lines = append(lines, extraLine)
	}
	content := strings.Join(lines, "\n")
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(content), 0o644); err != nil {
		return fmt.Errorf("写入临时文件失败: %w", err)
	}
	return os.Rename(tmp, path)
}
