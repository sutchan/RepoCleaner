// src/lib/data/steps.ts v1.2.2
// 使用步骤（三步开始）
import type { Lang } from "../i18n";

export const steps: { title: Record<Lang, string>; desc: Record<Lang, string> }[] = [
  {
    title: { zh: "放置脚本", en: "Place the script" },
    desc: {
      zh: "将 RepoCleaner.bat 复制到仓库根目录。",
      en: "Copy RepoCleaner.bat to your repo root.",
    },
  },
  {
    title: { zh: "选择目录", en: "Pick a directory" },
    desc: {
      zh: "支持自动检测、手动输入或读取 config.ini。",
      en: "Auto-detect, type a path, or load config.ini.",
    },
  },
  {
    title: { zh: "运行功能", en: "Run a feature" },
    desc: {
      zh: "按数字菜单或命令行参数执行对应功能。",
      en: "Use the numbered menu or a CLI arg to run a feature.",
    },
  },
];
