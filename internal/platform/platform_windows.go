//go:build windows

// internal/platform/platform_windows.go v2.0.2

package platform

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// windowsPlatform Windows 实现：attrib 隐藏属性、setx/reg 环境变量持久化。
type windowsPlatform struct{}

// New 返回当前平台的 Platform 实现（由 build tag 分发）。
func New() Platform { return windowsPlatform{} }

func (windowsPlatform) Name() string            { return "windows" }
func (windowsPlatform) DefaultRepoRoot() string { return `E:\Github` }

func (windowsPlatform) ListProjectDirs(repoRoot string) ([]string, error) {
	entries, err := os.ReadDir(repoRoot)
	if err != nil {
		return nil, err
	}
	var dirs []string
	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, filepath.Join(repoRoot, e.Name()))
		}
	}
	return dirs, nil
}

func (windowsPlatform) RemoveAll(path string) error  { return os.RemoveAll(path) }
func (windowsPlatform) RemoveFile(path string) error { return os.Remove(path) }

// SetHidden 对应 v1.x `attrib +h +s +r`。
func (windowsPlatform) SetHidden(path string) error {
	return run("attrib", "+h", "+s", "+r", path)
}

// ClearHidden 对应 v1.x `attrib -h -s -r`（功能 4 必须先清除属性再删除）。
func (windowsPlatform) ClearHidden(path string) error {
	return run("attrib", "-h", "-s", "-r", path)
}

func (windowsPlatform) TempDir() string { return os.TempDir() }

// SetUserEnv 对应 v1.x `setx`。
func (windowsPlatform) SetUserEnv(key, value string) error {
	return run("setx", key, value)
}

// RemoveUserEnv 对应 v1.x `setx KEY ""` + `reg delete HKCU\Environment /v KEY /f`。
func (windowsPlatform) RemoveUserEnv(key string) error {
	clearErr := run("setx", key, "")
	delErr := run("reg", "delete", `HKCU\Environment`, "/v", key, "/f")
	// 变量不存在时 setx 清空会失败，属预期情形；仅当注册表删除也失败时才报错。
	if delErr != nil {
		return errors.Join(clearErr, delErr)
	}
	return nil
}

func (windowsPlatform) GlobalNodeModulesPath() (string, error) {
	out, err := exec.Command("npm", "root", "-g").Output()
	if err != nil {
		return "", fmt.Errorf("npm root -g 执行失败: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// run 执行外部命令；失败时附带输出便于定位。
func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s %s: %w: %s",
			name, strings.Join(args, " "), err, strings.TrimSpace(string(out)))
	}
	return nil
}
