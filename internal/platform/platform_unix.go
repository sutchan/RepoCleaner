//go:build !windows

// internal/platform/platform_unix.go v2.0.2

package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// unixPlatform macOS / Linux 实现：chflags / 只读位隐藏、~/.profile 环境变量持久化。
type unixPlatform struct{}

// New 返回当前平台的 Platform 实现（由 build tag 分发）。
func New() Platform { return unixPlatform{} }

func (unixPlatform) Name() string { return runtime.GOOS }

func (unixPlatform) DefaultRepoRoot() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "Github"
	}
	return filepath.Join(home, "Github")
}

func (unixPlatform) ListProjectDirs(repoRoot string) ([]string, error) {
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

func (unixPlatform) RemoveAll(path string) error  { return os.RemoveAll(path) }
func (unixPlatform) RemoveFile(path string) error { return os.Remove(path) }

// SetHidden 隐藏 .next：macOS 用 chflags hidden，并统一设置只读位；
// Linux 下点前缀天然隐藏，无需额外命令。
func (unixPlatform) SetHidden(path string) error {
	if err := applyChflags(path, "hidden"); err != nil {
		return err
	}
	return os.Chmod(path, 0o444)
}

// ClearHidden 先清除隐藏标志再恢复写权限（功能 4 必须先于删除）。
func (unixPlatform) ClearHidden(path string) error {
	if err := applyChflags(path, "nohidden"); err != nil {
		return err
	}
	return os.Chmod(path, 0o644)
}

// applyChflags 仅在 macOS 且系统提供 chflags 时执行；其余 Unix 平台降级为点前缀隐藏。
func applyChflags(path, flag string) error {
	if runtime.GOOS != "darwin" {
		return nil
	}
	bin, err := exec.LookPath("chflags")
	if err != nil {
		return nil
	}
	out, err := exec.Command(bin, flag, path).CombinedOutput()
	if err != nil {
		return fmt.Errorf("chflags %s 执行失败: %w: %s", flag, err, strings.TrimSpace(string(out)))
	}
	return nil
}

func (unixPlatform) TempDir() string { return os.TempDir() }

// profilePath 返回环境变量持久化目标文件（~/.profile）。
func profilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".profile"), nil
}

// SetUserEnv 以「读旧 → 过滤同名行 → 追加 → 回写」写入 ~/.profile。
func (unixPlatform) SetUserEnv(key, value string) error {
	if err := removeProfileLine(key); err != nil {
		return err
	}
	path, err := profilePath()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("打开 %s 失败: %w", path, err)
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("export %s=%q\n", key, value))
	return err
}

// RemoveUserEnv 从 ~/.profile 过滤掉目标变量行。
func (unixPlatform) RemoveUserEnv(key string) error { return removeProfileLine(key) }

func removeProfileLine(key string) error {
	path, err := profilePath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	prefix := "export " + key + "="
	var kept []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, prefix) {
			continue
		}
		kept = append(kept, line)
	}
	return os.WriteFile(path, []byte(strings.Join(kept, "\n")), 0o644)
}

func (unixPlatform) GlobalNodeModulesPath() (string, error) {
	out, err := exec.Command("npm", "root", "-g").Output()
	if err != nil {
		return "", fmt.Errorf("npm root -g 执行失败: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
