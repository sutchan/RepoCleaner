// src/lib/data/features.ts v1.2.2
// 功能数据（对齐 OPENSPEC 7 项功能）
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
    title: { zh: "启用全局 node_modules", en: "Enable Global node_modules" },
    desc: {
      zh: "配置 NODE_PATH 并为所有项目写入 .npmrc，项目共享全局依赖，永不在本地生成 node_modules。",
      en: "Set NODE_PATH and write .npmrc for every project so they share global deps and never create local node_modules.",
    },
  },
  {
    id: 2,
    icon: "rotate-ccw",
    title: { zh: "还原默认 Node 配置", en: "Restore Default Node Config" },
    desc: {
      zh: "清除 NODE_PATH，删除所有 .npmrc，恢复 npm 默认行为。",
      en: "Clear NODE_PATH, remove all .npmrc and restore npm default behavior.",
    },
  },
  {
    id: 3,
    icon: "trash",
    title: { zh: "清理 & 禁用 .next 缓存", en: "Clean & Disable .next Cache" },
    desc: {
      zh: "扫描 NextJS 项目，删除 .next 并写入不可见 .next 文件阻止缓存生成。",
      en: "Scan NextJS projects, delete .next and drop an invisible .next file to block cache generation.",
    },
  },
  {
    id: 4,
    icon: "rotate-cw",
    title: { zh: "还原 NextJS 缓存", en: "Restore NextJS Cache" },
    desc: {
      zh: "删除禁止缓存的 .next 文件，恢复 NextJS 正常编译与热更新。",
      en: "Remove the cache-blocking .next file to restore normal NextJS compilation and HMR.",
    },
  },
  {
    id: 5,
    icon: "download",
    title: { zh: "一键安装常用全局依赖", en: "Install Common Global Deps" },
    desc: {
      zh: "批量安装 react/react-dom/next、vue、axios、express、pnpm、yarn 等开发库。",
      en: "Batch-install react/react-dom/next, vue, axios, express, pnpm, yarn and more dev libs.",
    },
  },
  {
    id: 6,
    icon: "broom",
    title: { zh: "清理所有项目缓存", en: "Clean All Project Caches" },
    desc: {
      zh: "清理 node_modules/.next/.nuxt/dist/build 及 lock 文件，一键回收磁盘空间。",
      en: "Clean node_modules/.next/.nuxt/dist/build and lockfiles to reclaim disk space in one pass.",
    },
  },
  {
    id: 7,
    icon: "zap",
    title: { zh: "清理 Vite/构建缓存", en: "Clean Vite/Build Caches" },
    desc: {
      zh: "清理 .vite/.turbo/.nuxt/dist/build 等构建缓存，适合 Vite 项目。",
      en: "Clean .vite/.turbo/.nuxt/dist/build build caches; ideal for Vite projects.",
    },
  },
];
