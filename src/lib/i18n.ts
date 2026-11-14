// src/lib/i18n.ts v2.3.3
export type Lang = "zh" | "zh-Hant" | "en";

export type Dict = {
  appTitle: string;
  language: string;
  statTotal: string;
  statSafe: string;
  statTweaks: string;
  securityTitle: string;
  securityDesc: string;
  tabScan: string;
  tabTweaks: string;
  safe: string;
  warning: string;
  clean: string;
  footer: string;
};

// 语言优先级：简中 → 繁中 → 英文；支持列表供 localStorage 枚举校验。
export const SUPPORTED_LANGS: Lang[] = ["zh", "zh-Hant", "en"];

// html lang 取值：简中 zh、繁中 zh-Hant、英文 en。
export const HTML_LANG: Record<Lang, string> = {
  zh: "zh",
  "zh-Hant": "zh-Hant",
  en: "en",
};

export const i18n: Record<Lang, Dict> = {
  zh: {
    appTitle: "Repo Manager",
    language: "语言",
    statTotal: "扫描仓库",
    statSafe: "安全项",
    statTweaks: "优化项",
    securityTitle: "安全提示",
    securityDesc: "删除操作不可逆，清理前请确认已备份重要文件。",
    tabScan: "扫描结果",
    tabTweaks: "优化建议",
    safe: "安全",
    warning: "需注意",
    clean: "清理",
    footer: "Repo Manager",
  },
  "zh-Hant": {
    appTitle: "Repo Manager",
    language: "語言",
    statTotal: "掃描儲存庫",
    statSafe: "安全項目",
    statTweaks: "優化項目",
    securityTitle: "安全提示",
    securityDesc: "刪除作業無法復原，清理前請確認已備份重要檔案。",
    tabScan: "掃描結果",
    tabTweaks: "優化建議",
    safe: "安全",
    warning: "需注意",
    clean: "清理",
    footer: "Repo Manager",
  },
  en: {
    appTitle: "Repo Manager",
    language: "Language",
    statTotal: "Scanned",
    statSafe: "Safe",
    statTweaks: "Tweaks",
    securityTitle: "Security Notice",
    securityDesc: "Deletion is irreversible. Back up important files before cleaning.",
    tabScan: "Scan Results",
    tabTweaks: "Optimizations",
    safe: "Safe",
    warning: "Caution",
    clean: "Clean",
    footer: "Repo Manager",
  },
};
