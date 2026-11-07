// internal/buildinfo/version_test.go v2.0.2

package buildinfo

import (
	"regexp"
	"testing"
	"time"
)

// TestFooterFormat 校验页脚格式为 vX.Y.Z · build yymmddhhmm。
func TestFooterFormat(t *testing.T) {
	original := BuildID
	t.Cleanup(func() { BuildID = original })

	BuildID = "2611051259"
	want := "v" + Version + " · build 2611051259"
	if got := Footer(); got != want {
		t.Fatalf("Footer() = %q, want %q", got, want)
	}
}

// TestFooterFallback 校验注入缺失时回退为仅版本号或 10 位推导值。
func TestFooterFallback(t *testing.T) {
	original := BuildID
	t.Cleanup(func() { BuildID = original })

	BuildID = "invalid"
	got := Footer()
	if !regexp.MustCompile(`^v\d+\.\d+\.\d+( · build \d{10})?$`).MatchString(got) {
		t.Fatalf("Footer() 回退格式不合法: %q", got)
	}
}

// TestBuildIDLayout 校验 yymmddhhmm 布局为 10 位。
func TestBuildIDLayout(t *testing.T) {
	got := time.Date(2026, 11, 5, 12, 59, 0, 0, time.Local).Format("0601021504")
	if got != "2611051259" {
		t.Fatalf("build 编号 = %q, want 2611051259", got)
	}
}
