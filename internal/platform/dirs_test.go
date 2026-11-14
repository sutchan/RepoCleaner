// internal/platform/dirs_test.go v2.3.2

package platform

import (
	"strings"
	"testing"
)

// TestIsSafeCleanDirName 覆盖合法目录名与各类路径穿越 / 系统保留名。
func TestIsSafeCleanDirName(t *testing.T) {
	valid := []string{"custom-build", ".cache", "out dir", "構建", "a1_b-2"}
	for _, name := range valid {
		if !IsSafeCleanDirName(name) {
			t.Errorf("合法目录名被拒绝: %q", name)
		}
	}
	invalid := []string{
		"", " ", " name", "name ", ".", "..", "../..", "sub/dir", `sub\dir`,
		`C:\Windows`, "/etc", "a:b", "CON", "nul.txt", "COM1", "line\nbreak",
		strings.Repeat("a", 256),
	}
	for _, name := range invalid {
		if IsSafeCleanDirName(name) {
			t.Errorf("不安全目录名被放行: %q", name)
		}
	}
}
