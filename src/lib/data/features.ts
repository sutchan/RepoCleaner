// src/lib/data/features.ts v2.3.3
// 功能数据（对齐 OPENSPEC 7 项功能），三语词条齐全
import type { Lang } from "../i18n";

export interface Feature {
  id: number;
  icon: "package" | "rotate-ccw" | "trash" | "rotate-cw" | "download" | "broom" | "zap";
  title: Record<Lang, string>;
  desc: Record<Lang, string>;
}

export const features: Feature[] = [
  {
    id: 1,
    icon: "package",
    title: { zh: "启用全局 node_modules", "zh-Hant": "啟用全域 node_modules", en: "Enable Global node_modules" },
    desc: {
      zh: "配置 NODE_PATH 并为所有项目写入 .npmrc，项目共享全局依赖，永不在本地生成 node_modules。",
      "zh-Hant": "設定 NODE_PATH 並為所有專案寫入 .npmrc，專案共用全域依賴，不在本機產生 node_modules。",
      en: "Set NODE_PATH and write .npmrc for every project so they share global deps and never create local node_modules.",
    },
  },
  {
    id: 2,
    icon: "rotate-ccw",
    title: { zh: "重置 Node 配置", "zh-Hant": "重設 Node 設定", en: "Reset Node Config" },
    desc: {
      zh: "清除 NODE_PATH，删除所有 .npmrc，恢复 npm 默认行为。",
      "zh-Hant": "清除 NODE_PATH、刪除所有 .npmrc，還原 npm 預設行為。",
      en: "Clear NODE_PATH, remove all .npmrc and restore npm default behavior.",
    },
  },
  {
    id: 3,
    icon: "trash",
    title: { zh: "清理并禁用 .next 缓存", "zh-Hant": "清理並停用 .next 快取", en: "Clean and Disable .next Cache" },
    desc: {
      zh: "扫描 Next.js 项目，删除 .next 并写入不可见 .next 文件阻止缓存生成。",
      "zh-Hant": "掃描 Next.js 專案，刪除 .next 並寫入隱藏的 .next 檔案以阻止快取產生。",
      en: "Scan Next.js projects, delete .next and drop an invisible .next file to block cache generation.",
    },
  },
  {
    id: 4,
    icon: "rotate-cw",
    title: { zh: "恢复 Next.js 缓存", "zh-Hant": "還原 Next.js 快取", en: "Restore Next.js Cache" },
    desc: {
      zh: "删除禁止缓存的 .next 文件，恢复 Next.js 正常编译与热更新。",
      "zh-Hant": "刪除禁止快取的 .next 檔案，還原 Next.js 正常編譯與熱更新。",
      en: "Remove the cache-blocking .next file to restore normal Next.js compilation and HMR.",
    },
  },
  {
    id: 5,
    icon: "download",
    title: { zh: "安装常用全局依赖", "zh-Hant": "安裝常用全域依賴", en: "Install Common Global Deps" },
    desc: {
      zh: "批量安装 react/react-dom/next、vue、axios、express、pnpm、yarn 等开发库。",
      "zh-Hant": "批次安裝 react/react-dom/next、vue、axios、express、pnpm、yarn 等開發套件。",
      en: "Batch-install react/react-dom/next, vue, axios, express, pnpm, yarn and more dev libs.",
    },
  },
  {
    id: 6,
    icon: "broom",
    title: { zh: "清理所有项目缓存", "zh-Hant": "清理所有專案快取", en: "Clean All Project Caches" },
    desc: {
      zh: "清理 node_modules/.next/.nuxt/dist/build 及 lock 文件，一键回收磁盘空间。",
      "zh-Hant": "清理 node_modules/.next/.nuxt/dist/build 與 lock 檔案，一次回收磁碟空間。",
      en: "Clean node_modules/.next/.nuxt/dist/build and lockfiles to reclaim disk space in one pass.",
    },
  },
  {
    id: 7,
    icon: "zap",
    title: { zh: "清理 Vite/构建缓存", "zh-Hant": "清理 Vite/建置快取", en: "Clean Vite/Build Caches" },
    desc: {
      zh: "清理 .vite/.turbo/.nuxt/dist/build 等构建缓存，适合 Vite 项目。",
      "zh-Hant": "清理 .vite/.turbo/.nuxt/dist/build 等建置快取，適合 Vite 專案。",
      en: "Clean .vite/.turbo/.nuxt/dist/build build caches; ideal for Vite projects.",
    },
  },
];
