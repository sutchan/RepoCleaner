# RepoCleaner Prototype v1.2.2

RepoCleaner 的参考界面（prototype），是 [OPENSPEC.md](../OPENSPEC.md) 规范的可视化落地，与 `RepoCleaner.bat` 的 7 项功能对齐。

## 技术栈

- Next.js 16（App Router）+ React 19 + TypeScript（strict）
- Tailwind CSS v4（CSS 变量令牌，见 `DESIGN-SYSTEM.md` / `tokens.json`）
- 零第三方 UI/图标依赖（组件自研 + 内联 SVG）

## 目录结构

```
src/
├── app/
│   ├── layout.tsx        # 根布局
│   ├── page.tsx          # 单页 dashboard（概览/安全提示/扫描结果/优化建议）
│   └── globals.css       # Tailwind 4 主题令牌
├── components/
│   ├── icons.tsx         # 内联 SVG 图标（ScanLine/ShieldCheck/...）
│   └── ui/               # 基础组件：button/card/alert/badge/tabs/table/code-block
└── lib/
    ├── i18n.ts           # 中/英文案字典（唯一 i18n 源）
    └── data/             # 业务数据（统一经 @/lib/data 导出）
        ├── features.ts       # 7 项功能
        ├── project-types.ts  # 适配项目类型
        ├── steps.ts          # 使用步骤
        ├── repos.ts          # 扫描结果 + 优化建议（dashboard 数据源）
        └── index.ts          # barrel 出口
```

## 数据流

- 页面文案 → `src/lib/i18n.ts`（`useState<Lang>` 实时切换）
- 展示数据 → `src/lib/data/*`（单一数据源，`@/lib/data` 统一导入）

## 开发

```bash
npm install
npm run dev        # 开发预览
npm run type-check # 类型检查（tsc --noEmit）
npm run build      # 生产构建
```

## 设计规范

- 设计令牌：`DESIGN-SYSTEM.md` + `tokens.json`（与 `globals.css` 的 CSS 变量一致）
- 所有 UI 容器/区块带语义化 `id`（如 `repo-cleaner-app`、`main-tabs`、`repo-card-<name>`）
