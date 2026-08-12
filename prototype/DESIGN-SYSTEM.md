# RepoCleaner 设计系统 (Design System) v1.2.1

> 本文件是 prototype/ 的**唯一设计规范来源**。所有页面、组件、文档均须对齐此处。
> 配套前端工程：Next.js 16 + React 19 + Tailwind 4 + shadcn/ui 风格组件。

---

## 1. 设计原则 (Design Principles)

1. **极简优先** — 留白即设计，单屏只解决一件事。
2. **数据真实** — 所有展示数据来自 `src/lib/data.ts`（与 OPENSPEC/README 一致）。
3. **一致令牌** — 颜色/间距/字号全部走 CSS 变量，禁止散落硬编码。
4. **响应式无妥协** — 移动端单列、桌面多列，断点 `sm/md/lg/xl`。
5. **克制动效** — 仅用于引导注意力与状态反馈，时长 ≤ 300ms。

---

## 2. 设计系统 (Design Tokens)

### 2.1 色彩 (Color)

基于中性灰 + 单一品牌强调色（青绿），暗色优先（开发者工具调性）。

| 令牌 | 浅色 | 暗色 | 用途 |
|------|------|------|------|
| `--background` | `#ffffff` | `#0a0a0b` | 页面底色 |
| `--foreground` | `#0a0a0b` | `#ededed` | 主文字 |
| `--card` | `#ffffff` | `#161618` | 卡片底 |
| `--muted` | `#f4f4f5` | `#1c1c1f` | 次级底 |
| `--muted-foreground` | `#71717a` | `#a1a1aa` | 次级文字 |
| `--border` | `#e4e4e7` | `#27272a` | 描边 |
| `--primary` | `#10b981` | `#10b981` | 品牌主色（emerald-500） |
| `--primary-foreground` | `#ffffff` | `#052e22` | 主色上文字 |
| `--destructive` | `#ef4444` | `#ef4444` | 危险/删除 |
| `--success` | `#22c55e` | `#22c55e` | 成功 |
| `--warning` | `#f59e0b` | `#f59e0b` | 警告 |

**语义规则**：删除类操作（功能3/6/7 清理）用 destructive；安装/启用类用 primary；成功反馈用 success。

### 2.2 字体 (Typography)

| 层级 | 字号 | 字重 | 行高 | 用途 |
|------|------|------|------|------|
| Display | 48px | 700 | 1.1 | 首页大标题 |
| H1 | 32px | 700 | 1.2 | 页面标题 |
| H2 | 24px | 600 | 1.3 | 区块标题 |
| H3 | 18px | 600 | 1.4 | 卡片标题 |
| Body | 15px | 400 | 1.6 | 正文 |
| Small | 13px | 400 | 1.5 | 辅助文字 |
| Code | 13px | 500 | 1.5 | 等宽（ui-monospace） |

字体栈：`-apple-system, "Segoe UI", "PingFang SC", "Microsoft YaHei", sans-serif`；
代码：`ui-monospace, "JetBrains Mono", "SF Mono", Consolas, monospace`。

### 2.3 间距 (Spacing)

4px 基准栅格：`--space-1=4px ... --space-8=32px`，组件内边距统一 `p-4/p-6`，卡片间距 `gap-4`。

### 2.4 圆角与阴影 (Radius & Shadow)

- 圆角：`--radius=12px`（卡片），按钮 `8px`，徽章 `9999px`。
- 阴影：暗色下以 `border` 代替阴影，浅色用 `shadow-sm`/`shadow-md`。

### 2.5 图标 (Icon)

- 库：`lucide-react`（线性、1.5px 描边、24px 默认）。
- 映射：全局 node_modules→`Boxes`，重置→`RotateCcw`，.next→`Zap`，恢复→`PlayCircle`，依赖→`PackagePlus`，清理→`Trash2`，Vite→`Bolt`，语言→`Languages`，退出→`LogOut`。

### 2.6 动效 (Motion)

- 时长：`--dur-fast=150ms`，`--dur=250ms`，缓动 `cubic-bezier(0.4,0,0.2,1)`。
- 用途：hover 抬升（`translateY(-2px)`）、卡片入场 `fade-in-up`、状态切换淡入。

---

## 3. 组件库 (Component Library)

### 3.1 基础组件 (Base) — `src/components/ui/`

| 组件 | 等价 shadcn | 说明 |
|------|------------|------|
| `button` | Button | 变体 primary/secondary/destructive/ghost/outline |
| `card` | Card | Card/CardHeader/CardTitle/CardContent |
| `badge` | Badge | 状态/标签 |
| `tabs` | Tabs | 语言/视图切换 |
| `table` | Table | 功能/清理项对照 |
| `alert` | Alert | 警告/成功反馈 |
| `code-block` | — | 配置/命令展示（等宽+复制） |

### 3.2 复合组件 (Composite) — `src/components/`

- `SiteHeader` / `SiteFooter` — 站点外壳
- `FeatureCard` — 单个功能卡（图标+标题+描述+操作）
- `StatGrid` — 数据指标网格
- `TerminalMock` — 模拟 CLI 菜单（还原 OPENSPEC 5.2）
- `LangToggle` — 中英文切换（受控）

### 3.3 业务组件 (Business)

- `ConsolePanel` — 控制台，渲染 7 个功能，可点选并执行（模拟）
- `FeatureTable` — 功能↔清理项对照表
- `ConfigViewer` — config.ini / .npmrc 配置展示

---

## 4. 交互标准 (Interaction)

### 4.1 模式 (Patterns)

- 首页 → 功能总览 → 控制台（可执行模拟）/ 文档。
- 所有导航走顶部 `SiteHeader`，移动端折叠为抽屉。

### 4.2 反馈 (Feedback)

- 操作成功：`alert` success + 图标；失败：`alert` destructive。
- 加载：按钮 `disabled` + `Spinner`（lucide `Loader2` 旋转）。

### 4.3 错误 (Error)

- 路径不存在：红色 `alert` + 重新输入提示（对齐 OPENSPEC 5.3）。
- 删除前：`alert` warning 二次确认（destructive 按钮）。

### 4.4 空状态 (Empty State)

- 无扫描结果：居中图标 + 文案 + 主操作按钮引导。

---

## 5. 响应式断点

| 断点 | 阈值 | 布局 |
|------|------|------|
| sm | ≥640px | 2 列 |
| md | ≥768px | 3 列 |
| lg | ≥1024px | 4 列 + 侧栏 |
| xl | ≥1280px | 最大宽 1200px 居中 |

---

## 6. 可访问性 (a11y)

- 语义标签 `header/nav/main/section/footer`；
- 交互元素 `aria-label`；颜色对比 ≥ WCAG AA；
- 键盘可达，focus 态可见（`ring-2 ring-primary`）。
