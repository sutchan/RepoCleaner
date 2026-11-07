// internal/engine/engine.go v2.0.2

// Package engine 实现功能 1 ～ 7 的平台无关业务逻辑；OS 差异一律经 platform.Platform。
package engine

import (
	"fmt"
	"os"
	"path/filepath"

	"repo-cleaner/internal/i18n"
	"repo-cleaner/internal/platform"
)

// Logger 由 UI / CLI 注入，用于输出执行日志。
type Logger interface {
	Logf(format string, args ...any)
}

// StdLogger 默认输出到 stdout（命令行模式）。
type StdLogger struct{}

// Logf 实现 Logger。
func (StdLogger) Logf(format string, args ...any) { fmt.Printf(format+"\n", args...) }

// 功能 6 清理白名单（对齐 OPENSPEC §3.6）。
var cleanDirsF6 = []string{"node_modules", ".next", ".nuxt", "dist", "build"}

// 功能 6 清理锁文件白名单。
var lockFilesF6 = []string{"package-lock.json", "yarn.lock", "pnpm-lock.yaml"}

// 功能 7 清理白名单（对齐 OPENSPEC §3.7）。
var cleanDirsF7 = []string{".vite", ".turbo", ".nuxt", "dist", "build", "coverage", ".nyc_output"}

// npmGroups 功能 5 安装分组（对齐 OPENSPEC §3.5）。
var npmGroups = [][]string{
	{"react", "react-dom", "next"},
	{"vue", "@vue/cli"},
	{"axios", "express"},
	{"pnpm", "yarn"},
	{"cross-env", "dotenv", "nodemon", "pm2", "rimraf", "mkdirp"},
}

// Run 按功能编号（1-7）执行对应功能，lang 决定执行日志语言。
func Run(p platform.Platform, fn int, repoRoot string, lang i18n.Lang, log Logger) error {
	switch fn {
	case 1:
		return EnableGlobalNodeModules(p, repoRoot, lang, log)
	case 2:
		return ResetNodeConfig(p, repoRoot, lang, log)
	case 3:
		return CleanNextCache(p, repoRoot, lang, log)
	case 4:
		return RestoreNextCache(p, repoRoot, lang, log)
	case 5:
		return InstallGlobalDeps(lang, log)
	case 6:
		return CleanAllCaches(p, repoRoot, lang, log)
	case 7:
		return CleanViteCache(p, repoRoot, lang, log)
	default:
		return fmt.Errorf("unknown function: %d", fn)
	}
}

// IsDangerous 返回功能是否为危险操作（需要二次确认）：1 / 6 / 7。
func IsDangerous(fn int) bool { return fn == 1 || fn == 6 || fn == 7 }

// EnableGlobalNodeModules 功能 1：启用全局 node_modules（对齐 OPENSPEC §3.1）。
func EnableGlobalNodeModules(p platform.Platform, repoRoot string, lang i18n.Lang, log Logger) error {
	d := i18n.Logs(lang)
	log.Logf("%s", d.F1Step1)
	globalPath, err := p.GlobalNodeModulesPath()
	if err != nil {
		return err
	}
	log.Logf(d.F1Path, globalPath)

	log.Logf("%s", d.F1Step2)
	if err := p.SetUserEnv("NODE_PATH", globalPath); err != nil {
		log.Logf(d.EnvSetWarn, err)
	}

	log.Logf("%s", d.F1Step3)
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if !platform.IsNodeProject(dir) {
			continue
		}
		if err := writeNpmrc(p, dir, globalPath); err != nil {
			log.Logf(d.WriteNpmrcErr, dir, err)
			continue
		}
		log.Logf(d.Configured, filepath.Base(dir))
	}

	log.Logf("%s", d.F1Step4)
	for _, dir := range dirs {
		nm := filepath.Join(dir, "node_modules")
		if info, err := os.Stat(nm); err == nil && info.IsDir() {
			log.Logf(d.Cleaning, filepath.Base(dir))
			if err := p.RemoveAll(nm); err != nil {
				log.Logf(d.DeleteErr, err)
			}
		}
	}
	return nil
}

// writeNpmrc 写入 .npmrc；tmp= 值已加引号（v2.0.0 修复 L1，见 OPENSPEC §7.3）。
func writeNpmrc(p platform.Platform, projectDir, globalPath string) error {
	content := fmt.Sprintf("prefix=%s\ncache=%s\ntmp=%q\nglobal=true\nprefer-global=true\n",
		globalPath, filepath.Join(globalPath, ".npm-cache"), p.TempDir())
	return os.WriteFile(filepath.Join(projectDir, ".npmrc"), []byte(content), 0o644)
}

// ResetNodeConfig 功能 2：重置 Node 配置（对齐 OPENSPEC §3.2）。
func ResetNodeConfig(p platform.Platform, repoRoot string, lang i18n.Lang, log Logger) error {
	d := i18n.Logs(lang)
	log.Logf("%s", d.F2Step1)
	if err := p.RemoveUserEnv("NODE_PATH"); err != nil {
		log.Logf(d.EnvDelWarn, err)
	}
	log.Logf("%s", d.F2Step2)
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if err := p.RemoveFile(filepath.Join(dir, ".npmrc")); err == nil {
			log.Logf(d.Removed, filepath.Base(dir))
		}
	}
	return nil
}
