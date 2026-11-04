# RepoCleaner 开放规格说明书 (OpenSpec) v1.2.5

> **本文定位**：产品功能规格，定义「做什么」——7 项功能的行为契约、交互方式、环境与兼容性。
> 开发规范（编码 / 提交 / 分支 / 审查）见 [CONTRIBUTING.md](./CONTRIBUTING.md)；UI 设计规范见 [DESIGN-SYSTEM.md](./DESIGN-SYSTEM.md)；变更历史见 [CHANGELOG.MD](../CHANGELOG.MD)；文档导航见 [docs/README.md](./README.md)。
> 本文与代码冲突时以代码为准，并按 `CONTRIBUTING.md` §10 同步修正本文。

---

## 1. 项目概述

| 项 | 内容 |
|----|------|
| 项目名称 | RepoCleaner - 项目清洁助手 |
| 产品本体 | `RepoCleaner.bat`，Windows 批处理脚本（cmd），无第三方运行时依赖 |
| 参考界面应用 | 仓库根 `src/` 下的 Next.js 16 应用，是本文功能规格的**可视化参考**，不执行实际清理 |
| 核心功能 | 一键管理仓库根目录下的多个前端项目：清理冗余缓存、统一管理全局 `node_modules`、节省磁盘空间 |
| 目标用户 | 前端开发者、全栈开发者、需要管理多个 Node.js 项目的团队 |
| 支持项目类型 | Next.js、React、Vue、Vite、常规 Node.js 项目 |
| 许可证 | MIT（见 [LICENSE](../LICENSE)） |

> **目录概念**：「仓库根目录」指 `REPO_ROOT` 指向的、包含多个项目的父目录（默认 `E:\Github`）；「项目根目录」指其中一个项目目录，是 `REPO_ROOT` 的直接子目录。二者不可混用。

---

## 2. 架构与模块化

### 2.1 批处理脚本模块划分

`RepoCleaner.bat` 按标签（label）分节，各节职责单一：

| 标签 / 区块 | 职责 |
|-------------|------|
| 配置加载（文件开头） | 读取 `config.ini` / `lang.ini`，确定语言 |
| `:LOAD_LANGUAGE` | 定义中英文全部文案变量，按 `CURRENT_LANG` 覆盖英文值 |
| `:PARSE_ARGS` | 解析 `-lang` / `-h` / `--help` / 数字 / 路径参数 |
| `:SHOW_HELP` | 输出帮助信息 |
| `:RUN_ARG_FUNCTION` | 统一分发数字参数（`ARG_MODE`），并接收可选路径 |
| `:CHECK_CONFIG` / `:LANG_MENU` / `:INTERACTIVE_INPUT` | 语言选择、目录选择、配置保存 |
| `:ROOT_SET` / `:MENU` | 规范化根目录、显示主菜单并处理输入 |
| `:GLOBAL_NODE` ～ `:CLEAN_VITE_CACHE` | 功能 1 ～ 功能 7 实现 |

> 规格文档不再标注源码行号范围（行号随修改漂移，属易失真信息）；定位请直接搜索标签名。

### 2.2 关键重构决策

| 决策 | 原因 | 当前实现 |
|------|------|----------|
| 数字参数统一由 `ARG_MODE` + `:RUN_ARG_FUNCTION` 分发 | 原 `ARG_RUN_1`～`ARG_RUN_7` 逻辑重复 | 单一标签 + 7 行条件跳转 |
| 每个功能自行显示成功消息 | 通用 `SHOW_SUCCESS` 会显示错误文案 | 各功能末尾独立输出 `!Fn_SUCCESS!` |
| 路径输入使用 `PATH_INPUT` 变量 | 原硬编码英文提示无法实现国际化 | 中英文双语 |
| 目录扫描统一引号包裹 | 支持含空格路径 | 所有 `for /d %%d in ("!REPO_ROOT!\*")` 均带引号 |
| 应用代码提升到仓库根 | `prototype/` 只保留设计规范 | 根 `src/`，见 `CONTRIBUTING.md` §3 |

### 2.3 参考界面应用模块划分

| 路径 | 职责 |
|------|------|
| `src/app/` | 路由与全局样式（`layout.tsx` / `page.tsx` / `globals.css`） |
| `src/components/icons.tsx` | 内联 SVG 图标，零外部图标依赖 |
| `src/components/ui/` | 无状态基础组件（alert / badge / button / card / code-block / table / tabs） |
| `src/lib/i18n.ts` | 唯一 UI 文案字典（zh / en） |
| `src/lib/data/` | 单一数据源（features / project-types / steps / repos），经 `@/lib/data` 导出 |

> `src/lib/data/repos.ts` 中的扫描结果与容量数字是**演示数据（mock）**，用于展示界面结构，非真实磁盘扫描结果，不得作为事实引用。

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

> 菜单上显示的短文案以 `lang.ini` 的 `MENU_OPTIONn` / `Fn_TITLE` 为准；文档只使用上表名称。

所有功能的作用域统一为：**遍历 `REPO_ROOT` 的直接子目录**（`for /d %%d in ("!REPO_ROOT!\*")`），不递归进入更深层级。

### 3.1 功能 1：启用全局 node_modules

- **描述**：配置系统环境变量 `NODE_PATH`，禁用项目本地 `node_modules` 生成。
- **操作步骤**：
  1. 显示警告并要求二次确认（输入 `Y` 继续，否则返回菜单，**不执行任何删除**）。
  2. `npm root -g` 获取全局 npm 模块路径。
  3. `setx NODE_PATH` 写入用户环境变量，并同步当前会话的 `NODE_PATH`。
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
- **已知限制**：`tmp=` 行的值**未加引号**（见 §7.3）。若 `%TEMP%` 含空格，npm 解析该配置可能异常。
- **预期结果**：项目共享全局依赖；此后仍需先运行功能 5 安装依赖，否则项目会因找不到模块而报错。

### 3.2 功能 2：重置 Node 配置

- **操作步骤**：`setx NODE_PATH ""` + 删除 `HKCU\Environment` 下的 `NODE_PATH` 注册表值；删除所有项目根目录的 `.npmrc`。
- **预期结果**：npm 恢复默认行为。

### 3.3 功能 3：清理并禁用 .next 缓存

- **操作步骤**：
  1. 遍历 `REPO_ROOT` 直接子目录，识别含 `next.config.js` **或** `next.config.ts` 的 Next.js 项目。
  2. `rd /s /q` 删除现有的 `.next` 缓存目录。
  3. 写入一个同名的 `.next` **文件**，阻止 Next.js 再创建 `.next` 目录。
  4. `attrib +h +s +r` 设置隐藏(H) + 系统(S) + 只读(R) 属性。
- **预期结果**：节省磁盘空间，但会影响 Next.js 热更新速度。

### 3.4 功能 4：恢复 Next.js 缓存

- **操作步骤**：
  1. 与功能 3 相同的项目识别规则。
  2. `attrib -h -s -r` **先清除** `.next` 文件的隐藏/系统/只读属性。
  3. `del /f /q` 删除该 `.next` 文件。
- **实现要点**：`del` 的多个 `/a:` 参数只有最后一个生效，且默认不删除隐藏/系统属性文件；因此必须先 `attrib` 清除属性，否则恢复会静默失败（v1.2.3 修复）。
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
- **注意**：此功能需要网络；单组失败不会中断后续组（脚本未检查 `%ERRORLEVEL%`）。

### 3.6 功能 6：清理所有项目缓存

- **前置**：显示警告并 `pause` 等待确认。
- **清理项**（每个 `REPO_ROOT` 直接子目录）：
  - 目录：`node_modules`、`.next`、`.nuxt`、`dist`、`build`
  - 文件：`package-lock.json`、`yarn.lock`、`pnpm-lock.yaml`
- **预期结果**：释放磁盘空间。删除不可逆，恢复需重新 `npm install`。

### 3.7 功能 7：清理 Vite/构建缓存

- **前置**：显示警告并 `pause` 等待确认。
- **清理项**：`.vite`、`.turbo`、`.nuxt`、`dist`、`build`、`coverage`、`.nyc_output`
- **预期结果**：清理构建缓存，适合 Vite / Turborepo 项目。

> 功能 6 与功能 7 有重叠项（`dist`、`build`、`.nuxt`），属预期设计：功能 6 面向「全量清理」，功能 7 面向「仅构建缓存」。

---

## 4. 用户交互规格

### 4.1 语言切换

| 项 | 内容 |
|----|------|
| 支持语言 | 中文（`zh`）、英文（`en`） |
| 切换方式 | 菜单 `[L]`；命令行 `-lang en\|zh`；`config.ini` 的 `LANG=` |
| 默认语言 | `zh` |
| 文案来源 | 脚本内 `:LOAD_LANGUAGE` 定义的变量；`lang.ini` 的 `[en]/[zh]` 节为文档参考，脚本不从中读取文案 |

> `-lang` 只写 `config.ini`，**禁止整文件覆盖 `lang.ini`**（会清空文案节）。

### 4.2 目录选择

优先级：命令行路径参数 → `config.ini` 的 `REPO_ROOT` → 交互式选择。

交互菜单 4 个选项：使用当前目录 (`%CD%`)、使用上级目录 (`%~dp0..\`)、手动输入完整路径、使用默认目录 (`E:\Github`)。输入路径不存在时报错并重新输入；选择保存时以「读旧 → 过滤 → 临时文件 → `move /y`」方式更新 `config.ini`。

### 4.3 命令行参数

| 参数 | 说明 | 示例 |
|------|------|------|
| `-lang en\|zh` | 设置语言并写入 `config.ini` | `RepoCleaner.bat -lang en` |
| `-h` / `--help` | 显示帮助信息 | `RepoCleaner.bat -h` |
| `1`-`7` | 直接运行对应功能 | `RepoCleaner.bat 1` |
| 路径 | 指定仓库根目录（支持含空格路径，需引号包裹） | `RepoCleaner.bat "D:\My Projects"` |
| 数字 + 路径 | 为指定目录直接运行功能 | `RepoCleaner.bat 5 E:\Projects` |

### 4.4 主菜单（中文默认）

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

- 路径不存在：显示 `PATH_ERROR` 并 `pause` 后返回目录选择。
- 危险操作：功能 1 需输入 `Y` 二次确认；功能 6 / 7 显示警告并 `pause`。
- 每个功能完成后显示自己的成功消息与 `pause`。

---

## 5. 技术规格

### 5.1 环境要求

| 项 | 要求 |
|----|------|
| 操作系统 | Windows（批处理脚本）；参考界面应用可跨平台开发 |
| 运行环境 | Windows 命令提示符 `cmd.exe` |
| 依赖软件 | Node.js ≥ 18 与 npm（功能 1 / 5 需要） |
| 参考界面应用 | Node.js ≥ 18、npm、Next.js 16（见 `package.json`） |

### 5.2 文件结构

```
RepoCleaner/
├── .github/PULL_REQUEST_TEMPLATE.md   # PR 审查清单模板
├── docs/                              # 所有规范与规格文档归档
│   ├── README.md                      # 文档中心（导航 + 职责边界）
│   ├── OPENSPEC.md                    # 本文：产品功能规格
│   ├── CONTRIBUTING.md                # 开发规范
│   ├── DESIGN-SYSTEM.md               # UI 设计规范
│   └── tokens.json                    # 设计令牌
├── public/                            # 参考界面应用静态资源
├── src/                               # 参考界面应用代码
│   ├── app/                           # layout.tsx / page.tsx / globals.css / favicon.ico
│   ├── components/                    # icons.tsx + ui/
│   └── lib/                           # i18n.ts / utils.ts / data/
├── prototype/README.md                # 原型目录说明（无代码、无规范正文）
├── RepoCleaner.bat                    # 产品本体
├── config.ini                         # REPO_ROOT / LANG
├── lang.ini                           # [en] / [zh] 文案参考节
├── README.MD / CHANGELOG.MD / LICENSE # 入口、变更历史、许可证
├── RepoCleaner.code-workspace
├── package.json / package-lock.json / tsconfig.json / next.config.ts
├── postcss.config.mjs / eslint.config.mjs / next-env.d.ts
└── .gitignore
```

### 5.3 配置文件格式

**`config.ini`**（脚本实际读取）：

```ini
REPO_ROOT=E:\Github
LANG=zh
```

**`lang.ini`**：首部 `LANG=` 记录当前语言选择；`[en]` / `[zh]` 两节保存中英文文案参考。脚本**不从该节读取文案**（文案在 `:LOAD_LANGUAGE` 中定义），此文件供文档与后续外部化使用。`-lang` 参数不得覆盖本文件。

### 5.4 编码要求

- 所有文本文件：**UTF-8**。
- 控制台：`chcp 65001` 切换到 UTF-8 代码页（脚本第 3 行）。
- 缩进：批处理 4 空格；TS/JS/CSS/JSON 2 空格。详见 `CONTRIBUTING.md` §4。

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
| `NODE_PATH` | 全局模块路径 | 功能 1 通过 `setx` 写入用户环境；功能 2 删除 |
| `TEMP` | 临时目录 | 系统环境变量，写入 `.npmrc` 的 `tmp=` |
| `REPO_ROOT` / `LANG` | 脚本内部变量 | 命令行 / `config.ini` / 交互输入 |

### 6.2 生成 / 修改的文件

| 文件 | 位置 | 用途 |
|------|------|------|
| `.npmrc` | 各项目根目录 | 配置 npm 使用全局模块（功能 1 写、功能 2 删） |
| `.next` | 各 Next.js 项目根目录 | 功能 3 创建为带 `+h+s+r` 的**文件**以阻止缓存；功能 4 清除属性后删除 |
| `config.ini` | 脚本同目录 | 保存 `REPO_ROOT` / `LANG` |

---

## 7. 兼容性规格

### 7.1 路径处理

- 支持带空格的路径（所有目录枚举与文件操作均带引号）。
- 支持相对路径与绝对路径。
- `:ROOT_SET` 自动去除 `REPO_ROOT` 末尾的反斜杠。

### 7.2 文件属性

- 功能 3：`attrib +h +s +r` 设置隐藏 + 系统 + 只读。
- 功能 4：**必须**先 `attrib -h -s -r` 清除全部属性，再 `del /f /q` 删除。

### 7.3 已知限制

| 编号 | 限制 | 影响 |
|------|------|------|
| L1 | `.npmrc` 的 `tmp=` 值未加引号 | `%TEMP%` 含空格时 npm 可能解析异常 |
| L2 | 功能 1 设置 `NODE_PATH` 后需重开终端才对新进程完全生效 | 当前会话已同步设置，新终端需重开 |
| L3 | 目录扫描仅覆盖 `REPO_ROOT` 的直接子目录，不递归 | monorepo 深层项目不会被处理 |
| L4 | 功能 5 不检查 `npm install -g` 的退出码 | 单组安装失败仍显示成功消息 |
| L5 | 功能 6 / 7 使用 `pause` 而非 `Y/N` 二次确认 | 任意键即继续，需使用者自行确认 |

> 记录限制是为了让规格与实现一致；修复限制时须同步更新本表与 `CHANGELOG.MD`。

---

## 8. 变更记录

**唯一来源：[CHANGELOG.MD](../CHANGELOG.MD)**（按版本倒序，含 `1.2.5` 及更早版本的完整变更明细）。

本文不再重复维护版本历史，避免两份记录漂移。

---

## 9. 验证与质量门禁

### 9.1 参考界面应用（自动，可在任意平台执行）

在仓库根目录依次执行，三项全部通过方可合并：

```bash
npm run type-check     # tsc --noEmit
npm run lint           # eslint
npm run build          # next build（含类型检查与静态预渲染）
```

预期：`type-check` / `lint` 无输出错误，`build` 输出 `✓ Compiled successfully` 并完成静态页面生成。

### 9.2 批处理脚本（手动，仅 Windows）

`RepoCleaner.bat` **没有自动化测试**，改动后必须在隔离的临时目录做冒烟验证。完整清单见 [`CONTRIBUTING.md` §9.3](./CONTRIBUTING.md)（帮助信息、`-lang` 写入、语言切换、功能 1 的 Y/非 Y 分支、功能 3 → 4 的 `.next` 文件生成与删除、含空格路径、路径不存在重试）。

> 严禁在真实仓库根目录执行删除类验证。

### 9.3 覆盖范围说明

- 当前仓库**未配置单元测试框架**，也没有 `npm test` 脚本；上述门禁是现有的全部自动检查。
- 历史版本文档中的「功能 1～7 ✅ 通过」是当时的手工验证结论，不作为当前自动化覆盖证据。

---

## 10. 扩展计划

- [ ] 支持更多框架缓存清理（Angular、Svelte）
- [ ] 添加 dry-run 模式预览待删除文件
- [ ] 支持自定义清理规则
- [ ] 添加日志记录功能
- [ ] 功能 5 检查 `npm install -g` 退出码并汇总失败项（对应 L4）
- [ ] 功能 6 / 7 改为 `Y/N` 二次确认（对应 L5）
- [ ] 支持递归扫描 monorepo 深层项目（对应 L3）
- [ ] 添加详细进度显示

---

## 11. 许可证

MIT License，详见 [LICENSE](../LICENSE)。
