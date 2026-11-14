// internal/platform/dirs.go v2.3.2

package platform

import (
	"path/filepath"
	"strings"
)

// IsSafeCleanDirName 校验自定义清理目录是否为「项目目录内的单个普通目录名」。
// 拒绝路径穿越、绝对路径、盘符、分隔符、空白与 Windows 保留设备名，
// 保证 filepath.Join(project, name) 不会指向 project 之外。
func IsSafeCleanDirName(name string) bool {
	if name == "" || name != strings.TrimSpace(name) {
		return false
	}
	if name == "." || name == ".." || len(name) > 255 {
		return false
	}
	if strings.ContainsAny(name, `/\`) || strings.Contains(name, ":") {
		return false
	}
	if filepath.IsAbs(name) || filepath.Base(name) != name {
		return false
	}
	if isReservedDeviceName(name) {
		return false
	}
	// 控制字符会造成跨平台路径解析异常，直接拒绝。
	for _, r := range name {
		if r < 0x20 || r == 0x7f {
			return false
		}
	}
	return true
}

// isReservedDeviceName 判断 Windows 设备保留名（不区分大小写，忽略扩展名）。
func isReservedDeviceName(name string) bool {
	base := strings.ToLower(name)
	if dot := strings.IndexByte(base, '.'); dot >= 0 {
		base = base[:dot]
	}
	switch base {
	case "con", "prn", "aux", "nul":
		return true
	}
	if len(base) == 4 && (strings.HasPrefix(base, "com") || strings.HasPrefix(base, "lpt")) {
		return base[3] >= '1' && base[3] <= '9'
	}
	return false
}
