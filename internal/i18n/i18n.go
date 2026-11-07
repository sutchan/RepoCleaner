// internal/i18n/i18n.go v2.2.0

// Package i18n 是唯一的 UI 文案来源（zh / en），禁止在 ui / engine 层硬编码文案。
package i18n

// Lang 界面语言。
type Lang string

// 支持的语言（对齐 OPENSPEC §4.1）。
const (
	Zh Lang = "zh"
	En Lang = "en"
)

// Normalize 校验并归一化语言值，非法值回落中文。
func Normalize(l string) Lang {
	if Lang(l) == En {
		return En
	}
	return Zh
}

// Dict 单语言界面文案；功能名称顺序与 OPENSPEC §3.0 一致（下标 0 对应功能 1）。
// 执行日志文案见 log.go 的 Logs()。
type Dict struct {
	AppTitle       string
	TargetDir      string
	ChooseDir      string
	Save           string
	Saved          string
	LangLabel      string
	FuncNames      [7]string
	RunHint        string
	DryRunLabel    string
	WarningF1      string
	WarningAll     string
	WarningVite    string
	Confirm        string
	Ok             string
	ErrInvalidRoot string
	About          string
	AboutText      string
	OutputTitle    string
}

var dicts = map[Lang]Dict{
	Zh: {
		AppTitle:       "RepoCleaner - 项目清洁助手",
		TargetDir:      "目标目录:",
		ChooseDir:      "选择目录",
		Save:           "保存配置",
		Saved:          "配置已保存到 config.ini",
		LangLabel:      "语言",
		FuncNames:      [7]string{"启用全局 node_modules", "重置 Node 配置", "清理 .next 缓存", "恢复 Next.js 缓存", "安装全局依赖", "清理所有项目缓存", "清理 Vite/构建缓存"},
		RunHint:        "点击功能按钮执行；危险操作会先确认",
		DryRunLabel:    "预演模式 (dry-run)",
		WarningF1:      "警告: 这将删除所有项目的本地 node_modules!",
		WarningAll:     "警告: 这将永久删除所有缓存文件!",
		WarningVite:    "警告: 这将删除 .vite/.turbo/.angular 等构建缓存目录!",
		Confirm:        "是否继续?",
		Ok:             "[完成]",
		ErrInvalidRoot: "错误: 路径不存在!",
		About:          "关于",
		AboutText:      "RepoCleaner 通过 7 项功能管理全局依赖与前端项目缓存；危险操作执行前会二次确认。版本与构建信息见窗口页脚。",
		OutputTitle:    "执行输出",
	},
	En: {
		AppTitle:       "RepoCleaner - Project Cleaner",
		TargetDir:      "Target:",
		ChooseDir:      "Choose directory",
		Save:           "Save config",
		Saved:          "Config saved to config.ini",
		LangLabel:      "Language",
		FuncNames:      [7]string{"Enable Global node_modules", "Reset Node Config", "Clean .next Cache", "Restore Next.js Cache", "Install Global Dependencies", "Clean All Project Cache", "Clean Vite/Build Cache"},
		RunHint:        "Click a function button; dangerous ones ask for confirmation",
		DryRunLabel:    "Dry-run mode",
		WarningF1:      "WARNING: This will delete all local node_modules!",
		WarningAll:     "WARNING: This will permanently delete all cache files!",
		WarningVite:    "WARNING: This will delete .vite/.turbo/.angular and other build cache directories!",
		Confirm:        "Continue?",
		Ok:             "[OK]",
		ErrInvalidRoot: "ERROR: Path does not exist!",
		About:          "About",
		AboutText:      "RepoCleaner manages global dependencies and front-end caches with 7 functions; destructive actions require confirmation. Version and build info are shown in the window footer.",
		OutputTitle:    "Output",
	},
}

// T 返回指定语言界面文案；非法语言回落中文。
func T(l Lang) Dict {
	if d, ok := dicts[l]; ok {
		return d
	}
	return dicts[Zh]
}
