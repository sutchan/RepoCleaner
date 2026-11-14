// internal/ui/progress_parse.go v2.3.3

package ui

import (
	"regexp"
	"strconv"
	"strings"
)

var stepPattern = regexp.MustCompile(`^\[(\d+)/(\d+)\]`)

// parseStep 从执行日志提取 [n/m] 分步标记；无标记或数值非法时返回 ok=false。
func parseStep(line string) (done, total int, ok bool) {
	m := stepPattern.FindStringSubmatch(strings.TrimSpace(line))
	if m == nil {
		return 0, 0, false
	}
	d, err1 := strconv.Atoi(m[1])
	t, err2 := strconv.Atoi(m[2])
	if err1 != nil || err2 != nil || t <= 0 {
		return 0, 0, false
	}
	if d > t {
		d = t
	}
	return d, t, true
}
