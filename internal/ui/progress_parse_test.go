// internal/ui/progress_parse_test.go v2.3.3

package ui

import "testing"

func TestParseStep(t *testing.T) {
	cases := []struct {
		line        string
		done, total int
		ok          bool
	}{
		{"[1/4] 获取全局 node_modules 路径...", 1, 4, true},
		{"[5/5] Installing development tools...", 5, 5, true},
		{"[9/5] overflow", 5, 5, true},
		{"  [2/3] 带前导空格", 2, 3, true},
		{"普通日志", 0, 0, false},
		{"[x/4] bad", 0, 0, false},
		{"[1/0] zero", 0, 0, false},
	}
	for _, c := range cases {
		d, n, ok := parseStep(c.line)
		if d != c.done || n != c.total || ok != c.ok {
			t.Fatalf("%q => (%d,%d,%v)，期望 (%d,%d,%v)", c.line, d, n, ok, c.done, c.total, c.ok)
		}
	}
}
