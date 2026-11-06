// internal/engine/install.go v2.0.0

package engine

import (
	"os/exec"
	"strings"
)

// InstallGlobalDeps 功能 5：分组安装常用全局依赖（对齐 OPENSPEC §3.5）。
// 单组失败不中断后续组（沿用 v1.x 行为），失败组记录日志。
func InstallGlobalDeps(log Logger) error {
	groupNames := [5]string{
		"[1/5] 安装 React 生态...",
		"[2/5] 安装 Vue 生态...",
		"[3/5] 安装 HTTP 库...",
		"[4/5] 安装包管理器...",
		"[5/5] 安装开发工具...",
	}
	for i, group := range npmGroups {
		log.Logf("%s", groupNames[i])
		out, err := runNpmInstall(group)
		if err != nil {
			log.Logf("[错误] 第 %d 组安装失败: %v", i+1, err)
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
