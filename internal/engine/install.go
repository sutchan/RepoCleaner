// internal/engine/install.go v2.0.2

package engine

import (
	"os/exec"
	"strings"

	"repo-cleaner/internal/i18n"
)

// InstallGlobalDeps 功能 5：分组安装常用全局依赖（对齐 OPENSPEC §3.5）。
// 单组失败不中断后续组（沿用 v1.x 行为），失败组记录日志。
func InstallGlobalDeps(lang i18n.Lang, log Logger) error {
	d := i18n.Logs(lang)
	for i, group := range npmGroups {
		log.Logf("%s", d.GroupNames[i])
		out, err := runNpmInstall(group)
		if err != nil {
			log.Logf(d.GroupFail, i+1, err)
			if out != "" {
				log.Logf("%s", out)
			}
		}
	}
	return nil
}

// runNpmInstall 执行 `npm install -g <包...>`，返回合并输出。
func runNpmInstall(pkgs []string) (string, error) {
	args := append([]string{"install", "-g"}, pkgs...)
	out, err := exec.Command("npm", args...).CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
