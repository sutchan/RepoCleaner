# RepoCleaner 开发规范与贡献指南 v1.2.5

> 本文是 RepoCleaner 仓库的**唯一开发规范来源**，规定「如何改代码、如何写文档、如何提交与评审」。
> 本文存放于 `docs/`；所有贡献者（含 AI 协作会话）在动手前必须阅读本文。
> 文档导航见 [docs/README.md](./README.md)；若本文与代码实现冲突，以代码为准并及时修正本文。

---

## 1. 文档体系与职责边界

仓库文档按「一个主题一份权威文档、统一归档 `docs/`」组织，禁止同一事实在多处重复维护。下表路径均为**仓库根相对路径**。

| 文档 | 职责（写什么） | 不写什么 |
|------|---------------|----------|
| `README.MD`（根） | 面向使用者：功能、快速开始、参数、FAQ、文档索引 | 不写开发规范、不写版本历史明细 |
| `docs/README.md` | 文档中心：文档地图、职责边界、按角色阅读路径、维护约定 | 不复述各规范正文 |
| `docs/OPENSPEC.md` | 产品功能规格：7 项功能的行为契约、交互、环境、兼容性、已知限制 | 不写代码风格、不写提交/分支规则、不重复变更日志 |
| `docs/CONTRIBUTING.md`（本文） | 开发规范：编码风格、命名、目录、提交、分支、审查、版本、术语、安全 | 不写功能介绍、不写设计令牌 |
| `docs/DESIGN-SYSTEM.md` | UI 设计规范：设计令牌、组件库、交互、响应式、a11y、差距清单 | 不写业务功能、不写开发流程 |
| `docs/tokens.json` | 机器可读设计令牌（与 `DESIGN-SYSTEM.md` §2 对齐） | 不作为产品版本号来源 |
| `prototype/README.md` | `prototype/` 目录说明：原型现状与目录用途 | 不复述设计规范正文、不存放代码 |
| `CHANGELOG.MD`（根） | **唯一版本历史来源**，按版本倒序记录变更 | 不重复功能规格 |
| `LICENSE`（根） | MIT 许可证正文 | — |
| `.github/PULL_REQUEST_TEMPLATE.md` | PR 审查清单模板 | — |

**引用规则**：需要版本历史 → `CHANGELOG.MD`；需要功能行为 → `docs/OPENSPEC.md`；需要协作流程 → 本文；需要视觉/组件规范 → `docs/DESIGN-SYSTEM.md`。禁止把 `CHANGELOG.MD` 的内容复制到 `docs/OPENSPEC.md` 或其它文档。

**文件名大小写**：根文档历史形成为大写扩展名 `README.MD` / `CHANGELOG.MD`；`docs/` 下全部为小写 `.md`；`docs/tokens.json` 为小写 `.json`。**所有链接必须按实际大小写与相对层级书写**（GitHub / Linux 区分大小写）：
- 从 `docs/` 内的文档引用根文件用 `../README.MD`、`../CHANGELOG.MD`、`../LICENSE`；
- 引用同目录文档用 `./OPENSPEC.md`、`./CONTRIBUTING.md`、`./DESIGN-SYSTEM.md`、`./tokens.json`；
- 禁止把链接目标写成小写扩展名的 `CHANGELOG.md` 或错误大小写的 `Openspec.md` 等不存在的路径。

**禁止**：在 `.github/` 下创建 `readme.md`（与根 README 重复且易误导）；Community Health 文件与 PR 模板不受此限。

---

## 2. 环境与可用命令

| 项 | 要求 |
|----|------|
| 操作系统 | Windows（运行 `RepoCleaner.bat`）；参考界面应用可跨平台开发 |
| Node.js | ≥ 18（建议 LTS）；npm 随 Node 安装 |
| 包管理器 | **npm**（仓库提交 `package-lock.json`，不使用 yarn / pnpm 锁文件） |
| 框架 | Next.js 16（App Router）+ React 19 + TypeScript 5（strict）+ Tailwind CSS v4 |

在**仓库根目录**执行：

```bash
npm install            # 安装依赖（改动 package.json 后同步更新 package-lock.json）
npm run dev            # 启动开发服务器
npm run type-check     # TypeScript 类型检查（tsc --noEmit）
npm run lint           # ESLint 检查
npm run check:docs     # 校验所有 Markdown 相对链接（含文件名大小写）
npm run build          # 生产构建（含类型检查 + 静态预渲染）
npm run start          # 启动生产构建产物（需先 build）
```

> 仓库**没有** `npm test` 脚本，不要臆造测试命令。当前质量门禁是 `type-check` + `lint` + `check:docs` + `build` 四项，外加 §9.3 的批处理手动冒烟清单。

---

## 3. 目录组织

### 3.1 根目录布局

```
RepoCleaner/
├── .github/                     # PR 模板等仓库协作配置（禁止创建 .github/readme.md）
│   └── PULL_REQUEST_TEMPLATE.md
├── docs/                        # 所有规范与规格文档归档
│   ├── README.md                # 文档中心
│   ├── OPENSPEC.md              # 产品功能规格
│   ├── CONTRIBUTING.md          # 本文：开发规范
│   ├── DESIGN-SYSTEM.md         # UI 设计规范
│   └── tokens.json              # 设计令牌
├── public/                      # 参考界面应用静态资源（SVG）
├── scripts/                     # 工程校验脚本
│   └── check-doc-links.mjs      # 文档相对链接校验（npm run check:docs）
├── src/                         # 参考界面应用代码（Next.js App Router）
│   ├── app/                     # 路由与全局样式：layout.tsx / page.tsx / globals.css / favicon.ico
│   ├── components/
│   │   ├── icons.tsx            # 内联 SVG 图标（零外部图标依赖）
│   │   └── ui/                  # 无状态基础组件：alert/badge/button/card/code-block/table/tabs
│   └── lib/
│       ├── i18n.ts              # 唯一 UI 文案字典（zh / en）
│       ├── utils.ts             # 纯工具函数（cn 等）
│       └── data/                # 单一数据源：features/project-types/steps/repos + index.ts barrel
├── prototype/                   # 原型目录，仅保留 README.md 说明，无代码、无规范正文
│   └── README.md
├── RepoCleaner.bat              # 产品本体：Windows 批处理脚本
├── config.ini                   # 运行配置：REPO_ROOT / LANG
├── lang.ini                     # 语言文案参考节 [en] / [zh]
├── README.MD / CHANGELOG.MD / LICENSE     # 入口、变更历史、许可证
├── RepoCleaner.code-workspace   # VS Code 工作区配置，禁止删除
├── package.json / package-lock.json / tsconfig.json / next.config.ts
├── postcss.config.mjs / eslint.config.mjs / next-env.d.ts
└── .gitignore
```

### 3.2 组织规则

1. **所有规范 / 规格文档放 `docs/`**；根只保留入口 `README.MD`、`CHANGELOG.MD`、`LICENSE`。**应用代码一律放根 `src/`**；`prototype/` 只放原型说明，禁止再放回可运行代码或规范正文。
2. **单一数据源**：业务展示数据只放 `src/lib/data/`，统一经 `@/lib/data` 导出；UI 文案只放 `src/lib/i18n.ts`。禁止在组件内硬编码展示数据或文案。
3. **路径别名**：统一用 `@/` 指向 `./src/*`（见 `tsconfig.json` `paths`），禁止跨层使用长串 `../../../`。
4. **新增模块按职责落位**：路由 → `src/app/`；可复用 UI → `src/components/ui/`；图标 → `src/components/icons.tsx`；纯函数 → `src/lib/`；数据 → `src/lib/data/`；规范 → `docs/`；工程校验脚本 → `scripts/`。
5. **文件行数**：所有源代码文件（`.ts/.tsx/.js/.jsx/.vue/.py/.dart/.go/.java` 等）单文件 **≤ 200 行**。超出时按职责拆分为更小模块（组件拆子组件、逻辑抽纯函数/常量/类型），且**不得改变对外导出契约**。
   - **文档文件（`.md/.markdown/.mdx/.rst`）不拆分**，即使超过 200 行也保持单文件完整连贯。
   - `.bat` 为批处理脚本，不纳入源码 200 行规则，但仍须按标签分节、保持可读。

---

## 4. 编码风格

### 4.1 通用

- **编码**：所有文本文件使用 **UTF-8**；`RepoCleaner.bat` 通过 `chcp 65001` 将控制台切到 UTF-8（不写「GBK/UTF-8」这类模棱两可的表述）。
- **缩进**：TypeScript / JS / CSS / JSON 用 2 空格；`RepoCleaner.bat` 用 4 空格；禁止 Tab。
- **引号**：TS/JS 统一双引号；JSX 属性用双引号；CSS 用双引号（字体名等）。
- **分号**：TS/JS 语句必须写分号。
- **注释**：关键逻辑、非直觉判断（尤其批处理的延迟展开与安全边界）必须写**中文注释**；禁止用注释复述显而易见的代码。
- **文件头注释**：源代码文件首行标注 `// <相对路径> vX.Y.Z`；CSS 用 `/* ... */`；`.bat` 用 `::`；`.ini` 用 `;`。示例：
  ```ts
  // src/components/ui/button.tsx v1.2.5
  ```
  头注释中的版本号是**该文件自身的版本**，随「本次实际修改该文件」更新；**未改动的文件不得批量刷新头注释**。

### 4.2 TypeScript / React

- 启用 `strict`，**禁止 `any`**；需要宽松类型时用 `unknown` + 类型收窄。
- 组件用**函数组件 + 命名导出**；仅页面/布局用 `export default`。
- 需要 hooks、事件、`navigator` 等浏览器 API 的组件，文件首行必须是 `"use client"`（放在 import 之前）。App Router 下服务端组件使用 `useState` 会导致构建失败。
- 组件 props 继承原生属性（如 `React.ButtonHTMLAttributes<HTMLButtonElement>`）并透传 `...props`，保证 `id`、`aria-*` 可传入。
- 样式用 `cn()`（`src/lib/utils.ts`）合并 className，禁止字符串拼接。
- 状态尽量局部化；无状态组件不引入 hooks。
- `Promise` 异常必须捕获处理（如 `CodeBlock` 的 clipboard 写入用 `try/catch`）。

### 4.3 样式与设计令牌

- 颜色/字体/圆角/动效**一律走 `src/app/globals.css` 的 `@theme` 令牌生成的工具类**：`bg-background`、`text-foreground`、`text-muted-foreground`、`bg-card`、`border-border`、`text-primary` 等。
- **禁止散落硬编码颜色**（`bg-[#0d1117]`、`text-gray-400` 等任意值）；新令牌先在 `docs/tokens.json` 与 `docs/DESIGN-SYSTEM.md` 登记，再写入 `globals.css`。
- 不引入 `next/font/google` 等**构建期联网**资源：离线环境会导致 `next build` 失败。使用 `font-sans` 系统字体栈。
- 图标使用 `src/components/icons.tsx` 内联 SVG，**不引入第三方图标库**。
- 当前仅实现**暗色主题**（`color-scheme: dark`）；浅色值属规划，不得在代码中假定已实现。

### 4.4 语义化 id

为页面级容器、主要区块、关键交互元素添加语义化 `id`，便于调试定位、ARIA 关联与端到端测试选择器。

- 命名用 **kebab-case**，表达用途：`app-header`、`overview`、`main-tabs`、`security-note`、`app-footer`。
- 列表项用稳定前缀 + 业务键：`repo-card-<repo>`、`tweak-card-<id>`、`btn-clean-<repo>`。
- 禁止 `div1`、`box`、`a` 等无意义命名；主要容器 `id` 全仓库唯一。
- 新增页面/组件时同步补齐；重构超 200 行文件时一并检查。

### 4.5 Windows 批处理（`RepoCleaner.bat`）

- 保持 `setlocal enabledelayedexpansion`；块内（`( )` 中）读取块内赋值的变量必须用 `!VAR!` 而非 `%VAR%`。
- 所有路径变量**必须加双引号**，支持含空格路径。
- 删除目标只能来自 `for /d` 枚举出的目录变量，禁止把用户输入直接拼接进 `rd` / `del`。
- 功能 3 生成的 `.next` 文件带 `+h +s +r` 属性；功能 4 必须先 `attrib -h -s -r` 再 `del /f /q`（`del` 的多 `/a:` 只有末项生效，且默认不删隐藏/系统文件）。
- 危险操作（删除 `node_modules`）前必须有 `Y` 二次确认。
- 写 `config.ini` 用「读旧文件 → 过滤目标键 → 补写 → 临时文件 → `move /y`」方式，**禁止整文件覆盖 `lang.ini`**（会清空 `[en]/[zh]` 文案节）。

---

## 5. 命名约定

| 类别 | 规范 | 示例 |
|------|------|------|
| React 组件文件 | `PascalCase.tsx`（本仓库 `src/components/ui/` 沿用既有小写文件名，新增文件优先 PascalCase） | `CodeBlock.tsx` |
| 工具 / hooks / 普通模块 | `camelCase.ts` | `useLang.ts`、`utils.ts` |
| 数据模块 | 语义复数或主题名，`kebab-case` 可用 | `features.ts`、`project-types.ts` |
| 类型 / 接口 | `PascalCase`，不加 `I` 前缀 | `ScanResult`、`ButtonProps` |
| 常量 | `SCREAMING_SNAKE_CASE` | `APP_VERSION`、`DEFAULT_REPO_ROOT` |
| CSS 类名 / DOM id | `kebab-case` | `security-note`、`app-footer` |
| 文档文件 | 根目录既有 `README.MD` / `CHANGELOG.MD` 保持历史大写；`docs/` 下统一小写 `.md`，新增文档一律小写 | `docs/OPENSPEC.md`、`docs/README.md` |
| 分支 | 见 §7 | `fix/next-restore`、`docs/spec-sync` |
| 图标组件 | `PascalCase`，语义化 | `ScanLine`、`ShieldCheck` |

**术语统一**（全仓库文档与代码注释一致，禁止混用）：

| 统一写法 | 禁止写法 |
|----------|----------|
| `Next.js` | NextJS、Nextjs、next.js（句首除外） |
| `Nuxt` | NuxtJS |
| `node_modules`、`.next`、`dist`、`build` | 中文意译变体；代码块/行内代码用反引号 |
| 仓库根目录（`REPO_ROOT`，含多个项目的父目录） | 项目根目录（二者含义不同，禁止混用） |
| 项目根目录（单个项目，`REPO_ROOT` 的直接子目录） | — |
| 参考界面应用（根 `src/` 的 Next.js 应用） | prototype 应用、前端工程 |
| 规范文档目录（`docs/`） | 原型目录（后者专指 `prototype/`） |
| 原型目录（`prototype/`，仅说明文件） | 设计规范目录（规范在 `docs/`） |
| 批处理脚本（`RepoCleaner.bat`） | bat 脚本、主程序 |
| 功能 1 ～ 功能 7（名称见 `OPENSPEC.md` §3.0） | 功能一、F1（仅代码变量名可用 F1_） |

---

## 6. 提交信息格式

采用 Conventional Commits 精简格式：

```
<type>: <description>

[可选正文]

[可选页脚]
```

### 6.1 type（必填，限以下取值）

| type | 用途 |
|------|------|
| `feat` | 新增功能 |
| `fix` | 修复缺陷 |
| `docs` | 文档更新 |
| `style` | 代码风格调整（不改行为） |
| `refactor` | 重构（不改行为） |
| `test` | 测试相关 |
| `chore` | 构建 / 依赖 / 配置变更 |
| `perf` | 性能优化 |
| `ci` | CI/CD 配置变更 |
| `revert` | 回滚提交 |

### 6.2 书写要求

- `description`：**≤ 50 字符，小写字母或中文开头、以动词开头、不加句号**，如 `fix: 修复功能4恢复 .next 缓存失效`。
- 正文（可选）：每行 ≤ 72 字符，解释「为什么」与影响范围，用 `-` 列表。
- 页脚（可选）：引用 issue / PR、标注 breaking changes；**提交信息正文或页脚须包含本次新版本号**（见 §8）。
- 一次提交只做一件事；禁止把无关改动混入同一提交。

示例：

```
docs: 建立 docs 文档体系并同步规范至 v1.2.5

- 新建 docs/ 归档全部规格与规范文档，补充文档中心索引
- 对齐 OPENSPEC 与实现，修正过时描述并统一术语
- 版本号改由 package.json 单一来源在构建时注入

v1.2.5
```

### 6.3 提交前检查

```bash
git status                  # 确认无 node_modules / .next / *.log / *.tsbuildinfo 被跟踪
npm run type-check
npm run lint
npm run build
```

---

## 7. 分支管理

### 7.1 分支模型

| 分支 | 类型 | 用途 |
|------|------|------|
| `main` | 常驻 | 生产/发布分支，始终保持可构建；**当前仓库实际仅启用此分支** |
| `dev` | 常驻（按需启用） | 开发集成分支，稳定后合并到 `main` |
| `feature/<短描述>` | 短期 | 新功能，从 `dev` 切出 |
| `fix/<短描述>` | 短期 | 非紧急缺陷修复 |
| `docs/<短描述>` | 短期 | 纯文档变更 |
| `chore/<短描述>` | 短期 | 依赖 / 配置 / 构建变更 |
| `hotfix/<短描述>` | 短期 | 生产紧急修复，从 `main` 切出，修完同时回合 `dev` |
| `release/<版本>` | 短期 | 发布准备（版本号、CHANGELOG 定稿） |

短描述用 `kebab-case`，如 `feature/dry-run-mode`、`fix/next-cache-restore`。

### 7.2 规则

1. **禁止直接向 `main` 强推**；禁止 `--force` 覆盖 `main` / `dev` 历史。
2. 短期分支合并后删除；保留提交历史，**不使用 squash 合并**，便于回滚。
3. 分支名与提交 type 保持一致语义。
4. 远程为 `https://github.com/sutchan/RepoCleaner.git`，推送使用 `git push origin <分支>`。
5. 单人小改动（≤ 3 个文件、纯文档/拼写）可直接提交到 `main`；涉及 ≥ 5 个文件、新增依赖或架构调整，须走 PR 流程（§9）。

---

## 8. 版本号与变更记录

### 8.1 版本语义（SemVer）

| 级别 | 适用 | 升级方式 |
|------|------|----------|
| patch | 修复、文档、重构、样式、配置、依赖等任意修改 | 第三位 +1（`1.2.4 → 1.2.5`） |
| minor | 新功能、向后兼容变更 | 第二位 +1，第三位归 0 |
| major | 破坏性变更 | 第一位 +1，其余归 0 |

**任何文件修改都必须至少 bump 一次 patch**；minor/major 变更不再叠加 patch。

### 8.2 版本号单一来源

**唯一权威源：`package.json` 的 `version` 字段。** 参考界面页面的版本号与 build 编号由 `next.config.ts` 在构建时读取 `package.json` 并注入 `NEXT_PUBLIC_APP_VERSION` / `NEXT_PUBLIC_BUILD_ID`，`src/app/page.tsx` 页脚读取展示 —— 禁止在页面/组件中硬编码版本号。

每次 bump 需同步以下**全局展示位**（各一处）：

- `package.json` → `version`
- `README.MD`（根） → 版本徽章 `Version-X.Y.Z` 与页脚版本号
- `CHANGELOG.MD`（根） → 新增对应版本小节（置于最上方）
- `docs/tokens.json` → `version`
- `docs/DESIGN-SYSTEM.md`、`docs/README.md`、`docs/OPENSPEC.md`、`docs/CONTRIBUTING.md`、`prototype/README.md` → 标题版本
- 被改动**代码文件**的头注释版本号

**禁止**：全仓库批量刷写所有源码文件的头注释；未改动文件头注释保持不变（头注释是文件自版本，不等于产品版本，看到旧版本与新版本并存属正常，不是版本漂移）。

**build 编号**：格式 `yyyymmddhhmm`（本地构建时间到分钟），由 `next.config.ts` 自动生成，`dev` / `build` 均生效，无需手工维护。页脚注入缺失时回退为仅显示版本号。

### 8.3 CHANGELOG 写法

- 按版本倒序，标题 `## [X.Y.Z] - YYYY-MM-DD`。
- 分组用 `Added` / `Changed` / `Fixed` / `Docs` / `Security`，条目以 `- **要点** - 说明` 书写。
- 历史版本号是**历史记录，不得改写**；批量替换版本文本时必须排除 CHANGELOG 中的旧版本号与示例占位符。

---

## 9. 代码审查流程

### 9.1 流程

1. 从目标分支切短期分支（§7.1）。
2. 完成改动，按 §6.3 跑完全部门禁。
3. `git add` → `git commit`（格式见 §6）→ `git push origin <分支>`。
4. 在 GitHub 创建 PR 到 `main`（或 `dev`），**PR 标题格式同提交信息**，描述按 `.github/PULL_REQUEST_TEMPLATE.md` 填写。
5. 审查人按 §9.2 清单逐项核对；有异议在对应代码行评论，修改后重新推送。
6. 全部通过后以 **merge commit** 合并（不 squash），删除短期分支。

### 9.2 审查清单

- [ ] 功能行为与 `docs/OPENSPEC.md` 对应条目一致；变更功能时已同步规格文档
- [ ] 新增/修改的 UI 对齐 `docs/DESIGN-SYSTEM.md` 与 `docs/tokens.json`，无硬编码颜色
- [ ] 主要容器与关键元素已加语义化 `id`（§4.4）
- [ ] 无 `console.log` / `debugger`；无 `any`；无未使用的 import
- [ ] `npm run type-check`、`npm run lint`、`npm run build` 全部通过
- [ ] 单文件 ≤ 200 行（代码文件）；超出已按职责拆分且不破坏导出契约
- [ ] 批处理改动：路径带引号、延迟展开正确、危险操作有二次确认
- [ ] 提交信息格式合规，且含新版本号
- [ ] 版本号已 bump，`package.json` / README 徽章 / CHANGELOG / 展示位一致
- [ ] `CHANGELOG.MD` 已新增对应小节；`README.MD` / `docs/OPENSPEC.md` 使用说明随功能变化更新
- [ ] 文档链接按实际大小写与相对层级书写，新增文档归入 `docs/`
- [ ] 无密钥 / Token / 真实用户数据入库；`node_modules`、`.next`、`*.log`、`*.tsbuildinfo` 未被跟踪

### 9.3 功能验证（批处理手动冒烟）

`RepoCleaner.bat` 无自动化测试，改动后须在**隔离的临时目录**（放 1～2 个含 `package.json` 的假项目）手动验证：

| 检查项 | 预期 |
|--------|------|
| `RepoCleaner.bat -h` | 输出帮助信息后退出 |
| `RepoCleaner.bat -lang en` | `config.ini` 写入 `LANG=en` 且不产生重复行，`lang.ini` 文案节完好 |
| 菜单 `[L]` | 中/英切换生效 |
| 功能 1（输入 Y） | 生成 `.npmrc`、清理本地 `node_modules`；输入非 Y 则返回菜单且不删除 |
| 功能 3 → 功能 4 | 生成带 `+h+s+r` 的 `.next` 文件；功能 4 能成功删除该文件（无残留、无报错） |
| 功能 6 / 7 | 仅清理白名单目录；含空格路径可正确处理 |
| 路径不存在 | 显示错误并允许重新输入 |

> 严禁在真实仓库根目录做删除类验证。

---

## 10. 文档维护约定

1. 改功能 → 同步 `docs/OPENSPEC.md` + 根 `README.MD`；改 UI → 同步 `docs/DESIGN-SYSTEM.md` + `docs/tokens.json`；改流程 → 同步本文；改文档结构 → 同步 `docs/README.md`。
2. 文档中的**路径、命令、文件名、版本号必须与仓库实况一致**；写文档前先 `git log -1` 与 `node -p "require('./package.json').version"` 实查，禁止凭记忆填写。
3. 禁止在文档中保留「规划中但未实现」的内容而不加标注；要么删除，要么明确写「规划，当前未实现」，或记入 `docs/DESIGN-SYSTEM.md` §7 差距清单。
4. 数据型描述（数量、行数、命令数）须实算或引用唯一来源，禁止估算。
5. 文档文件不拆分；超长文档保持单文件完整。
6. 禁止创建 `.github/readme.md`；`RepoCleaner.code-workspace` 是工作区配置，禁止删除。
7. 改动文档后校验所有 Markdown 相对链接可达（含大小写），链接检查脚本思路见 `docs/README.md` §5。

---

## 11. 安全红线

- 代码、配置、日志中不得包含密钥 / Token / 凭据。
- 用户输入必须校验：批处理路径只用 `if exist` 判定，删除目标只取 `for /d` 枚举值并加引号。
- 前端不拼接未净化的 HTML（本仓库无 `dangerouslySetInnerHTML`）。
- 删除类操作必须二次确认，且不得在真实仓库根目录做验证。
- 依赖定期审计（`npm outdated`）；升级 `next` / `react` 等核心依赖须跑完整门禁。
