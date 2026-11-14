// internal/platform/platform.go v2.3.3

// Package platform 收口全部操作系统差异（目录扫描、删除、隐藏属性、环境变量持久化）。
// engine 只依赖本接口，禁止直接调用 OS 专属命令；新增平台仅需增加对应 build tag 实现。
package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Platform 抽象操作系统差异，是 engine 与操作系统之间的唯一边界。
type Platform interface {
	// Name 返回平台标识（windows / darwin / linux）。
	Name() string
	// DefaultRepoRoot 返回默认仓库根目录（Windows: E:\Github；Unix: ~/Github）。
	DefaultRepoRoot() string
	// ListProjectDirs 枚举 repoRoot 的直接子目录（仅目录，不递归）。
	ListProjectDirs(repoRoot string) ([]string, error)
	// RemoveAll 删除目录及其全部内容。
	RemoveAll(path string) error
	// RemoveFile 删除单个文件。
	RemoveFile(path string) error
	// SetHidden 设置隐藏属性（功能 3）。
	SetHidden(path string) error
	// ClearHidden 清除隐藏属性（功能 4，必须先于删除，否则删除静默失败）。
	ClearHidden(path string) error
	// TempDir 返回系统临时目录（写入 .npmrc 的 tmp=）。
	TempDir() string
	// SetUserEnv 持久化写入用户级环境变量。
	SetUserEnv(key, value string) error
	// RemoveUserEnv 删除用户级环境变量。
	RemoveUserEnv(key string) error
	// GlobalNodeModulesPath 返回 `npm root -g` 的全局模块路径。
	GlobalNodeModulesPath() (string, error)
}

// deepSkipDirs 递归枚举时跳过的依赖 / 构建产物目录名（避免深入依赖树与输出目录）。
var deepSkipDirs = map[string]bool{
	"node_modules": true, ".next": true, ".nuxt": true, "dist": true, "build": true,
	".vite": true, ".turbo": true, ".angular": true, ".svelte-kit": true,
	"coverage": true, ".nyc_output": true, ".git": true,
}

// ListProjectDirsDeep 递归枚举 repoRoot 下 depthLimit 层内的目录（1 = 仅直接子目录，
// 2 = 含孙目录，V1 monorepo 支持）。跳过 deepSkipDirs 中的目录；单个子树读取失败
// 不影响其余结果。返回排序后的路径列表（v2.3.0，对齐 OPENSPEC §3.0）。
func ListProjectDirsDeep(repoRoot string, depthLimit int) ([]string, error) {
	var dirs []string
	seen := map[string]bool{}
	var walk func(dir string, depth int)
	walk = func(dir string, depth int) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return // 子树不可读时跳过，不影响其余枚举
		}
		for _, e := range entries {
			if !e.IsDir() || deepSkipDirs[e.Name()] {
				continue
			}
			full := filepath.Join(dir, e.Name())
			if !seen[full] {
				seen[full] = true
				dirs = append(dirs, full)
			}
			if depth < depthLimit {
				walk(full, depth+1)
			}
		}
	}
	walk(repoRoot, 1)
	sort.Strings(dirs)
	return dirs, nil
}

// HasFile 判断 dir 下是否存在指定文件（探测 package.json / next.config.*）。
func HasFile(dir, name string) bool {
	info, err := os.Stat(filepath.Join(dir, name))
	return err == nil && !info.IsDir()
}

// IsNodeProject 判断目录是否为 Node 项目（含 package.json）。
func IsNodeProject(dir string) bool { return HasFile(dir, "package.json") }

// IsNextProject 判断目录是否为 Next.js 项目（含 next.config.js 或 next.config.ts）。
func IsNextProject(dir string) bool {
	return HasFile(dir, "next.config.js") || HasFile(dir, "next.config.ts")
}

// NormalizeRoot 规范化仓库根目录：去空白、去末尾分隔符、转绝对路径（对齐 OPENSPEC §7.1）。
func NormalizeRoot(root string) string {
	root = strings.TrimSpace(root)
	root = strings.TrimRight(root, `\/`)
	if abs, err := filepath.Abs(root); err == nil {
		return abs
	}
	return root
}

// ValidateRoot 校验仓库根目录存在且为目录。
func ValidateRoot(root string) error {
	info, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("路径不存在: %s", root)
	}
	if !info.IsDir() {
		return fmt.Errorf("不是目录: %s", root)
	}
	return nil
}
