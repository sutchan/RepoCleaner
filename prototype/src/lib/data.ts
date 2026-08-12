// src/lib/data.ts v1.2.1
// 功能数据（对齐 OPENSPEC 7 项功能）与中英文 i18n 字典
// 单一数据来源：页面、README、bat 帮助均从此派生，避免口径不一致

export type Lang = "zh" | "en";

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

export interface ProjectType {
  type: string;
  icon: string;
  size: Record<Lang, string>;
  desc: Record<Lang, string>;
}

export const projectTypes: ProjectType[] = [
  {
    type: "NextJS",
    icon: "▲",
    size: { zh: "≈ 450 MB / 项目", en: "≈ 450 MB / project" },
    desc: {
      zh: "含 .next 缓存，禁用后热更新略慢但省空间。",
      en: "Includes .next cache; disabling trades HMR speed for disk space.",
    },
  },
  {
    type: "React",
    icon: "⚛",
    size: { zh: "≈ 280 MB / 项目", en: "≈ 280 MB / project" },
    desc: {
      zh: "Vite 构建缓存 .vite 占用明显。",
      en: "Vite build cache .vite takes noticeable space.",
    },
  },
  {
    type: "Vue",
    icon: "💚",
    size: { zh: "≈ 310 MB / 项目", en: "≈ 310 MB / project" },
    desc: {
      zh: "含 .nuxt 缓存与 node_modules。",
      en: "Ships with .nuxt cache and node_modules.",
    },
  },
  {
    type: "Node",
    icon: "🟢",
    size: { zh: "≈ 120 MB / 项目", en: "≈ 120 MB / project" },
    desc: {
      zh: "常规后端项目，依赖可全局共享。",
      en: "Plain backend project; deps can be shared globally.",
    },
  },
];

export interface I18nDict {
  appName: string;
  appTagline: string;
  navHome: string;
  navFeatures: string;
  navProjects: string;
  navSteps: string;
  heroBadge: string;
  heroTitle: string;
  heroSubtitle: string;
  ctaPrimary: string;
  ctaSecondary: string;
  statSaved: string;
  statProjects: string;
  statLanguages: string;
  featuresTitle: string;
  featuresSubtitle: string;
  projectsTitle: string;
  projectsSubtitle: string;
  projectsAvg: string;
  stepsTitle: string;
  stepsSubtitle: string;
  stepCopy: string;
  footerDesc: string;
  footerBuilt: string;
  langLabel: string;
}

export const i18n: Record<Lang, I18nDict> = {
  zh: {
    appName: "项目清洁助手",
    appTagline: "RepoCleaner",
    navHome: "首页",
    navFeatures: "功能",
    navProjects: "项目类型",
    navSteps: "使用步骤",
    heroBadge: "一键让所有前端项目保持干净",
    heroTitle: "告别冗余 node_modules，回收 10GB+ 磁盘空间",
    heroSubtitle:
      "自动检测 NextJS / React / Vue / Node 项目，统一全局依赖、清理构建缓存，一次配置永久干净。",
    ctaPrimary: "立即开始",
    ctaSecondary: "查看功能",
    statSaved: "已回收空间",
    statProjects: "适配项目类型",
    statLanguages: "界面语言",
    featuresTitle: "核心功能",
    featuresSubtitle: "7 项功能覆盖依赖管理与缓存清理全流程",
    projectsTitle: "适配项目类型",
    projectsSubtitle: "不同框架的缓存占用一览",
    projectsAvg: "平均缓存占用",
    stepsTitle: "三步开始",
    stepsSubtitle: "把 RepoCleaner.bat 放到项目根目录即可使用",
    stepCopy: "复制命令",
    footerDesc: "开源的 Windows 批处理工具，让前端项目目录保持整洁。",
    footerBuilt: "原型基于 Next.js 构建",
    langLabel: "语言",
  },
  en: {
    appName: "RepoCleaner",
    appTagline: "RepoCleaner",
    navHome: "Home",
    navFeatures: "Features",
    navProjects: "Project Types",
    navSteps: "How to Use",
    heroBadge: "Keep every frontend project clean in one click",
    heroTitle: "Drop redundant node_modules, reclaim 10GB+ disk space",
    heroSubtitle:
      "Auto-detect NextJS / React / Vue / Node projects, unify global deps and clean build caches — configure once, stay clean forever.",
    ctaPrimary: "Get Started",
    ctaSecondary: "View Features",
    statSaved: "Space Reclaimed",
    statProjects: "Project Types",
    statLanguages: "UI Languages",
    featuresTitle: "Core Features",
    featuresSubtitle: "7 features covering dependency management and cache cleanup",
    projectsTitle: "Supported Project Types",
    projectsSubtitle: "Cache footprint by framework",
    projectsAvg: "Avg Cache Size",
    stepsTitle: "Get Started in 3 Steps",
    stepsSubtitle: "Drop RepoCleaner.bat at your project root and run it",
    stepCopy: "Copy",
    footerDesc: "An open-source Windows batch tool that keeps frontend project dirs tidy.",
    footerBuilt: "Prototype built with Next.js",
    langLabel: "Language",
  },
};

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
