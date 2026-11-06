// internal/engine/engine.go v2.0.0

// Package engine 实现功能 1 ～ 7 的平台无关业务逻辑；OS 差异一律经 platform.Platform。
package engine

import (
	"fmt"
	"os"
	"path/filepath"

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

// Run 按功能编号（1-7）执行对应功能。
func Run(p platform.Platform, fn int, repoRoot string, log Logger) error {
	switch fn {
	case 1:
		return EnableGlobalNodeModules(p, repoRoot, log)
	case 2:
		return ResetNodeConfig(p, repoRoot, log)
	case 3:
		return CleanNextCache(p, repoRoot, log)
	case 4:
		return RestoreNextCache(p, repoRoot, log)
	case 5:
		return InstallGlobalDeps(log)
	case 6:
		return CleanAllCaches(p, repoRoot, log)
	case 7:
		return CleanViteCache(p, repoRoot, log)
	default:
		return fmt.Errorf("未知功能编号: %d", fn)
	}
}

// IsDangerous 返回功能是否为危险操作（需要二次确认）：1 / 6 / 7。
func IsDangerous(fn int) bool { return fn == 1 || fn == 6 || fn == 7 }

// EnableGlobalNodeModules 功能 1：启用全局 node_modules（对齐 OPENSPEC §3.1）。
func EnableGlobalNodeModules(p platform.Platform, repoRoot string, log Logger) error {
	log.Logf("[1/4] 获取全局 node_modules 路径...")
	globalPath, err := p.GlobalNodeModulesPath()
	if err != nil {
		return err
	}
	log.Logf("Path: %s", globalPath)

	log.Logf("[2/4] 设置 NODE_PATH 环境变量...")
	if err := p.SetUserEnv("NODE_PATH", globalPath); err != nil {
		log.Logf("[警告] 写入环境变量失败: %v", err)
	}

	log.Logf("[3/4] 禁用所有项目的本地 node_modules...")
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if !platform.IsNodeProject(dir) {
			continue
		}
		if err := writeNpmrc(p, dir, globalPath); err != nil {
			log.Logf("[错误] 写入 %s 失败: %v", dir, err)
			continue
		}
		log.Logf("Configured: %s", filepath.Base(dir))
	}

	log.Logf("[4/4] 清理现有本地 node_modules 目录...")
	for _, dir := range dirs {
		nm := filepath.Join(dir, "node_modules")
		if info, err := os.Stat(nm); err == nil && info.IsDir() {
			log.Logf("Cleaning: %s", filepath.Base(dir))
			if err := p.RemoveAll(nm); err != nil {
				log.Logf("[错误] 删除失败: %v", err)
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
func ResetNodeConfig(p platform.Platform, repoRoot string, log Logger) error {
	log.Logf("删除 NODE_PATH 环境变量...")
	if err := p.RemoveUserEnv("NODE_PATH"); err != nil {
		log.Logf("[警告] 删除环境变量失败: %v", err)
	}
	log.Logf("清理项目 .npmrc 配置...")
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if err := p.RemoveFile(filepath.Join(dir, ".npmrc")); err == nil {
			log.Logf("Removed: %s", filepath.Base(dir))
		}
	}
	return nil
}
