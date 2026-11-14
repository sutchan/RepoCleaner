// src/lib/data/project-types.ts v2.3.3
// 适配项目类型说明；不写入未经核实的容量数字（数据真实性红线）
import type { Lang } from "../i18n";

export interface ProjectType {
  type: string;
  icon: string;
  size: Record<Lang, string>;
  desc: Record<Lang, string>;
}

export const projectTypes: ProjectType[] = [
  {
    type: "Next.js",
    icon: "▲",
    size: { zh: "依项目而定", "zh-Hant": "依專案而定", en: "Varies by project" },
    desc: {
      zh: "含 .next 缓存；功能 3 写入同名文件以停止缓存写入。",
      "zh-Hant": "含 .next 快取；功能 3 寫入同名檔案以停止快取寫入。",
      en: "Includes .next cache; feature 3 writes a same-name file to stop cache writes.",
    },
  },
  {
    type: "React",
    icon: "⚛",
    size: { zh: "依项目而定", "zh-Hant": "依專案而定", en: "Varies by project" },
    desc: {
      zh: "Vite 构建缓存目录 .vite 属于功能 7 的清理白名单。",
      "zh-Hant": "Vite 建置快取目錄 .vite 屬於功能 7 的清理白名單。",
      en: "The .vite build cache directory is in feature 7's allowlist.",
    },
  },
  {
    type: "Vue",
    icon: "💚",
    size: { zh: "依项目而定", "zh-Hant": "依專案而定", en: "Varies by project" },
    desc: {
      zh: "项目目录内的 node_modules 与构建缓存按白名单清理。",
      "zh-Hant": "專案目錄內的 node_modules 與建置快取依白名單清理。",
      en: "node_modules and build caches inside the project are cleaned by allowlist.",
    },
  },
  {
    type: "Node",
    icon: "🟢",
    size: { zh: "依项目而定", "zh-Hant": "依專案而定", en: "Varies by project" },
    desc: {
      zh: "常规后端项目；功能 1 启用全局 node_modules 共享依赖。",
      "zh-Hant": "一般後端專案；功能 1 啟用全域 node_modules 共用依賴。",
      en: "Plain backend project; feature 1 enables shared global node_modules.",
    },
  },
];
