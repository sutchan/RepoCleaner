// src/lib/data/project-types.ts v1.2.2
// 适配项目类型与平均缓存占用
import type { Lang } from "../i18n";

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
