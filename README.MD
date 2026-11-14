# 🧰 Repo Manager

<div align="center">

[![Platform](https://img.shields.io/badge/Platform-Win%20%7C%20macOS%20%7C%20Linux-blue?style=flat)](https://github.com/sutchan/RepoManager)
[![Go](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev)
[![Version](https://img.shields.io/badge/Version-2.3.2-green?style=flat)](https://github.com/sutchan/RepoManager)
[![License](https://img.shields.io/badge/License-MIT-orange?style=flat)](https://github.com/sutchan/RepoManager)
[![Node](https://img.shields.io/badge/Node.js-18%2B-339933?style=flat&logo=node.js&logoColor=white)](https://nodejs.org)
[![i18n](https://img.shields.io/badge/i18n-中文%2FEnglish-9C27B0?style=flat)](https://github.com/sutchan/RepoManager)

</div>

> 一键让所有前端项目保持干净：无冗余 `node_modules`、无 `.next` 缓存、无打包垃圾。

适配项目类型：**Next.js / React / Vue / Vite / Angular / Svelte / 常规 Node 项目**

---

## 📑 目录

- [功能亮点](#-功能亮点)
- [快速开始](#-快速开始)
- [功能详解](#-功能详解)
- [智能配置](#-智能配置)
- [常见问题 FAQ](#-常见问题-faq)
- [注意事项](#-注意事项)
- [排障与维护](#-排障与维护)
- [文档体系与开发规范](#-文档体系与开发规范-docs)
- [目录结构](#-目录结构)
- [许可证](#-许可证)
- [更新日志](#-更新日志)

---

## ✨ 功能亮点

本助手用于管理项目目录，自动解决以下痛点：

| 问题 | 解决方案 |
|------|----------|
| 每个项目重复生成巨大的 `node_modules` 占用空间 | 全局 node_modules，项目永不生成本地依赖 |
| Next.js 项目产生大量 `.next` 缓存文件 | 禁用/启用 `.next` 缓存一键切换 |
| 项目打包产生 `dist`/`build` 垃圾文件 | 一键清理所有缓存 |
| 依赖混乱、无法统一管理 | 统一安装到全局，告别重复安装 |
| Vite / Turborepo / Angular / Svelte 构建缓存膨胀 | 专清 `.vite`/`.turbo`/`.angular`/`.svelte-kit`/`coverage` 等目录 |

**特性**

- 🖥️ Go 编译的**单文件可执行程序**（Windows / macOS / Linux），零第三方运行时依赖
- 🎨 Fyne 桌面 GUI，一键操作；给出功能号 1–7 时进入**非交互命令行模式**
- 🧪 预演模式（dry-run）：先看将要删除什么，再决定是否执行
- 🌲 递归扫描（monorepo）：`-recursive` / GUI 复选框 / `RECURSIVE=true` 覆盖深层项目（深度 3 层），默认关闭
- 🌏 中文 / English 双语界面与执行日志（GUI 按钮 / `-lang` / `config.ini` 切换）
- 🗂️ 自定义清理目录：`config.ini` 的 `EXTRA_CLEAN_DIRS` 追加目录名，仅限项目目录内部
- 📝 执行日志落盘：CLI 与 GUI 同步追加写入可执行文件同目录的 `repo-manager.log`
- ⚙️ 目录来源：GUI 输入框与目录选择、`config.ini`、CLI 位置路径（须与功能号同用）
- 💾 仓库根目录可保存到 `config.ini`，一次设置长期生效
- 🛡️ 危险操作（功能 1 / 6 / 7）前二次确认，避免误删
- 📊 全局依赖安装逐组检查退出码，失败项汇总提示

---

## 🚀 快速开始

### 前置要求

- Windows / macOS / Linux 系统
- 已安装 Node.js 和 npm（建议 Node 18+；功能 1 / 5 需要）
- 源码构建需 Go ≥ 1.25；Fyne GUI 编译需 C 工具链与系统 OpenGL

### 使用步骤

1. **获取 `repo-manager` 可执行文件**：从 GitHub Release 下载 `repo-manager-<os>-<arch>[.exe]`（Windows 带 `.exe`），用随附的 `SHA256SUMS.txt` 校验；或源码构建（`./scripts/build.ps1` / `./scripts/build.sh` 交叉编译到 `bin/`；本机快速构建 `go build -o .tmp/repo-manager ./cmd/repo-manager`，临时产物统一进 `.tmp/`，Fyne GUI 需 C 工具链）
2. **运行**：不带参数启动 GUI（Windows 双击，macOS/Linux 命令行执行）；带功能号 `1`–`7` 则直接进入命令行模式
3. **选择仓库根目录**（GUI 输入框 / 目录选择按钮，可保存到 `config.ini`；CLI 用位置路径）
4. **按功能按钮（GUI）或功能号（CLI）选择功能**
5. **一次配置，所有项目永久干净**

> 位置路径只在给出功能号 1–7 时生效；单独给路径会启动 GUI 并忽略该路径。
> 应用页脚展示 `vX.Y.Z · build yymmddhhmm`（10 位构建编号，如 `v2.2.1 · build 2611071100`），便于核对当前构建版本与时间；格式与注入规则见[开发规范 §8.2](./docs/CONTRIBUTING.md)。

### 🌟 推荐使用流程

| 场景 | 操作 |
|------|------|
| 首次配置 | 运行 `[1]` 启用全局 node_modules → 运行 `[5]` 安装所有常用依赖 → 运行 `[3]` 禁用 `.next` 缓存 |
| 空间不足 | 运行 `[6]` 清理所有缓存 |
| Vite / Turborepo / Angular / Svelte 项目 | 运行 `[7]` 清理 Vite/构建缓存 |

---

## 📖 功能详解

### 功能 1：启用全局 node_modules

- 自动获取全局 npm 模块路径（`npm root -g`）
- 配置系统环境变量 `NODE_PATH`
- 为所有包含 `package.json` 的项目生成 `.npmrc`，禁用本地依赖生成
- 清理现有的本地 `node_modules` 目录（执行前二次确认：GUI 弹窗 / CLI 输入 `Y`）

> 首次使用后，项目将共享全局依赖，不再生成 `node_modules` 文件夹

### 功能 2：重置 Node 配置

- 清除 `NODE_PATH` 环境变量
- 删除所有项目的 `.npmrc` 配置文件
- 恢复 npm 默认行为

### 功能 3：清理并禁用 .next 缓存

- 扫描所有包含 `next.config.js` / `next.config.ts` 的项目
- 删除现有的 `.next` 缓存目录
- 创建不可见的 `.next` 文件（隐藏 + 系统 + 只读属性）阻止缓存生成

> 适合磁盘空间紧张时使用，会影响 Next.js 热更新速度

### 功能 4：恢复 Next.js 缓存

- 删除禁止缓存的 `.next` 文件（先清除隐藏/系统/只读属性再删除）
- 恢复 Next.js 正常编译和热更新功能

### 功能 5：安装常用全局依赖

一次性安装以下常用开发库：

```
react react-dom next
vue @vue/cli
axios express
pnpm yarn
cross-env dotenv nodemon pm2 rimraf mkdirp
```

### 功能 6：清理所有项目缓存

清理以下内容：

| 清理项 | 说明 |
|--------|------|
| `node_modules` | 项目依赖目录 |
| `.next` | Next.js 缓存 |
| `.nuxt` | Nuxt 缓存 |
| `dist` | 打包输出目录 |
| `build` | 构建输出目录 |
| `package-lock.json` | npm 锁文件 |
| `yarn.lock` | yarn 锁文件 |
| `pnpm-lock.yaml` | pnpm 锁文件 |

### 功能 7：清理 Vite/构建缓存

清理以下内容：

| 清理项 | 说明 |
|--------|------|
| `.vite` | Vite 缓存目录 |
| `.turbo` | Turborepo 缓存 |
| `.nuxt` | Nuxt 缓存 |
| `.angular` | Angular 缓存（v2.2.0 新增） |
| `.svelte-kit` | SvelteKit 生成目录（v2.2.0 新增） |
| `dist` | 打包输出目录 |
| `build` | 构建输出目录 |
| `coverage` | 测试覆盖率报告 |
| `.nyc_output` | NYC 输出目录 |

---

## ⚙️ 智能配置

### 目录选择方式

运行程序时可通过以下方式指定仓库根目录：

| 方式 | 说明 |
|------|------|
| CLI 位置路径 | 须与功能号同用，如 `repo-manager 6 "D:\Projects"` |
| 配置文件 | **可执行文件同目录**的 `config.ini`（`REPO_ROOT=`；本机生成、不入库，示例见 [`docs/config.example.ini`](./docs/config.example.ini)） |
| GUI | 窗口输入框 / 目录选择按钮，可用「保存」写回 `config.ini` |

> 不带功能号时启动 GUI，位置路径被忽略；GUI 使用 `config.ini` 的 `REPO_ROOT`，缺省时取平台默认目录。

### 语言切换

支持中文/英文界面与执行日志切换：

| 方式 | 说明 |
|------|------|
| GUI 下拉框 | 顶部「中文 / English」实时切换并写入 `config.ini` |
| 命令行参数 | `repo-manager -lang en`（写入 `LANG=en`） |
| 配置文件 | `config.ini` 中 `LANG=zh` 或 `LANG=en` |

> 语言仅 `zh` / `en` 两种：除 `en` 外的任何值（含非法值、空值）均回落中文 `zh`。

### 配置文件示例

完整模板见 [`docs/config.example.ini`](./docs/config.example.ini)（共 4 个键，键缺失时按内部默认值处理）：

```ini
REPO_ROOT=E:\Github
LANG=zh
RECURSIVE=false
EXTRA_CLEAN_DIRS=
```

### 命令行参数

```bash
repo-manager [选项] [功能号] [目录路径]
```

| 参数 | 说明 |
|------|------|
| `-lang en\|zh` | 设置界面语言并写入 `config.ini`（非 `en` 值回落 `zh`） |
| `-dry-run` | 预演模式：仅显示将要执行的操作，不修改任何文件与环境变量 |
| `-recursive` | 递归扫描 monorepo 深层项目（深度 3 层；也可经 `config.ini` 的 `RECURSIVE=true` 持久化） |
| `EXTRA_CLEAN_DIRS`（config.ini） | 自定义清理目录名（逗号分隔），并入功能 6 / 7 白名单，仅限项目目录内部 |
| `-h` / `-help` / `--help` | 显示帮助信息（固定英文，含版本与 build 信息） |
| `1`-`7` | 直接运行对应功能，进入非交互命令行模式；无功能号 `0` |
| 目录路径 | 指定仓库根目录，**仅在给出功能号时生效** |

> 无 `-root` 标志、无 `-version` 标志、无控制台交互菜单；功能号超出 `1`–`7` 的数字会被当作路径处理。
> CLI 与 GUI 的执行日志都会追加写入可执行文件同目录的 `repo-manager.log`；`-h` 帮助固定英文（不随 `-lang` 切换），其中 `-recursive` 的递归深度与实际一致，均为 3 层。

**示例**

```bash
repo-manager                       # GUI 模式
repo-manager 1                     # CLI：直接启用全局 node_modules
repo-manager 5 E:\Projects         # CLI：为指定根目录安装全局依赖
repo-manager 6 "D:\My Projects"    # CLI：指定含空格的仓库根目录
repo-manager -lang en              # 写入 LANG=en 后启动（默认 GUI）
repo-manager -dry-run 6            # 预演清理所有项目缓存（不删除任何文件）
repo-manager -h                    # 查看帮助与构建信息
```

---

## 🛠️ 排障与维护

- 使用中遇到问题（启动失败、配置不生效、清理结果不符预期、门禁报错）→ [排障手册](./docs/TROUBLESHOOTING.md)
- 提问前自查与有效 Issue 应含信息 → [支持指南](./.github/SUPPORT.md)；报缺陷用 [Issue 模板](https://github.com/sutchan/RepoManager/issues/new/choose)
- **安全漏洞不要公开提 Issue** → [安全策略](./.github/SECURITY.md)
- 维护者操作（GitHub About 推荐设置、命名现状、发布校验、文档治理）→ [维护者手册](./docs/MAINTAINERS.md)
- 产品名 `Repo Manager` 已统一到 GUI 标题、CLI 帮助、i18n 字典、Go 模块与二进制 `repo-manager`、日志 `repo-manager.log`、Release 资产与 Git 远程 `sutchan/RepoManager`；命名对照与展示位见[维护者手册 §1](./docs/MAINTAINERS.md)

---

## ❓ 常见问题 FAQ

**Q：功能 1 会影响现有项目运行吗？**
A：启用全局 node_modules 后，项目不再生成本地 `node_modules`，统一从全局读取。请确保已先运行功能 5 安装所需依赖，否则项目会因找不到模块而报错。

**Q：功能 3 禁用 `.next` 缓存后热更新变慢？**
A：是的，禁用缓存会略微影响 Next.js 热更新速度。开发完成后可运行功能 4 恢复。

**Q：功能 6 / 7 删除后如何恢复？**
A：清理是不可逆的。删除 `node_modules` 后需重新 `npm install`；建议仅在确认不再开发或已备份时执行。

**Q：默认扫描哪个目录？**
A：默认扫描 Windows 下的 `E:\Github\`（macOS / Linux 为 `~/Github`），可通过 GUI 输入框、`config.ini` 的 `REPO_ROOT`，或与功能号同用的 CLI 位置路径修改。

**Q：支持含空格的仓库路径吗？**
A：支持。Go 版目录枚举用 `os.ReadDir`、路径一律用 `path/filepath` 拼接，不经过 shell 字符串拼接；命令行传参时仍需按终端规则给路径加引号。

**Q：为什么 `repo-manager E:\Github` 没有切换目录？**
A：不带功能号时程序启动 GUI 并忽略位置路径。请写成 `repo-manager <功能号> <路径>`（如 `repo-manager 6 E:\Github`），或在 GUI 里选择目录后保存。

**Q：有交互式文本菜单吗？**
A：没有。v1.x 批处理脚本提供过文本菜单；v2.0.0 起只有 Fyne GUI 与「功能号 CLI」两种模式。

**Q：怎么确认哪些文件会被删除？**
A：加 `-dry-run` 或在 GUI 勾选「预演模式」，只输出将执行的操作，不修改任何文件与环境变量；预演下危险操作跳过二次确认。

**Q：执行日志在哪里？**
A：可执行文件同目录的 `repo-manager.log`，CLI 与 GUI 均追加写入并带时间戳；日志可能含本机路径，公开提交前请脱敏。更多症状见[排障手册](./docs/TROUBLESHOOTING.md)。

---

## ⚠️ 注意事项

1. **首次使用建议先运行 `[5]` 安装全局依赖**
2. **功能 3 会影响 Next.js 热更新速度**，开发完成后可恢复
3. **功能 6 删除后需要重新 `npm install`**，建议在确认不再开发时使用
4. **默认扫描 `E:\Github\` 目录**（Unix 为 `~/Github`），可通过 GUI、`config.ini` 或 CLI 位置路径修改
5. **删除操作不可逆**，执行前请确认已备份重要文件
6. **默认只扫描仓库根目录的直接子目录**；需要覆盖 monorepo 深层项目时显式开启 `-recursive` / GUI 复选框 / `RECURSIVE=true`（深度 3 层，见 [OPENSPEC.md](./docs/OPENSPEC.md) §3.0 / §7.3 L3）

---

## 📚 文档体系与开发规范 (docs/)

所有规格与规范文档统一归档在 [`docs/`](./docs/README.md)，根 README 只作使用者入口：

| 文档 | 内容 |
|------|------|
| [docs/README.md](./docs/README.md) | 文档中心：文档地图、职责边界、按角色阅读路径 |
| [docs/OPENSPEC.md](./docs/OPENSPEC.md) | 产品功能规格：7 项功能行为契约、交互、兼容性、已知限制 |
| [docs/CONTRIBUTING.md](./docs/CONTRIBUTING.md) | 开发规范：编码风格、命名约定、目录组织、提交信息格式、分支管理、代码审查流程、版本规则、发布流程、术语表 |
| [docs/DESIGN-SYSTEM.md](./docs/DESIGN-SYSTEM.md) | UI 设计规范：设计令牌、组件库、交互、响应式、a11y、规范与实现差距清单 |
| [docs/tokens.json](./docs/tokens.json) | 机器可读设计令牌 |
| [docs/config.example.ini](./docs/config.example.ini) | 运行配置示例模板（`config.ini` 不入库） |
| [docs/TROUBLESHOOTING.md](./docs/TROUBLESHOOTING.md) | 使用与构建排障：症状、原因、处理与门禁失败定位 |
| [docs/MAINTAINERS.md](./docs/MAINTAINERS.md) | 维护者手册：命名现状、仓库设置、发布校验、文档治理、已知差距 |
| [docs/tasks.md](./docs/tasks.md) / [tasks-done.md](./docs/tasks-done.md) | 进行中 / 待办任务清单与已完成任务归档 |
| [landing/README.md](./landing/README.md) | Web 展示项目落地页目录：用途、边界与内容约定（页面文件规划中） |
| [.github/CODE_OF_CONDUCT.md](./.github/CODE_OF_CONDUCT.md) | 行为准则 |
| [.github/SECURITY.md](./.github/SECURITY.md) | 安全策略与漏洞私密报告 |
| [.github/SUPPORT.md](./.github/SUPPORT.md) | 支持与提问指南 |
| [.github/PULL_REQUEST_TEMPLATE.md](./.github/PULL_REQUEST_TEMPLATE.md) | PR 审查清单模板 |

### 社区与贡献

- 参与贡献前请阅读[开发规范](./docs/CONTRIBUTING.md)与[行为准则](./.github/CODE_OF_CONDUCT.md)。
- 报缺陷 / 提建议：使用 [Issue 模板](https://github.com/sutchan/RepoManager/issues/new/choose)，提问前先看[支持指南](./.github/SUPPORT.md)与[常见问题 FAQ](#-常见问题-faq)。
- **安全漏洞不要公开提 Issue**，请按[安全策略](./.github/SECURITY.md)走私密渠道；本工具会删除文件，相关边界务必先读。
- 社区行为问题按[行为准则](./.github/CODE_OF_CONDUCT.md)举报。

`prototype/` 现仅保留[原型目录说明](./prototype/README.md)，不再包含代码或规范正文。

### 参考界面应用（已弃用）

根 `src/` 是 v1.x 的 Next.js + React + TypeScript + Tailwind v4 参考界面（纯展示，不执行真实清理；真实清理由 Go 版实现）。自 v2.0.0 起已弃用，仅保留作历史参考：

- **单一数据源** `src/lib/data/`，文案统一经 `src/lib/i18n.ts` 简中 / 繁中 / 英文三语切换
- **设计令牌** `src/app/globals.css` 的 CSS 变量，与 `docs/tokens.json` 对齐
- **语义化 id** 主要容器/区块均带 `id`（如 `repo-manager-app`、`main-tabs`、`repo-card-<name>`）
- **页脚** 展示 `vX.Y.Z · build yymmddhhmm`（10 位 build 编号，由 `next.config.ts` 从 `package.json` 注入）

本地预览与全部可用命令（含 `dev` / `type-check` / `lint` / `check:docs` / `check:root-clean` / `build`）见[文档中心 §4 常用命令](./docs/README.md)；Go 产品本体的构建与门禁见[开发规范 §2](./docs/CONTRIBUTING.md)。

---

## 📂 目录结构

```
RepoManager/
├── .github/                           # Community Health Files
│   ├── workflows/                     # ci.yml（门禁）+ release.yml（push vX.Y.Z tag 触发，release job 依赖 gate + build）
│   ├── CODE_OF_CONDUCT.md / SECURITY.md / SUPPORT.md
│   ├── FUNDING.yml / CODEOWNERS / PULL_REQUEST_TEMPLATE.md
│   └── ISSUE_TEMPLATE/                # config.yml + bug_report.yml + feature_request.yml
├── docs/                              # 所有规范与规格文档归档
│   ├── README.md                      # 文档中心（导航 + 职责边界）
│   ├── OPENSPEC.md                    # 产品功能规格
│   ├── CONTRIBUTING.md                # 开发规范
│   ├── DESIGN-SYSTEM.md               # UI 设计规范
│   ├── TROUBLESHOOTING.md             # 使用与构建排障
│   ├── MAINTAINERS.md                 # 维护者手册
│   ├── tokens.json                    # 设计令牌
│   ├── config.example.ini             # 运行配置示例模板
│   └── tasks.md / tasks-done.md       # 任务清单与完成归档
├── go.mod / go.sum                    # Go 模块与依赖（Fyne v2）
├── cmd/repo-manager/                  # 产品入口：main.go + gui.go（默认）/ headless.go（nogui 占位）
├── internal/                          # Go 源码：buildinfo / config / engine / i18n / logs / platform / ui（7 个包）
├── scripts/                           # build.ps1 / build.sh（交叉编译）+ check-doc-links.mjs / check-root-clean.mjs（门禁）
├── .tmp/                              # 会话级临时文件唯一落点（脚本 / 探针 / 日志 / 中间产物，用完即清）
├── reports/                           # 工具正式产物目录（审计报告 / 覆盖率 / 测试结果）
├── RepoCleaner.bat                    # v1.x 旧版 Windows 脚本（已弃用）
├── config.ini                         # 运行配置（REPO_ROOT / LANG / RECURSIVE / EXTRA_CLEAN_DIRS，exe 同目录、本机生成不入库）
├── lang.ini                           # v1.x 批处理语言文案（历史保留，Go 程序不读取）
├── README.MD / CHANGELOG.MD / LICENSE # 入口 / 变更历史 / 许可证
├── RepoManager.code-workspace          # VS Code 工作区配置（禁止删除；命名现状见 docs/MAINTAINERS.md §1）
├── .nvmdrc                            # 本机 Node 版本标记
├── package.json                       # 参考界面依赖与版本同步位（产品唯一权威源在 internal/buildinfo/version.go）
├── next.config.ts                     # 参考界面 Next.js 配置（注入版本 + 10 位 build 编号）
├── tsconfig.json                      # TypeScript 配置（@/* → ./src/*）
├── postcss.config.mjs                 # Tailwind v4 PostCSS 配置
├── eslint.config.mjs                  # ESLint 配置
├── public/                            # 参考界面静态资源
├── src/                               # 已弃用参考界面应用（v1.x 遗留）
│   ├── app/                           # layout.tsx / page.tsx / globals.css
│   ├── components/                    # icons.tsx + ui/ 基础组件
│   └── lib/                           # i18n.ts + utils.ts + data/（单一数据源）
├── landing/                           # Web 展示项目落地页（README.md 为目录说明，页面文件规划中）
├── prototype/README.md                # 原型目录说明（无代码、无规范正文）
└── .gitignore
```

> **临时文件约定**：一次性脚本、探针、调试日志、中间产物一律写 `.tmp/`；工具正式产物（报告 / 覆盖率 / 测试结果）写 `reports/`。目录忽略规则以 `.gitignore` 实况为准，根目录白名单由 `npm run check:root-clean` 强制校验；新增根目录工程文件须同时核对二者。详见 [开发规范 §3.2](./docs/CONTRIBUTING.md) 与 [维护者手册 §5](./docs/MAINTAINERS.md)。

---

## 📄 许可证

[MIT License](./LICENSE)

---

## 📝 更新日志

详细更新内容请查看 [CHANGELOG.MD](./CHANGELOG.MD)。

---

<div align="center">

Repo Manager v2.3.3 · MIT License · [文档中心](./docs/README.md)

</div>
