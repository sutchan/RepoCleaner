# RepoCleaner 原型与设计规范 (prototype) v1.2.3

本目录仅保存 RepoCleaner 的**原型参考与设计规范文件**，不再包含可运行的应用代码。

> 可运行的应用代码已提升至仓库根目录（`src/` + 根级 `package.json` / `next.config.ts` / `tsconfig.json` 等）。
> 根目录应用即 OPENSPEC 规范的可视化落地，与 `RepoCleaner.bat` 的 7 项功能对齐。

## 本目录内容

| 文件 | 说明 |
|------|------|
| `DESIGN-SYSTEM.md` | 唯一设计规范来源（色彩/字体/间距/圆角/动效/组件库/交互/a11y） |
| `tokens.json` | 设计令牌（与 `DESIGN-SYSTEM.md` §2 及 `src/app/globals.css` 的 CSS 变量一致） |

## 设计规范要点

- 设计令牌：颜色/间距/字号全部走 `src/app/globals.css` 的 CSS 变量，**禁止散落硬编码**（页面与组件只使用 `bg-background` / `text-foreground` / `text-muted-foreground` 等令牌工具类）。
- 字体：系统字体栈（`-apple-system, "Segoe UI", "PingFang SC", "Microsoft YaHei", sans-serif`），不引入外部网络字体。
- 暗色优先（开发者工具调性），`color-scheme: dark`。
- 所有 UI 容器/区块带语义化 `id`（如 `repo-cleaner-app`、`main-tabs`、`repo-card-<name>`），便于调试定位与无障碍访问。

## 根目录应用开发

```bash
npm install
npm run dev        # 开发预览
npm run type-check # 类型检查（tsc --noEmit）
npm run build      # 生产构建
```
