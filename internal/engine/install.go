// internal/engine/install.go v2.1.0

package engine

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"repo-manager/internal/i18n"
)

// npmRunner 功能 5 的安装执行入口；测试可替换以避免真实联网安装。
var npmRunner = runNpmInstall

// InstallGlobalDeps 功能 5：分组安装常用全局依赖（对齐 OPENSPEC §3.5）。
// 单组失败不中断后续组；结束后汇总失败组并返回错误（v2.1.0 修复 L4）。
// DryRun 下仅列出各组将要安装的包，不执行 npm。
func InstallGlobalDeps(lang i18n.Lang, opts Options, log Logger) error {
	d := i18n.Logs(lang)
	var failed []string
	for i, group := range npmGroups {
		log.Logf("%s", d.GroupNames[i])
		if opts.DryRun {
			log.Logf(d.DryRunInstall, strings.Join(group, " "))
			continue
		}
		out, err := npmRunner(group)
		if err != nil {
			log.Logf(d.GroupFail, i+1, err)
			if out != "" {
				log.Logf("%s", out)
			}
			failed = append(failed, fmt.Sprintf("%d", i+1))
		}
	}
	if len(failed) > 0 {
		summary := fmt.Sprintf(d.GroupSummary, strings.Join(failed, ", "))
		log.Logf("%s", summary)
		return errors.New(summary)
	}
	return nil
}

// runNpmInstall 执行 `npm install -g <包...>`，返回合并输出。
func runNpmInstall(pkgs []string) (string, error) {
	args := append([]string{"install", "-g"}, pkgs...)
	out, err := exec.Command("npm", args...).CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
