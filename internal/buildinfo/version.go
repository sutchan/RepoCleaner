// internal/buildinfo/version.go v2.1.0

// Package buildinfo 是产品版本的唯一权威来源（见 docs/CONTRIBUTING.md §8.2）。
package buildinfo

import (
	"fmt"
	"runtime/debug"
	"time"
)

// Version 产品版本常量；bump 时须同步全部全局展示位（CONTRIBUTING.md §8.2 清单）。
const Version = "2.1.0"

// BuildID 构建编号，格式 yymmddhhmm（10 位），由 scripts/build.ps1 / build.sh
// 在构建期以 -ldflags "-X repo-cleaner/internal/buildinfo.BuildID=<10位>" 注入。
// 未注入时 Footer 回退到 VCS 构建时间推导，仍缺失则只显示版本号。
var BuildID = ""

// Footer 返回应用页脚唯一展示的「版本 + build 编号」文本。
// 格式：vX.Y.Z · build yymmddhhmm，例如 v1.2.3 · build 2611051259。
func Footer() string {
	if id := effectiveBuildID(); id != "" {
		return fmt.Sprintf("v%s · build %s", Version, id)
	}
	return fmt.Sprintf("v%s", Version)
}

// BuildInfo 返回命令行 -h 使用的完整构建信息（页脚文本 + VCS 短提交号）。
func BuildInfo() string {
	rev, _ := vcsSettings()
	base := Footer()
	if len(rev) >= 7 {
		return fmt.Sprintf("%s · %s", base, rev[:7])
	}
	return base
}

// effectiveBuildID 优先采用构建期注入值，其次由 vcs.time 推导本地时间。
func effectiveBuildID() string {
	if len(BuildID) == 10 && isDigits(BuildID) {
		return BuildID
	}
	_, vcsTime := vcsSettings()
	if vcsTime == "" {
		return ""
	}
	t, err := time.Parse(time.RFC3339, vcsTime)
	if err != nil {
		return ""
	}
	return t.Local().Format("0601021504")
}

// isDigits 判断字符串是否全为十进制数字。
func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// vcsSettings 读取 debug.ReadBuildInfo 注入的 VCS 提交号与构建时间。
func vcsSettings() (revision, vcsTime string) {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.time":
				vcsTime = s.Value
			}
		}
	}
	return revision, vcsTime
}
