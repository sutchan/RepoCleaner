// src/lib/data/steps.ts v2.3.3
// 使用步骤（三步开始）；描述以当前 Go 产品 repo-manager 为准
import type { Lang } from "../i18n";

export const steps: { title: Record<Lang, string>; desc: Record<Lang, string> }[] = [
  {
    title: { zh: "取得程序", "zh-Hant": "取得程式", en: "Get the binary" },
    desc: {
      zh: "从 Release 下载对应平台的归档，或用 Go 构建 repo-manager。",
      "zh-Hant": "從 Release 下載對應平台的歸檔，或用 Go 建置 repo-manager。",
      en: "Download the archive for your platform from Releases, or build repo-manager with Go.",
    },
  },
  {
    title: { zh: "选择目录", "zh-Hant": "選擇目錄", en: "Pick a directory" },
    desc: {
      zh: "GUI 选择目录或 CLI 传入路径；config.ini 可保存默认目录。",
      "zh-Hant": "GUI 選擇目錄或 CLI 傳入路徑；config.ini 可儲存預設目錄。",
      en: "Choose a folder in the GUI or pass a CLI path; config.ini stores the default.",
    },
  },
  {
    title: { zh: "执行功能", "zh-Hant": "執行功能", en: "Run a feature" },
    desc: {
      zh: "点击功能按钮 1–7，或用 repo-manager <1-7> [路径]；危险操作会二次确认。",
      "zh-Hant": "點擊功能按鈕 1–7，或用 repo-manager <1-7> [路徑]；危險操作會二次確認。",
      en: "Click feature buttons 1–7, or use repo-manager <1-7> [path]; destructive actions require confirmation.",
    },
  },
];
