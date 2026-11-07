# RepoCleaner 设计系统 (Design System) v2.1.0

> **本文定位**：RepoCleaner 的**唯一 UI 设计规范来源**，规定设计令牌、界面结构与交互标准。
> **v2.0.0 起**：实际 UI 为 **Go + Fyne 桌面 GUI**（`internal/ui`），配色沿用本文 §2.1 设计令牌；仓库根 `src/` 的 Next.js 参考界面为 **v1.x 遗留、已弃用**，§3.1 / §3.2 等 Next.js 章节仅作历史参考。
> 本文归档于 `docs/`（v1.2.5 起从 `prototype/` 迁入）；`prototype/` 现仅保留原型说明，不含可运行代码或规范正文。
> 机器可读令牌见 [tokens.json](./tokens.json)；开发规范见 [CONTRIBUTING.md](./CONTRIBUTING.md)；文档导航见 [README.md](./README.md)。

**阅读约定**：本文用「**已实现**」标注当前代码确实具备的能力，用「**规划**」标注尚未落地的设计意图。凡「规划」项均不得在文档其它位置被当作既成事实引用。完整差距清单见 §7。

---

## 1. 设计原则

1. **极简优先** — 留白即设计，单屏只解决一件事。
2. **数据可追溯** — 展示数据统一来自 `src/lib/data/`；其中 `repos.ts` 为**演示数据（mock）**，已在文件头注释标明，容量数字为示意，不可作为事实引用。
3. **一致令牌** — 颜色/字体全部走 `src/app/globals.css` 的 `@theme` CSS 变量生成的工具类，禁止散落硬编码颜色（`bg-background` / `text-foreground` / `text-muted-foreground` / `bg-muted` 等）。
4. **响应式无妥协** — 移动端单列、桌面多列，断点 `sm/md/lg/xl`。
5. **克制动效** — 仅用于引导注意力与状态反馈，时长 ≤ 300ms。

---

## 2. 设计令牌 (Design Tokens) — 已实现

### 2.1 色彩

中性灰 + 单一品牌强调色（青绿），**当前仅实现暗色主题**（`:root { color-scheme: dark }`）。

Tailwind v4 的 `@theme` 使用 `--color-<name>` 命名，生成 `bg-<name>` / `text-<name>` / `border-<name>` 工具类。下表「令牌」列写逻辑名，「CSS 变量」列写实现名。

| 令牌 | 亮色（规划） | 暗色（已实现） | CSS 变量 / 工具类 | 用途 |
|------|-------------|---------------|-------------------|------|
| background | `#ffffff` | `#0a0a0b` | `--color-background` / `bg-background` | 页面底色 |
| foreground | `#0a0a0b` | `#ededed` | `--color-foreground` / `text-foreground` | 主文字 |
| card | `#ffffff` | `#161618` | `--color-card` / `bg-card` | 卡片底 |
| muted | `#f4f4f5` | `#1c1c1f` | `--color-muted` / `bg-muted` | 次级底 |
| muted-foreground | `#71717a` | `#a1a1aa` | `--color-muted-foreground` / `text-muted-foreground` | 次级文字 |
| border | `#e4e4e7` | `#27272a` | `--color-border` / `border-border` | 描边 |
| primary | `#10b981` | `#10b981` | `--color-primary` / `bg-primary` | 品牌主色（emerald-500） |
| primary-foreground | `#ffffff` | `#052e22` | `--color-primary-foreground` | 主色上文字 |
| destructive | `#ef4444` | `#ef4444` | `--color-destructive` | 危险 / 删除 |
| success | `#22c55e` | `#22c55e` | `--color-success` | 成功 |
| warning | `#f59e0b` | `#f59e0b` | `--color-warning` | 警告 |

**语义规则（已实现）**：删除/风险类用 `destructive`，安全/成功用 `success`，提示用 `warning`，主操作用 `primary`。

> v2.0.0 起 Fyne 版按上表在 `internal/ui/theme.go` 实现同名配色（背景 / 前景 / 主色 / 危险等）；亮暗主题随 Fyne 系统设置，`tokens.json` 仍为令牌契约。

> 亮色列的值保存在 `tokens.json` 中作为规划基线，但 `globals.css` **未**实现亮色模式切换；不得假定亮色主题可用。

### 2.2 字体 — 已实现

| 层级 | 字号 | 字重 | 行高 | 用途 |
|------|------|------|------|------|
| Display | 48px | 700 | 1.1 | 首页大标题（规划，当前页面未使用） |
| H1 | 32px | 700 | 1.2 | 页面标题（当前用 `text-2xl`） |
| H2 | 24px | 600 | 1.3 | 区块标题 |
| H3 | 18px | 600 | 1.4 | 卡片标题 |
| Body | 15px | 400 | 1.6 | 正文（`body` 已设置） |
| Small | 13px | 400 | 1.5 | 辅助文字（`text-sm` / `text-xs`） |
| Code | 13px | 500 | 1.5 | 等宽（`font-mono`） |

字体栈（已实现，写入 `--font-sans` / `--font-mono`）：

```
sans: -apple-system, "Segoe UI", "PingFang SC", "Microsoft YaHei", sans-serif
mono: ui-monospace, "JetBrains Mono", "SF Mono", Consolas, monospace
```

**禁止**引入 `next/font/google` 等构建期联网资源：离线/沙箱环境会导致 `next build` 失败。布局使用 `font-sans antialiased`。

### 2.3 间距 — 部分实现

4px 基准栅格。`tokens.json` 定义 `space-1`(4) ～ `space-8`(32)，作为设计契约；**`globals.css` 未逐项定义 `--space-*` 变量**，当前直接使用 Tailwind 内置间距工具类：

- 组件内边距：`p-4` / `p-6`；列表项 `py-4`
- 卡片间距：`gap-4`；区块间距 `space-y-3` / `gap-2`
- 页面容器：`px-4 py-8`

新增间距优先复用上述既有档位，避免任意值（`p-[13px]`）。

### 2.4 圆角与阴影 — 部分实现

- `globals.css` 已定义 `--radius: 12px`（卡片）、`--radius-sm: 8px`（按钮），徽章用 `rounded-full`。
- **当前组件使用具名数值类**（`rounded-[12px]` / `rounded-[8px]` / `rounded-[10px]` / `rounded-[6px]`），未消费 `--radius` 变量。统一改造列入 §7 待办。
- 阴影：暗色下以 `border` 代替阴影；仅激活态 Tab 用 `shadow-sm`。

### 2.5 图标 — 已实现

- 实现：[`src/components/icons.tsx`](../src/components/icons.tsx) 内联 SVG，**零外部图标依赖**。
- 规格：`viewBox="0 0 24 24"`、`fill="none"`、`stroke="currentColor"`、`strokeWidth={2}`、`strokeLinecap/Linejoin="round"`，默认 `width/height=16`，`aria-hidden="true"`。
- 现有图标：`ScanLine`（扫描）、`ShieldCheck`（安全）、`FolderTree`（目录）、`CheckCircle2`（成功）、`AlertCircle`（警告）、`Trash2`（清理）。
- 尺寸由调用方用 `className="h-4 w-4"` 等覆盖；颜色继承 `currentColor`。

### 2.6 动效 — 已实现（部分未消费）

- `globals.css` 定义 `--dur-fast: 150ms`、`--dur: 250ms`，缓动 `cubic-bezier(0.4,0,0.2,1)`。
- 已定义关键帧：`fade-in-up`（卡片入场，`.animate-fade-in-up`）、`spin`（旋转加载，`.animate-spin-slow`）——**当前页面未使用**。
- 组件过渡实际使用 Tailwind `transition-all duration-150` / `transition-colors`；按钮 `active:scale-[0.98]`。
- 焦点态：`:focus-visible { outline: 2px solid var(--color-primary); outline-offset: 2px }`。
- 文本选中：`::selection` 用 `color-mix` 混合 `primary` 35%。

---

## 3. 组件库

### 3.0 Fyne 界面（v2.0.0 已实现，产品本体 UI）

Go 版 GUI 位于 `internal/ui`，遵循本文 §2.1 配色与 §4 交互标准：

| 区块 | 说明 |
|------|------|
| 主窗口 | 标题 + 目标目录展示（可修改并保存到 `config.ini`） |
| 功能按钮 1–7 | 与 `OPENSPEC.md` §3.0 功能一一对应；危险操作（1 / 6 / 7）先弹确认对话框 |
| 语言按钮 | 「中文 / English」实时切换界面与执行日志文案（`internal/i18n` 的 `T()` / `Logs()`） |
| 关于按钮 | 弹窗展示功能说明，不出现版本号（版本信息仅页脚） |
| 结果区 | 展示每个功能的执行日志与成功消息 |
| 预演复选框 | 「预演模式 (dry-run)」：仅输出将要执行的操作，不修改任何文件；会话级选择，不写入 `config.ini`（见 `OPENSPEC.md` §4.7） |
| 页脚区（Border bottom） | 唯一展示 `vX.Y.Z · build yymmddhhmm`（10 位，如 `v1.2.3 · build 2611051259`），文本来自 `internal/buildinfo.Footer()`；页眉不得出现版本号（见 `CONTRIBUTING.md` §8.2） |

### 3.1 基础组件 — `src/components/ui/`（已弃用，v1.x 遗留）

| 组件 | 用途 | 当前页面是否使用 |
|------|------|------------------|
| `button.tsx` | 变体 `default/primary/secondary/destructive/ghost/outline`；尺寸 `sm/md/lg/icon`；`forwardRef` 透传原生属性 | ✅ |
| `card.tsx` | `Card` / `CardHeader` / `CardTitle` / `CardDescription` / `CardContent` | ✅（Description 未使用） |
| `badge.tsx` | `primary/secondary/destructive/success/warning` 状态徽章 | ✅ |
| `tabs.tsx` | `Tabs`（受控 `value` + 非受控 `defaultValue`）/ `TabsList` / `TabsTrigger` / `TabsContent` | ✅ |
| `alert.tsx` | `Alert`（tone）+ `AlertTitle` + `AlertDescription`，`role="alert"` | ✅（仅 warning） |
| `table.tsx` | `Table` / `THead` / `TBody` / `TR` / `TH` / `TD` | ❌ 已提供，未在页面使用 |
| `code-block.tsx` | 等宽代码展示 + 剪贴板复制（客户端组件） | ❌ 已提供，未在页面使用 |

> 所有基础组件透传 `...props`，因此可传入 `id`、`aria-*` 等属性（`CONTRIBUTING.md` §4.2 / §4.4）。

### 3.2 业务页面 — `src/app/page.tsx`（已弃用，v1.x 遗留）

单页 dashboard，客户端组件（`"use client"` + `useState<Lang>`），自上而下：

| 区块 | 语义化 id | 内容 |
|------|-----------|------|
| 页面容器 | `repo-cleaner-app`（`main`）、`app-container` | 全屏暗色底，`max-w-5xl` 居中 |
| 顶部栏 | `app-header`、`app-title`、`lang-switch`、`lang-btn-zh`、`lang-btn-en` | 标题 + 中/英实时切换 |
| 概览卡片 | `overview`、`stat-total`、`stat-safe`、`stat-tweaks` | 扫描总数 / 安全数 / 优化项数 |
| 安全提示 | `security-note` | `warning` 级 Alert，提示删除不可逆 |
| 主内容 | `main-tabs`、`main-tabs-list`、`tab-scan`、`tab-tweaks` | Tabs 切换「扫描结果」/「优化建议」 |
| 扫描结果 | `tab-content-scan`、`scan-results`、`repo-card-<repo>`、`repo-info-<repo>`、`repo-status-<repo>`、`btn-clean-<repo>` | 仓库卡片 + 风险徽章 + 清理按钮（当前为展示按钮，无实际删除行为） |
| 优化建议 | `tab-content-tweaks`、`tweak-list`、`tweak-card-<id>`、`tweak-head-<id>` | 分类徽章 + 标题 + 说明 |
| 页脚 | `app-footer` | 文案 + `vX.Y.Z · build yymmddhhmm`（10 位，由 `next.config.ts` 注入，见 `CONTRIBUTING.md` §8.2） |

根布局 `layout.tsx`：`root-html` / `root-body` / `app-shell`，`metadata` 提供标题与描述。

---

## 4. 交互标准

### 4.1 已实现

- 单页 dashboard，`Tabs` 在「扫描结果」与「优化建议」间切换（非受控 `defaultValue="scan"`）。
- 语言开关实时切换中/英（`useState` 受控，切换全部 UI 文案）。
- 按钮 hover / active 过渡；焦点可见轮廓。
- `CodeBlock` 的复制按钮有「复制 / 已复制」状态反馈（`setTimeout` 1.5s 复位），剪贴板不可用则静默降级。

### 4.2 规划（当前未实现）

以下为设计意图，**已弃用的参考界面是纯展示 dashboard，不执行任何真实清理动作**（真实清理由 v2.0.0 的 Go + Fyne 版实现，含二次确认），因此尚未落地：

- 操作成功 / 失败的 `alert` 反馈（success / destructive）。
- 加载态：按钮 `disabled` + 旋转指示器。原设计设想使用第三方图标库的 Spinner，但仓库**不引入任何图标库**；若实现须使用 `src/components/icons.tsx` 内联 SVG 或已有的 `.animate-spin-slow`。
- 路径不存在的错误 `alert` + 重新输入（对齐批处理的行为，见 `OPENSPEC.md` §4.5）。
- 删除前的 warning 二次确认弹层。
- 空状态（无扫描结果时的居中图标 + 引导按钮）。
- 清理按钮的真实行为（当前 `btn-clean-*` 无 `onClick`）。

---

## 5. 响应式断点

| 断点 | 阈值 | 设计约定 | 当前实现 |
|------|------|----------|----------|
| 默认 | < 640px | 单列 | `grid-cols-1` |
| `sm` | ≥ 640px | 概览卡片 3 列 | `sm:grid-cols-3`（已实现） |
| `md` | ≥ 768px | 多列 | 未单独定义，沿用 `sm` |
| `lg` | ≥ 1024px | 多列 + 侧栏 | **无侧栏**（规划） |
| `xl` | ≥ 1280px | 内容居中 | 容器 `max-w-5xl`（1024px），**非** 1200px |

> 历史版本曾写「xl 最大宽 1200px 居中」「lg 4 列 + 侧栏」，与代码不符，已按实现修正；若后续调整布局，须同步本表与 `tokens.json`（若涉及令牌）。

---

## 6. 可访问性 (a11y)

**已实现**

- 语义标签：`main` / `header` / `section` / `footer`。
- `TabsTrigger` 有 `role="tab"` 与 `aria-selected`；`TabsContent` 有 `role="tabpanel"`；`Alert` 有 `role="alert"`。
- 装饰性 SVG 图标带 `aria-hidden="true"`。
- 焦点可见：全局 `:focus-visible` 轮廓（2px `primary`，`outline-offset: 2px`）。
- 语言切换按钮使用可读文本（「中文」/「English」）而非纯图标。
- 颜色对比：暗色底 `#0a0a0b` 配 `#ededed` 主文字、`#a1a1aa` 次级文字，满足 WCAG AA。

**已知限制 / 待办**

- `<html lang="zh">` 在 `layout.tsx` 中固定，语言切到英文时**不更新**，屏幕阅读器仍按中文朗读。
- `TabsList` 未加 `role="tablist"`，Tab 之间缺少 `aria-controls` / 方向键导航。
- `CodeBlock` 复制按钮的 `aria-label="复制代码"` 未随界面语言切换。

---

## 7. 规范与实现差距清单

| 编号 | 差距 | 位置 | 处理建议 |
|------|------|------|----------|
| G1 | 亮色主题仅有 `tokens.json` 规划值，CSS 未实现 | §2.1 | 保持规划标注，或实现 `prefers-color-scheme` 切换 |
| G2 | `--space-*` / `--radius` 已定义但组件使用具名数值类 | §2.3 / §2.4 | 逐步替换为令牌驱动的工具类 |
| G3 | `animate-fade-in-up` / `animate-spin-slow` 已定义未消费 | §2.6 | 需要时启用，否则保留为组件库能力 |
| G4 | 成功/失败反馈、加载态、错误态、空状态、删除确认均未实现 | §4.2 | 明确列为规划，禁止写成既有能力 |
| G5 | `table` / `code-block` 组件已提供但页面未使用 | §3.1 | 保留为组件库能力，或按需接入 |
| G6 | `lg` 侧栏、1200px 容器为历史设计，实际是 1024px 无侧栏 | §5 | 已按实现修正文档 |
| G7 | `<html lang>` 不随语言切换、Tabs 缺少完整 ARIA | §6 | 后续补齐 |
| G8 | ~~清理按钮无 `onClick`，界面为纯展示~~ | §3.2 | **v2.0.0 已解决**：Go + Fyne 版实现真实清理（含二次确认）；Next.js 参考界面保持纯展示并已弃用 |

> 本清单是「规范与实现一致」的核对表：修复任一项时，同步更新本文对应章节与本表，并在 `CHANGELOG.MD` 记录。

---

## 8. 变更同步约定

- 新增 / 修改令牌：先改 `tokens.json` 与本文 §2，再改 `internal/ui` 的 Fyne 主题配色（`src/app/globals.css` 为已弃用参考界面遗留），最后跑 `go build ./...`。
- 新增组件：更新 §3.1 表格并标注是否已接入页面；补语义化 `id`。
- 修改布局或断点：同步 §5 表格。
- 变更页脚版本 / build 编号格式：同步 `CONTRIBUTING.md` §8.2、`OPENSPEC.md` §4.6、本文 §3.0 / §3.2、`internal/buildinfo`、`next.config.ts` 与 `scripts/build.*`。
- 本文与 `docs/tokens.json` 的 `version` 随产品版本 bump 同步（见 `CONTRIBUTING.md` §8.2）。
