# RepoCleaner 开放规格说明书 (OpenSpec) v2.0.2

> **本文定位**：产品功能规格，定义「做什么」——7 项功能的行为契约、交互方式、环境与兼容性。
> 开发规范（编码 / 提交 / 分支 / 审查）见 [CONTRIBUTING.md](./CONTRIBUTING.md)；UI 设计规范见 [DESIGN-SYSTEM.md](./DESIGN-SYSTEM.md)；变更历史见 [CHANGELOG.MD](../CHANGELOG.MD)；文档导航见 [docs/README.md](./README.md)。
> 本文与代码冲突时以代码为准，并按 `CONTRIBUTING.md` §10 同步修正本文。

---

## 1. 项目概述

| 项 | 内容 |
|----|------|
| 项目名称 | RepoCleaner - 项目清洁助手 |
| 产品本体 | `repo-cleaner`（Go 编译的跨平台可执行文件，Windows / macOS / Linux 单二进制；构建需 C 工具链，运行依赖系统 OpenGL）；`RepoCleaner.bat` 为 Windows 旧版兼容脚本（**已弃用**） |
| 参考界面应用 | 仓库根 `src/` 的 Next.js 应用为 **v1.x 遗留、已弃用的参考界面**（纯展示，不执行真实清理）；v2.0.0 起实际 UI 为 **Go + Fyne 桌面 GUI** |
| 核心功能 | 一键管理仓库根目录下的多个前端项目：清理冗余缓存、统一管理全局 `node_modules`、节省磁盘空间 |
| 目标用户 | 前端开发者、全栈开发者、需要管理多个 Node.js 项目的团队 |
| 支持平台 | Windows / macOS / Linux（单一代码库，交叉编译为各平台单二进制） |
| 支持项目类型 | Next.js、React、Vue、Vite、常规 Node.js 项目 |
| 许可证 | MIT（见 [LICENSE](../LICENSE)） |

> **目录概念**：「仓库根目录」指 `REPO_ROOT` 指向的、包含多个项目的父目录（默认 `E:\Github`）；「项目根目录」指其中一个项目目录，是 `REPO_ROOT` 的直接子目录。二者不可混用。

---

## 2. 架构与模块化

### 2.1 应用分层架构（Go）

`repo-cleaner` 采用分层设计，所有操作系统差异收口到 `internal/platform`，业务与界面各自独立：

| 层 / 包 | 职责 |
|---------|------|
| `cmd/repo-cleaner/main.go` | 程序入口：解析参数、目录解析、启动 Fyne GUI |
| `internal/platform` | **操作系统差异抽象**（关键）：目录扫描、删除、隐藏/系统/只读属性、环境变量持久化、默认 `REPO_ROOT`，按 `build tag` 分 `platform_windows.go` / `platform_unix.go` 实现 |
| `internal/engine` | 功能 1 ～ 功能 7 的**平台无关**业务逻辑（扫描、写 `.npmrc`、清缓存），只调用 `platform` 接口 |
| `internal/config` | `config.ini` 读写（去重键、保留未知行）、默认 `REPO_ROOT` / `LANG` |
| `internal/i18n` | 中文 / 英文文案字典（唯一来源） |
| `internal/ui` | Fyne 桌面 GUI：主菜单、语言切换、二次确认、进度与结果展示 |

> 新增平台只需在 `internal/platform` 增加对应 `build tag` 实现，`engine` 与 `ui` 零改动。

### 2.2 关键设计决策

| 决策 | 原因 | 当前实现 |
|------|------|----------|
| 平台差异全部收口到 `internal/platform` | 避免业务 / UI 散落 `runtime.GOOS` 分支 | 接口 + 按 `build tag` 的 Windows / Unix 实现 |
| 业务与平台解耦 | 跨平台复用同一份功能逻辑 | `engine` 只依赖 `platform` 接口 |
| 语言文案集中在 `internal/i18n` | 单一数据源，避免散落硬编码 | zh / en 字典，GUI 与日志统一引用 |
| 危险操作二次确认 | 删除不可逆，防止误删 | 功能 1 / 6 / 7 均需确认：GUI 确认弹窗、CLI 输入 `Y` |
| UI 选 Fyne（Go + 系统 OpenGL） | 单二进制、跨平台一致 | 构建需 C 工具链；桌面窗口复用既有设计令牌配色 |

### 2.3 已弃用的参考界面应用（Next.js，v1.x 遗留）

仓库根 `src/` 下的 Next.js 应用是 **v1.x 的参考界面**（纯展示 dashboard，不执行真实清理），自 v2.0.0 起**已弃用**，仅保留作历史参考，不再随产品演进；实际 UI 为 Go + Fyne（见 §2.1）。

> 仍可 `npm install && npm run dev` 启动作 Web 预览，但属遗留路径。`src/lib/data/repos.ts` 中的扫描结果与容量数字是**演示数据（mock）**，不得作为事实引用。

---

## 3. 功能规格

### 3.0 功能名称对照表（全仓库文档统一使用）

| 编号 | 统一名称 | 批处理标签 |
|------|----------|------------|
| 功能 1 | 启用全局 node_modules | `:GLOBAL_NODE` |
| 功能 2 | 重置 Node 配置 | `:RESET_NODE` |
| 功能 3 | 清理并禁用 .next 缓存 | `:CLEAN_NEXT` |
| 功能 4 | 恢复 Next.js 缓存 | `:RESET_NEXT` |
| 功能 5 | 安装常用全局依赖 | `:INSTALL_DEPS` |
| 功能 6 | 清理所有项目缓存 | `:CLEAN_ALL_CACHE` |
| 功能 7 | 清理 Vite/构建缓存 | `:CLEAN_VITE_CACHE` |

> 菜单上显示的短文案以 `internal/i18n` 文案字典为准；文档只使用上表名称。
> v2.0.0 起各功能由 `internal/engine` 的 `EnableGlobalNodeModules` / `ResetNodeConfig` / `CleanNextCache` / `RestoreNextCache` / `InstallGlobalDeps` / `CleanAllCaches` / `CleanViteCache` 实现；上表「批处理标签」列仅为 v1.x 历史对照。

所有功能的作用域统一为：**遍历 `REPO_ROOT` 的直接子目录**（`internal/platform.ListProjectDirs` 用 `os.ReadDir` 枚举），不递归进入更深层级。

### 3.1 功能 1：启用全局 node_modules

- **描述**：配置系统环境变量 `NODE_PATH`，禁用项目本地 `node_modules` 生成。
- **操作步骤**：
  1. 显示警告并要求二次确认（GUI 确认弹窗 / CLI 输入 `Y`；取消则**不执行任何删除**）。
  2. `npm root -g` 获取全局 npm 模块路径。
  3. 调用 `internal/platform.SetUserEnv` 写入用户级 `NODE_PATH`（Windows `setx`；Unix 写 `~/.profile`）。
  4. 为所有含 `package.json` 的项目根目录写入 `.npmrc`。
  5. 删除这些项目现有的本地 `node_modules` 目录。
- **`.npmrc` 实际生成内容**（5 行）：
  ```ini
  prefix=<全局模块路径>
  cache=<全局模块路径>\.npm-cache
  tmp=<%TEMP% 的实际值>
  global=true
  prefer-global=true
  ```
- **已知限制**：`tmp=` 行的值**已加引号**（v2.0.0 修复 L1），兼容含空格的 `%TEMP%` 路径。
- **预期结果**：项目共享全局依赖；此后仍需先运行功能 5 安装依赖，否则项目会因找不到模块而报错。

### 3.2 功能 2：重置 Node 配置

- **操作步骤**：调用 `internal/platform.RemoveUserEnv`（Windows 清空 `setx` 并删除 `HKCU\Environment` 注册表值；Unix 从 `~/.profile` 过滤 `export NODE_PATH=` 行）；删除所有项目根目录的 `.npmrc`。
- **预期结果**：npm 恢复默认行为。

### 3.3 功能 3：清理并禁用 .next 缓存

- **操作步骤**：
  1. 遍历 `REPO_ROOT` 直接子目录，识别含 `next.config.js` **或** `next.config.ts` 的 Next.js 项目。
  2. `internal/platform.RemoveAll` 删除现有的 `.next` 缓存目录。
  3. 写入一个同名的 `.next` **文件**，阻止 Next.js 再创建 `.next` 目录。
  4. 设置隐藏属性（Windows：`attrib +h +s +r`；macOS：`chflags hidden`；Linux：点前缀 + 只读位），使该文件在文件管理器中默认不可见。
- **预期结果**：节省磁盘空间，但会影响 Next.js 热更新速度。

### 3.4 功能 4：恢复 Next.js 缓存

- **操作步骤**：
  1. 与功能 3 相同的项目识别规则。
  2. **先清除**隐藏/系统/只读属性（Windows：`attrib -h -s -r`；macOS：`chflags nohidden`；Linux：移除只读位）。
  3. 删除该 `.next` 文件。
- **实现要点**：必须先清除属性再删除，否则删除会静默失败（沿用 v1.2.3 修复的约束；v2.0.0 在 `internal/platform` 中按平台实现）。
- **预期结果**：Next.js 下次构建时重新生成 `.next` 目录，恢复正常编译与热更新。

### 3.5 功能 5：安装常用全局依赖

按 5 组依次执行 `npm install -g`：

| 组 | 包 |
|----|----|
| React 生态 | `react` `react-dom` `next` |
| Vue 生态 | `vue` `@vue/cli` |
| HTTP 库 | `axios` `express` |
| 包管理器 | `pnpm` `yarn` |
| 开发工具 | `cross-env` `dotenv` `nodemon` `pm2` `rimraf` `mkdirp` |

- **预期结果**：依赖安装到全局，项目可通过 `NODE_PATH` 直接引用。
- **注意**：此功能需要网络；单组失败不会中断后续组（`runNpmInstall` 返回错误后仅记录日志，对应 L4）。

### 3.6 功能 6：清理所有项目缓存

- **前置**：显示警告并要求二次确认（GUI 确认弹窗 / CLI 输入 `Y`）。
- **清理项**（每个 `REPO_ROOT` 直接子目录）：
  - 目录：`node_modules`、`.next`、`.nuxt`、`dist`、`build`
  - 文件：`package-lock.json`、`yarn.lock`、`pnpm-lock.yaml`
- **预期结果**：释放磁盘空间。删除不可逆，恢复需重新 `npm install`。

### 3.7 功能 7：清理 Vite/构建缓存

- **前置**：显示警告并要求二次确认（GUI 确认弹窗 / CLI 输入 `Y`）。
- **清理项**：`.vite`、`.turbo`、`.nuxt`、`dist`、`build`、`coverage`、`.nyc_output`
- **预期结果**：清理构建缓存，适合 Vite / Turborepo 项目。

> 功能 6 与功能 7 有重叠项（`dist`、`build`、`.nuxt`），属预期设计：功能 6 面向「全量清理」，功能 7 面向「仅构建缓存」。

---

## 4. 用户交互规格

### 4.1 语言切换

| 项 | 内容 |
|----|------|
| 支持语言 | 中文（`zh`）、英文（`en`） |
| 切换方式 | GUI 顶部语言按钮；命令行 `-lang en\|zh`；`config.ini` 的 `LANG=` |
| 默认语言 | `zh` |
| 文案来源 | `internal/i18n`：界面字典 `T()` + 执行日志字典 `Logs()`；`lang.ini` 为 v1.x 历史保留，Go 程序不读取 |

> `-lang` 只写 `config.ini`，**禁止整文件覆盖 `lang.ini`**（会清空文案节）。

### 4.2 目录选择

优先级：命令行路径参数 → `config.ini` 的 `REPO_ROOT` → 平台默认目录（由 `internal/platform.DefaultRepoRoot()` 提供）。

GUI 提供 3 种目录来源：使用当前目录、手动输入完整路径（带目录选择按钮）、使用默认目录（Windows：`E:\Github`；macOS / Linux：`~/Github`）。输入路径不存在时报错并允许重新输入；选择保存时以「读旧 → 过滤目标键 → 补写 → 回写」方式更新 `config.ini`（保留未知行与顺序）。

### 4.3 命令行参数

| 参数 | 说明 | 示例 |
|------|------|------|
| `-lang en\|zh` | 设置语言并写入 `config.ini` | `repo-cleaner -lang en` |
| `-h` / `--help` | 显示帮助信息 | `repo-cleaner -h` |
| `1`-`7` | 直接运行对应功能（命令行模式，跳过 GUI） | `repo-cleaner 1` |
| 路径 | 指定仓库根目录（支持含空格路径，需引号包裹） | `repo-cleaner "D:\My Projects"` |
| 数字 + 路径 | 为指定目录直接运行功能 | `repo-cleaner 5 E:\Projects` |

### 4.4 主菜单（中文默认）

> v2.0.0 起主菜单为 **Fyne 桌面窗口**（按钮列表，非控制台文本）；下方文本图为 v1.x 菜单对照示意，功能编号与文案保持一致。

```
======================================================================
               RepoCleaner - 项目清洁助手
======================================================================

            目标目录: E:\Github

         [1] 启用全局 node_modules     - 无需本地 node_modules
         [2] 重置 Node 配置            - 恢复默认设置

         [3] 清理 .next 缓存           - 禁用 Next.js 缓存
         [4] 恢复 Next.js 缓存         - 启用正常编译

         [5] 安装全局依赖              - React/Vue/Next/axios/工具
         [6] 清理所有项目缓存          - 删除 node_modules/.next/dist
         [7] 清理 Vite/构建缓存        - 删除 .vite/dist/.turbo

         [L] English / 中文
         [0] 退出

======================================================================
```

### 4.5 错误处理

- 路径不存在：GUI 弹窗提示（对齐 `PATH_ERROR` 文案），关闭后返回目录选择。
- 危险操作：功能 1 / 6 / 7 均需二次确认（GUI 确认弹窗；CLI 输入 `Y`），取消不执行任何清理。
- 每个功能完成后在结果区显示该功能的成功消息；命令行模式等同文本输出。

### 4.6 应用页脚与构建标识

- 应用页面的**页脚区**是版本号与 build 编号的唯一展示位置；页眉 / 顶部栏不得出现版本号。
- 展示格式 `vX.Y.Z · build yymmddhhmm`，示例 `v1.2.3 · build 2611051259`；build 编号为 10 位（2 位年份 + 月日时分，24 小时制）。
- Go GUI 页脚文本由 `internal/buildinfo.Footer()` 生成；`scripts/build.ps1` / `scripts/build.sh` 在构建期通过 `-ldflags -X` 注入 `BuildID`，未注入时由 `vcs.time` 推导，仍缺失则仅显示版本号。
- 已弃用 Web 参考界面由 `next.config.ts` 注入 `NEXT_PUBLIC_APP_VERSION` / `NEXT_PUBLIC_BUILD_ID`，在 `<footer id="app-footer">` 展示；注入缺失时回退仅显示版本号。
- 版本权威源与同步清单见 [`CONTRIBUTING.md` §8.2](./CONTRIBUTING.md)，变更历史见 [CHANGELOG.MD](../CHANGELOG.MD)。

---

## 5. 技术规格

### 5.1 环境要求

| 项 | 要求 |
|----|------|
| 操作系统 | Windows / macOS / Linux（单一代码库，交叉编译为各平台单二进制） |
| 运行环境 | 编译后的 `repo-cleaner` 单二进制；无需额外运行时安装，运行依赖系统 OpenGL（Fyne） |
| 构建环境 | Go ≥ 1.25（`go` 工具链）；交叉编译命令见 `CONTRIBUTING.md` §2 |
| 依赖软件 | Node.js ≥ 18 与 npm（功能 1 / 5 需要，运行时检测） |
| 已弃用参考界面 | Node.js ≥ 18、npm、Next.js 16（见 `package.json`，仅遗留 Web 预览） |

### 5.2 文件结构

```
RepoCleaner/
├── .github/                           # Community Health Files
│   ├── CODE_OF_CONDUCT.md / SECURITY.md / SUPPORT.md
│   ├── FUNDING.yml / CODEOWNERS / PULL_REQUEST_TEMPLATE.md
│   └── ISSUE_TEMPLATE/                # config.yml + bug_report.yml + feature_request.yml
├── docs/                              # 所有规范与规格文档归档
│   ├── README.md                      # 文档中心（导航 + 职责边界）
│   ├── OPENSPEC.md                    # 本文：产品功能规格
│   ├── CONTRIBUTING.md                # 开发规范
│   ├── DESIGN-SYSTEM.md               # UI 设计规范
│   └── tokens.json                    # 设计令牌
├── public/                            # 参考界面应用静态资源
├── scripts/check-doc-links.mjs        # 文档链接校验（npm run check:docs）
├── src/                               # 参考界面应用代码
│   ├── app/                           # layout.tsx / page.tsx / globals.css / favicon.ico
│   ├── components/                    # icons.tsx + ui/
│   └── lib/                           # i18n.ts / utils.ts / data/
├── prototype/README.md                # 原型目录说明（无代码、无规范正文）
├── go.mod / go.sum                    # Go 模块定义与依赖锁定（Fyne）
├── cmd/repo-cleaner/main.go           # 程序入口
├── internal/                          # Go 源码：platform / engine / config / i18n / ui
├── scripts/build.ps1 / build.sh       # 跨平台交叉编译脚本
├── RepoCleaner.bat                    # v1.x 旧版 Windows 脚本（已弃用，保留兼容）
├── config.ini                         # REPO_ROOT / LANG
├── lang.ini                           # [en] / [zh] 文案参考节（历史保留）
├── README.MD / CHANGELOG.MD / LICENSE # 入口、变更历史、许可证
├── RepoCleaner.code-workspace
├── package.json / package-lock.json / tsconfig.json / next.config.ts
├── postcss.config.mjs / eslint.config.mjs / next-env.d.ts
└── .gitignore
```

### 5.3 配置文件格式

**`config.ini`**（Go 程序实际读取，位于可执行文件同目录）：

```ini
REPO_ROOT=E:\Github
LANG=zh
```

**`lang.ini`**：首部 `LANG=` 记录当前语言选择；`[en]` / `[zh]` 两节保存中英文文案参考。Go 程序**不从该节读取文案**（文案在 `internal/i18n` 中定义），此文件为 v1.x 历史保留。`-lang` 参数不得覆盖本文件。

### 5.4 编码要求

- 所有文本文件：**UTF-8**。
- Go：`gofmt` 格式化（Tab 缩进）；文件头注释 `// <相对路径> vX.Y.Z`。
- 缩进：Go 用 `gofmt` 默认（Tab）；TS/JS/CSS/JSON 2 空格；`.bat` 4 空格。详见 `CONTRIBUTING.md` §4。

### 5.5 代码统计（历史数据，仅供追溯）

| 指标 | 原始 | v1.1.0 重构后 |
|------|------|---------------|
| 批处理总行数 | 675 | ~630 |
| 命令行数字分发标签 | 7 个重复 `ARG_RUN_*` | 1 个 `:RUN_ARG_FUNCTION` |
| 语言变量定义 | 2 处重复 | `:LOAD_LANGUAGE` 统一管理 |

> 以上为 v1.1.0 重构时的历史统计，**不代表当前行数**，请勿据此判断现状。

---

## 6. 数据规格

### 6.1 环境变量

| 变量名 | 说明 | 来源 |
|--------|------|------|
| `NODE_PATH` | 全局模块路径 | 功能 1 写入用户环境（Windows：`setx`；macOS / Linux：写 shell profile）；功能 2 删除 |
| `TEMP` | 临时目录 | 系统环境变量，写入 `.npmrc` 的 `tmp=` |
| `REPO_ROOT` / `LANG` | 脚本内部变量 | 命令行 / `config.ini` / 交互输入 |

### 6.2 生成 / 修改的文件

| 文件 | 位置 | 用途 |
|------|------|------|
| `.npmrc` | 各项目根目录 | 配置 npm 使用全局模块（功能 1 写、功能 2 删） |
| `.next` | 各 Next.js 项目根目录 | 功能 3 创建为带隐藏属性的**文件**以阻止缓存；功能 4 清除属性后删除（属性操作按平台实现） |
| `config.ini` | 可执行文件同目录 | 保存 `REPO_ROOT` / `LANG` |

---

## 7. 兼容性规格

### 7.1 路径处理

- 支持带空格的路径：目录枚举用 `os.ReadDir`，路径统一经 `path/filepath` 拼接，无命令行拼接注入风险。
- 支持相对路径与绝对路径。
- `platform.NormalizeRoot` 去除空白与末尾分隔符并转绝对路径。

### 7.2 文件属性

- 功能 3：设置隐藏属性（Windows：`attrib +h +s +r`；macOS：`chflags hidden`；Linux：点前缀 + 只读位）。
- 功能 4：**必须先清除全部隐藏/系统/只读属性**再删除（Windows：`attrib -h -s -r` 后删除；macOS：`chflags nohidden` 后删除；Linux：移除只读位后删除）。

### 7.3 已知限制

| 编号 | 限制 | 影响 |
|------|------|------|
| L1 | ~~`.npmrc` 的 `tmp=` 值未加引号~~ | **v2.0.0 已修复**：`tmp=` 值加引号，兼容含空格的 `TEMP` 路径 |
| L2 | 功能 1 设置 `NODE_PATH` 后需重开终端才对新进程完全生效 | 当前会话已同步设置，新终端需重开 |
| L3 | 目录扫描仅覆盖 `REPO_ROOT` 的直接子目录，不递归 | monorepo 深层项目不会被处理 |
| L4 | 功能 5 不检查 `npm install -g` 的退出码 | 单组安装失败仍显示成功消息 |
| L5 | ~~功能 6 / 7 使用 `pause` 而非 `Y/N` 二次确认~~ | **v2.0.0 已修复**：GUI 确认弹窗 + CLI `Y/N` |

> 记录限制是为了让规格与实现一致；修复限制时须同步更新本表与 `CHANGELOG.MD`。

---

## 8. 变更记录

**唯一来源：[CHANGELOG.MD](../CHANGELOG.MD)**（按版本倒序记录全部发布版本的完整变更明细）。

本文不再重复维护版本历史，避免两份记录漂移。

---

## 9. 验证与质量门禁

### 9.1 仓库级自动门禁（可在任意平台执行）

在仓库根目录依次执行，四项全部通过方可合并：

```bash
gofmt -l .                                  # 输出应为空
go vet ./cmd/... ./internal/...             # 静态检查
go test ./cmd/... ./internal/...            # 单元 / 临时目录集成测试
go build ./cmd/...                          # 编译产品（需 OpenGL C 工具链）
npm run check:docs                          # 校验所有 Markdown 相对链接（含大小写）
```

预期：`gofmt` 无输出、`go vet` 无告警、`go test` 全部通过、`go build` 编译成功、`check:docs` 报告失效链接 0 个。

无 OpenGL C 工具链时，上述 Go 命令统一加 `-tags nogui`（`go build -tags nogui ./...`），GUI 包按 build tag 排除；不得以此为由跳过门禁。

> 删除类行为通过临时目录集成测试覆盖（`internal/engine` / `internal/platform` 的 `*_test.go`）；`RepoCleaner.bat` 为已弃用旧版，不再纳入自动门禁。

### 9.2 GUI 手动冒烟（跨平台）

`repo-cleaner` 的 GUI 交互无自动化测试，发版前须在**隔离的临时目录**（放 1～2 个含 `package.json` 的假项目）手动验证：帮助信息、`-lang` 写入、语言切换、功能 1 的确认/取消分支、功能 3 → 4 的 `.next` 文件生成与删除、含空格路径、路径不存在重试。完整清单见 [`CONTRIBUTING.md` §9.3](./CONTRIBUTING.md)。

> 严禁在真实仓库根目录执行删除类验证。

### 9.3 覆盖范围说明

- Go 模块已配置 `go test ./...`，覆盖 `internal/engine`、`internal/config`、`internal/i18n` 的纯逻辑与 `internal/platform` 的临时目录集成测试；这是当前的全部自动检查。
- 历史版本文档中的「功能 1～7 ✅ 通过」是当时的手工验证结论，不作为当前自动化覆盖证据。

---

## 10. 扩展计划

- [ ] 支持更多框架缓存清理（Angular、Svelte）
- [ ] 添加 dry-run 模式预览待删除文件
- [ ] 支持自定义清理规则
- [ ] 添加日志记录功能
- [ ] 功能 5 检查 `npm install -g` 退出码并汇总失败项（对应 L4）
- [x] ~~功能 6 / 7 改为 `Y/N` 二次确认~~（v2.0.0 已完成，见 L5）
- [ ] 支持递归扫描 monorepo 深层项目（对应 L3）
- [ ] 添加详细进度显示

---

## 11. 许可证

MIT License，详见 [LICENSE](../LICENSE)。
