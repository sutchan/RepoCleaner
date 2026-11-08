// internal/i18n/log.go v2.3.0

package i18n

// LogDict 是功能执行日志、命令行提示与 nogui 提示的文案字典。
// 所有值均为 fmt 格式串，调用方按原有参数顺序传参，禁止在 engine / ui 层硬编码文案。
type LogDict struct {
	F1Step1       string
	F1Step2       string
	F1Step3       string
	F1Step4       string
	F1Path        string
	EnvSetWarn    string
	EnvDelWarn    string
	WriteNpmrcErr string
	Configured    string
	Cleaning      string
	DeleteErr     string
	F2Step1       string
	F2Step2       string
	Removed       string
	F3Start       string
	Processing    string
	RemoveDirErr  string
	ClearAttrWarn string
	SetAttrWarn   string
	WriteNextErr  string
	F4Start       string
	Restoring     string
	RemoveErr     string
	RemoveNameErr string
	GroupNames    [5]string
	GroupFail     string
	GroupSummary  string
	DryRunOn      string
	RecursiveOn   string
	DryRunWrite   string
	DryRunDelete  string
	DryRunEnv     string
	DryRunEnvDel  string
	DryRunInstall string
	ErrorPrefix   string
	ConfirmCLI    string
	Cancelled     string
	LoadCfgErr    string
	SaveCfgErr    string
	NoGUI         [2]string
}

var logDicts = map[Lang]LogDict{
	Zh: {
		F1Step1:       "[1/4] 获取全局 node_modules 路径...",
		F1Step2:       "[2/4] 设置 NODE_PATH 环境变量...",
		F1Step3:       "[3/4] 禁用所有项目的本地 node_modules...",
		F1Step4:       "[4/4] 清理现有本地 node_modules 目录...",
		F1Path:        "Path: %s",
		EnvSetWarn:    "[警告] 写入环境变量失败: %v",
		EnvDelWarn:    "[警告] 删除环境变量失败: %v",
		WriteNpmrcErr: "[错误] 写入 %s 失败: %v",
		Configured:    "Configured: %s",
		Cleaning:      "Cleaning: %s",
		DeleteErr:     "[错误] 删除失败: %v",
		F2Step1:       "删除 NODE_PATH 环境变量...",
		F2Step2:       "清理项目 .npmrc 配置...",
		Removed:       "Removed: %s",
		F3Start:       "扫描并清理 Next.js 缓存...",
		Processing:    "Processing: %s",
		RemoveDirErr:  "[错误] 删除目录失败: %v",
		ClearAttrWarn: "[警告] 清除属性失败: %v",
		SetAttrWarn:   "[警告] 设置隐藏属性失败: %v",
		WriteNextErr:  "[错误] 写入 .next 文件失败: %v",
		F4Start:       "恢复 Next.js 缓存功能...",
		Restoring:     "Restoring: %s",
		RemoveErr:     "[错误] 删除失败: %v",
		RemoveNameErr: "[错误] 删除 %s 失败: %v",
		GroupNames: [5]string{
			"[1/5] 安装 React 生态...",
			"[2/5] 安装 Vue 生态...",
			"[3/5] 安装 HTTP 库...",
			"[4/5] 安装包管理器...",
			"[5/5] 安装开发工具...",
		},
		GroupFail:     "[错误] 第 %d 组安装失败: %v",
		GroupSummary:  "以下分组安装失败: %s",
		DryRunOn:      "=== 预演模式 (dry-run)：仅显示将要执行的操作，不会修改任何文件 ===",
		RecursiveOn:   "[递归] 已启用递归扫描（深度 2 层，覆盖 monorepo 深层项目）",
		DryRunWrite:   "[预演] 将写入: %s",
		DryRunDelete:  "[预演] 将删除: %s",
		DryRunEnv:     "[预演] 将设置环境变量: %s",
		DryRunEnvDel:  "[预演] 将删除环境变量: %s",
		DryRunInstall: "[预演] 将安装: %s",
		ErrorPrefix:   "[错误] %v",
		ConfirmCLI:    "是否继续? [Y/N]: ",
		Cancelled:     "已取消",
		LoadCfgErr:    "读取配置失败: %v\n",
		SaveCfgErr:    "保存配置失败: %v\n",
		NoGUI: [2]string{
			"本构建未包含 GUI（构建标签 nogui）。",
			"请使用命令行模式：repo-cleaner <1-7> [目录路径]，或 repo-cleaner -h 查看帮助。",
		},
	},
	En: {
		F1Step1:       "[1/4] Resolving global node_modules path...",
		F1Step2:       "[2/4] Setting NODE_PATH...",
		F1Step3:       "[3/4] Disabling local node_modules for all projects...",
		F1Step4:       "[4/4] Removing existing local node_modules...",
		F1Path:        "Path: %s",
		EnvSetWarn:    "[WARN] Failed to write environment variable: %v",
		EnvDelWarn:    "[WARN] Failed to remove environment variable: %v",
		WriteNpmrcErr: "[ERROR] Failed to write %s: %v",
		Configured:    "Configured: %s",
		Cleaning:      "Cleaning: %s",
		DeleteErr:     "[ERROR] Delete failed: %v",
		F2Step1:       "Removing NODE_PATH environment variable...",
		F2Step2:       "Removing project .npmrc files...",
		Removed:       "Removed: %s",
		F3Start:       "Scanning and cleaning Next.js cache...",
		Processing:    "Processing: %s",
		RemoveDirErr:  "[ERROR] Failed to remove directory: %v",
		ClearAttrWarn: "[WARN] Failed to clear attributes: %v",
		SetAttrWarn:   "[WARN] Failed to set hidden attributes: %v",
		WriteNextErr:  "[ERROR] Failed to write .next file: %v",
		F4Start:       "Restoring Next.js cache...",
		Restoring:     "Restoring: %s",
		RemoveErr:     "[ERROR] Delete failed: %v",
		RemoveNameErr: "[ERROR] Failed to delete %s: %v",
		GroupNames: [5]string{
			"[1/5] Installing React packages...",
			"[2/5] Installing Vue packages...",
			"[3/5] Installing HTTP libraries...",
			"[4/5] Installing package managers...",
			"[5/5] Installing development tools...",
		},
		GroupFail:     "[ERROR] Group %d installation failed: %v",
		GroupSummary:  "Failed install groups: %s",
		DryRunOn:      "=== Dry-run mode: showing planned actions only, nothing will be modified ===",
		RecursiveOn:   "[RECURSIVE] Recursive scan enabled (depth 2, covers monorepo sub-projects)",
		DryRunWrite:   "[DRY-RUN] Would write: %s",
		DryRunDelete:  "[DRY-RUN] Would delete: %s",
		DryRunEnv:     "[DRY-RUN] Would set env: %s",
		DryRunEnvDel:  "[DRY-RUN] Would remove env: %s",
		DryRunInstall: "[DRY-RUN] Would install: %s",
		ErrorPrefix:   "[ERROR] %v",
		ConfirmCLI:    "Continue? [Y/N]: ",
		Cancelled:     "Cancelled",
		LoadCfgErr:    "Failed to read config: %v\n",
		SaveCfgErr:    "Failed to save config: %v\n",
		NoGUI: [2]string{
			"This build has no GUI (build tag nogui).",
			"Use CLI mode: repo-cleaner <1-7> [path], or repo-cleaner -h for help.",
		},
	},
}

// Logs 返回指定语言的执行日志文案；非法语言回落中文。
func Logs(l Lang) LogDict {
	if d, ok := logDicts[l]; ok {
		return d
	}
	return logDicts[Zh]
}
