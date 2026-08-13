// src/lib/data/repos.ts v1.2.2
// 扫描结果与优化建议（prototype dashboard 数据源）

export type Risk = "safe" | "warning";

export interface ScanResult {
  repo: string;
  path: string;
  risk: Risk;
}

export const scanResults: ScanResult[] = [
  { repo: "next-dashboard", path: "E:\\Github\\next-dashboard", risk: "safe" },
  { repo: "vue-admin", path: "E:\\Github\\vue-admin", risk: "warning" },
  { repo: "react-portfolio", path: "E:\\Github\\react-portfolio", risk: "safe" },
  { repo: "node-api", path: "E:\\Github\\node-api", risk: "safe" },
];

export interface Tweak {
  id: string;
  category: string;
  title: string;
  desc: string;
}

export const tweaks: Tweak[] = [
  {
    id: "global-node",
    category: "依赖",
    title: "启用全局 node_modules",
    desc: "为全部仓库写入 .npmrc 并配置 NODE_PATH，预计回收约 1.2 GB。",
  },
  {
    id: "next-cache",
    category: "缓存",
    title: "清理 .next 构建缓存",
    desc: "2 个 NextJS 项目共占用 860 MB，禁用后热更新略慢但省空间。",
  },
  {
    id: "vite-cache",
    category: "缓存",
    title: "清理 Vite 构建缓存",
    desc: "3 个 Vite 项目 .vite 目录共占用 540 MB。",
  },
  {
    id: "lockfiles",
    category: "依赖",
    title: "移除 lock 文件",
    desc: "6 个 package-lock.json / yarn.lock，共 4.2 MB。",
  },
];
