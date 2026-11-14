// landing/assets/app.js v2.3.3
// 三语词条（简中 → 繁中 → 英文）与语言切换；偏好持久化并做枚举校验。
const SUPPORTED = ["zh", "zh-Hant", "en"];
const HTML_LANG = { zh: "zh", "zh-Hant": "zh-Hant", en: "en" };
const STORAGE_KEY = "repo-manager-landing-lang";

const dict = {
  zh: {
    navFeatures: "功能", navWorkflow: "使用步骤", navFaq: "常见问题",
    heroEyebrow: "跨平台 · 本地优先", heroTitle: "让前端项目保持干净",
    heroDesc: "7 项功能统一管理全局 node_modules 与前端项目缓存；危险操作执行前二次确认，支持预演模式与执行日志。",
    ctaDownload: "下载 Release", ctaDocs: "阅读文档",
    fact1: "7 项功能", fact2: "简中 / 繁中 / 英文", fact3: "Windows / macOS / Linux",
    featuresTitle: "七项功能", featuresLead: "功能编号与命令行参数一致，可直接调用。",
    f1: "启用全局 node_modules", d1: "配置 NODE_PATH 并清理项目内本地依赖。",
    f2: "重置 Node 配置", d2: "移除环境变量与项目 .npmrc。",
    f3: "清理 .next 缓存", d3: "写入同名文件以停止 Next.js 缓存写入。",
    f4: "恢复 Next.js 缓存", d4: "清除隐藏属性并删除占位文件。",
    f5: "安装全局依赖", d5: "分五组安装常用前端工具链。",
    f6: "清理所有项目缓存", d6: "按白名单清理缓存目录与本地依赖。",
    f7: "清理 Vite / 构建缓存", d7: "清理 .vite / .turbo / .angular 等目录。",
    workflowTitle: "三步开始",
    w1t: "取得程序", w1d: "从 Release 下载对应平台归档，或用 Go 构建 repo-manager。",
    w2t: "选择目录", w2d: "GUI 选择目录或 CLI 传入路径，config.ini 可保存默认目录。",
    w3t: "执行功能", w3d: "点击功能按钮 1–7，或用 repo-manager <1-7> [路径]。",
    safetyTitle: "安全设计",
    s1: "危险功能（1 / 6 / 7）执行前二次确认。",
    s2: "预演模式只显示将执行的操作，不修改文件。",
    s3: "递归扫描限制为 3 层，自定义清理目录名经安全校验。",
    s4: "执行日志同步写入可执行文件同目录的日志文件。",
    faqTitle: "常见问题",
    q1: "是否需要联网或上传项目数据？",
    a1: "不需要。工具在本地运行，不上传仓库内容；联网仅发生在功能 5 安装全局依赖时。",
    q2: "支持哪些操作系统？",
    a2: "Go 版支持 Windows、macOS 与 Linux 单二进制；GUI 运行需要系统 OpenGL，纯命令行可用 nogui 构建。",
    q3: "界面语言可以切换吗？",
    a3: "支持简体中文、繁体中文与英文；CLI 帮助固定为英文，不随语言切换。",
    footerNote: "开源 · MIT 许可 · 内容以仓库文档与 Release 说明为准",
    versionMissing: "版本信息待构建生成",
  },
  "zh-Hant": {
    navFeatures: "功能", navWorkflow: "使用步驟", navFaq: "常見問題",
    heroEyebrow: "跨平台 · 本機優先", heroTitle: "讓前端專案保持乾淨",
    heroDesc: "7 項功能統一管理全域 node_modules 與前端專案快取；危險操作執行前二次確認，支援預演模式與執行紀錄。",
    ctaDownload: "下載 Release", ctaDocs: "閱讀文件",
    fact1: "7 項功能", fact2: "簡中 / 繁中 / 英文", fact3: "Windows / macOS / Linux",
    featuresTitle: "七項功能", featuresLead: "功能編號與命令列參數一致，可直接呼叫。",
    f1: "啟用全域 node_modules", d1: "設定 NODE_PATH 並清理專案內本機依賴。",
    f2: "重設 Node 設定", d2: "移除環境變數與專案 .npmrc。",
    f3: "清理 .next 快取", d3: "寫入同名檔案以停止 Next.js 快取寫入。",
    f4: "還原 Next.js 快取", d4: "清除隱藏屬性並刪除佔位檔案。",
    f5: "安裝全域依賴", d5: "分五組安裝常用前端工具鏈。",
    f6: "清理所有專案快取", d6: "依白名單清理快取目錄與本機依賴。",
    f7: "清理 Vite / 建置快取", d7: "清理 .vite / .turbo / .angular 等目錄。",
    workflowTitle: "三步開始",
    w1t: "取得程式", w1d: "從 Release 下載對應平台歸檔，或用 Go 建置 repo-manager。",
    w2t: "選擇目錄", w2d: "GUI 選擇目錄或 CLI 傳入路徑，config.ini 可儲存預設目錄。",
    w3t: "執行功能", w3d: "點擊功能按鈕 1–7，或用 repo-manager <1-7> [路徑]。",
    safetyTitle: "安全設計",
    s1: "危險功能（1 / 6 / 7）執行前二次確認。",
    s2: "預演模式只顯示將執行的操作，不修改檔案。",
    s3: "遞迴掃描限制為 3 層，自訂清理目錄名稱經安全校驗。",
    s4: "執行紀錄同步寫入可執行檔同目錄的紀錄檔。",
    faqTitle: "常見問題",
    q1: "是否需要連網或上傳專案資料？",
    a1: "不需要。工具在本機執行，不上傳儲存庫內容；連網僅發生在功能 5 安裝全域依賴時。",
    q2: "支援哪些作業系統？",
    a2: "Go 版支援 Windows、macOS 與 Linux 單一執行檔；GUI 需要系統 OpenGL，純命令列可用 nogui 建置。",
    q3: "介面語言可以切換嗎？",
    a3: "支援簡體中文、繁體中文與英文；CLI 說明固定為英文，不隨語言切換。",
    footerNote: "開源 · MIT 授權 · 內容以儲存庫文件與 Release 說明為準",
    versionMissing: "版本資訊待建置產生",
  },
  en: {
    navFeatures: "Features", navWorkflow: "How it works", navFaq: "FAQ",
    heroEyebrow: "Cross-platform · local first", heroTitle: "Keep your front-end projects clean",
    heroDesc: "Seven functions manage global node_modules and front-end caches. Destructive actions require confirmation; dry-run mode and execution logs are built in.",
    ctaDownload: "Download Release", ctaDocs: "Read the docs",
    fact1: "7 functions", fact2: "Simplified / Traditional Chinese / English", fact3: "Windows / macOS / Linux",
    featuresTitle: "Seven functions", featuresLead: "Function numbers match the CLI arguments.",
    f1: "Enable global node_modules", d1: "Configure NODE_PATH and remove local project dependencies.",
    f2: "Reset Node config", d2: "Remove the environment variable and project .npmrc files.",
    f3: "Clean .next cache", d3: "Write a same-name file to stop Next.js cache writes.",
    f4: "Restore Next.js cache", d4: "Clear the hidden attribute and remove the placeholder file.",
    f5: "Install global dependencies", d5: "Install common front-end tooling in five groups.",
    f6: "Clean all project caches", d6: "Clean cache directories and local dependencies by allowlist.",
    f7: "Clean Vite / build caches", d7: "Clean .vite / .turbo / .angular and similar directories.",
    workflowTitle: "Three steps to start",
    w1t: "Get the binary", w1d: "Download the archive for your platform from Releases, or build repo-manager with Go.",
    w2t: "Pick a directory", w2d: "Choose a folder in the GUI or pass a CLI path; config.ini stores the default.",
    w3t: "Run a feature", w3d: "Click feature buttons 1–7, or use repo-manager <1-7> [path].",
    safetyTitle: "Safety by design",
    s1: "Functions 1 / 6 / 7 require a second confirmation.",
    s2: "Dry-run mode only shows planned actions and changes nothing.",
    s3: "Recursive scan is limited to depth 3; custom clean directory names are validated.",
    s4: "Execution logs are written next to the executable.",
    faqTitle: "Frequently asked questions",
    q1: "Does it need network access or upload my project data?",
    a1: "No. It runs locally and never uploads repository content; network access is only used by function 5 to install global dependencies.",
    q2: "Which operating systems are supported?",
    a2: "The Go build ships a single binary for Windows, macOS and Linux. The GUI needs system OpenGL; a nogui build works from the CLI only.",
    q3: "Can I change the interface language?",
    a3: "Simplified Chinese, Traditional Chinese and English are supported. CLI help stays in English regardless of the language setting.",
    footerNote: "Open source · MIT license · Facts follow the repository docs and Release notes",
    versionMissing: "Version information pending build",
  },
};

function readLang() {
  try {
    const saved = localStorage.getItem(STORAGE_KEY);
    return SUPPORTED.includes(saved) ? saved : "zh";
  } catch {
    return "zh";
  }
}

function applyLang(lang) {
  const chosen = SUPPORTED.includes(lang) ? lang : "zh";
  document.documentElement.lang = HTML_LANG[chosen];
  document.querySelectorAll("[data-i18n]").forEach((el) => {
    const value = dict[chosen][el.dataset.i18n];
    if (value) el.textContent = value;
  });
  document.querySelectorAll("#landing-lang-switch button").forEach((btn) => {
    const active = btn.dataset.lang === chosen;
    btn.setAttribute("aria-pressed", String(active));
    btn.classList.toggle("is-active", active);
  });
  // 版本与 build 编号来自构建期生成的 version.js，页面不现场生成
  const label = window.LANDING_VERSION && window.LANDING_VERSION.label;
  document.getElementById("landing-version").textContent = label || dict[chosen].versionMissing;
  try {
    localStorage.setItem(STORAGE_KEY, chosen);
  } catch {
    // 隐私模式下无法写入偏好时静默降级，保持页面可用
  }
}

document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll("#landing-lang-switch button").forEach((btn) => {
    btn.addEventListener("click", () => applyLang(btn.dataset.lang));
  });
  applyLang(readLang());
});
