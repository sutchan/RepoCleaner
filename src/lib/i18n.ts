// src/lib/i18n.ts v1.2.2
export type Lang = "zh" | "en";

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

export const i18n: Record<Lang, Dict> = {
  zh: {
    appTitle: "RepoCleaner 项目清洁助手",
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
    footer: "RepoCleaner 跨平台仓库清理工具",
  },
  en: {
    appTitle: "RepoCleaner",
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
    footer: "RepoCleaner cross-platform repo cleaner",
  },
};
