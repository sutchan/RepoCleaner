# 🧰 项目清洁助手 (RepoCleaner)

<div align="center">

[![Windows](https://img.shields.io/badge/Windows-0078D6?style=flat&logo=windows&logoColor=white)](https://github.com/sutchan/RepoCleaner)
[![Batch](https://img.shields.io/badge/Batch-.bat-blue?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Version](https://img.shields.io/badge/Version-1.2.4-green?style=flat)](https://github.com/sutchan/RepoCleaner)
[![License](https://img.shields.io/badge/License-MIT-orange?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Node](https://img.shields.io/badge/Node.js-339933?style=flat&logo=node.js&logoColor=white)](https://nodejs.org)
[![Disk](https://img.shields.io/badge/Saves-10GB%2B-red?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Auto](https://img.shields.io/badge/Auto-Clean-4CAF50?style=flat)](https://github.com/sutchan/RepoCleaner)
[![i18n](https://img.shields.io/badge/i18n-中文%2FEnglish-9C27B0?style=flat)](https://github.com/sutchan/RepoCleaner)

</div>

> 一键让所有前端项目保持干净：无冗余 `node_modules`、无 `.next` 缓存、无打包垃圾。

适配项目类型：**Next.js / React / Vue / Vite / 常规 Node 项目**

---

## 📑 目录

- [功能亮点](#-功能亮点)
- [快速开始](#-快速开始)
- [功能详解](#-功能详解)
- [智能配置](#-智能配置)
- [GitHub 仓库信息（推荐设置）](#-github-仓库信息推荐设置)
- [常见问题 FAQ](#-常见问题-faq)
- [注意事项](#-注意事项)
- [原型与设计规范](#-原型与设计规范-prototype)
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
| Vite / Turborepo 等构建缓存膨胀 | 专清 `.vite`/`.turbo`/`coverage` 等目录 |

**特性**

- 🪟 纯 Windows 批处理（`.bat`），零运行时依赖，双击即用
- 🌏 中文 / English 双语菜单，可交互或命令行切换
- ⚙️ 支持命令行参数、配置文件、交互式选择三种运行方式
- 💾 自动保存仓库根目录配置，一次设置永久生效
- 🛡️ 危险操作（删除 `node_modules`）前二次确认，避免误删

---

## 🚀 快速开始

### 前置要求

- Windows 系统
- 已安装 Node.js 和 npm（建议 Node 18+）

### 使用步骤

1. **将 `RepoCleaner.bat` 放在项目根目录**
2. **双击运行**
3. **选择仓库根目录**（支持自动检测、手动输入、保存配置）
4. **按数字菜单选择功能**
5. **一次配置，所有项目永久干净**

### 🌟 推荐使用流程

| 场景 | 操作 |
|------|------|
| 首次配置 | 运行 `[1]` 启用全局 node_modules → 运行 `[5]` 安装所有常用依赖 → 运行 `[3]` 禁用 `.next` 缓存 |
| 空间不足 | 运行 `[6]` 清理所有缓存 |
| Vite 项目 | 运行 `[7]` 清理 Vite/构建缓存 |

---

## 📖 功能详解

### 功能 1：启用全局 node_modules

- 自动获取全局 npm 模块路径（`npm root -g`）
- 配置系统环境变量 `NODE_PATH`
- 为所有包含 `package.json` 的项目生成 `.npmrc`，禁用本地依赖生成
- 清理现有的本地 `node_modules` 目录（执行前需输入 `Y` 确认）

> 首次使用后，项目将共享全局依赖，不再生成 `node_modules` 文件夹

### 功能 2：还原默认 Node 配置

- 清除 `NODE_PATH` 环境变量
- 删除所有项目的 `.npmrc` 配置文件
- 恢复 npm 默认行为

### 功能 3：清理 & 禁用 .next 缓存

- 扫描所有包含 `next.config.js` / `next.config.ts` 的项目
- 删除现有的 `.next` 缓存目录
- 创建不可见的 `.next` 文件（隐藏 + 系统 + 只读属性）阻止缓存生成

> 适合磁盘空间紧张时使用，会影响 Next.js 热更新速度

### 功能 4：还原 NextJS 缓存

- 删除禁止缓存的 `.next` 文件（先清除隐藏/系统/只读属性再删除）
- 恢复 Next.js 正常编译和热更新功能

### 功能 5：一键安装常用全局依赖

一次性安装以下常用开发库：

```
react react-dom next
vue @vue/cli
axios express
pnpm yarn
cross-env dotenv nodemon pm2 rimraf mkdirp
```

### 功能 6：一键清理所有项目缓存

清理以下内容：

| 清理项 | 说明 |
|--------|------|
| `node_modules` | 项目依赖目录 |
| `.next` | Next.js 缓存 |
| `.nuxt` | NuxtJS 缓存 |
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
| `.nuxt` | NuxtJS 缓存 |
| `dist` | 打包输出目录 |
| `build` | 构建输出目录 |
| `coverage` | 测试覆盖率报告 |
| `.nyc_output` | NYC 输出目录 |

---

## ⚙️ 智能配置

### 目录选择方式

运行脚本时支持多种目录选择方式：

| 方式 | 说明 |
|------|------|
| 命令行参数 | `RepoCleaner.bat "D:\Projects"` |
| 配置文件 | 同目录下的 `config.ini` |
| 交互选择 | 运行后手动选择目录 |

### 语言切换

支持中文/英文界面切换：

| 方式 | 说明 |
|------|------|
| 交互选择 | 运行后按 `[L]` 键切换语言 |
| 命令行参数 | `RepoCleaner.bat -lang en` |
| 配置文件 | 在 `config.ini` 中添加 `LANG=zh` 或 `LANG=en` |

### 配置文件示例

```ini
REPO_ROOT=E:\Github
LANG=zh
```

### 命令行参数

```bash
RepoCleaner.bat [选项] [功能号] [目录路径]
```

| 参数 | 说明 |
|------|------|
| `-lang en\|zh` | 设置语言 |
| `-h` 或 `--help` | 显示帮助信息 |
| `1-7` | 直接运行对应功能 |
| 目录路径 | 指定仓库根目录 |

**示例**

```bash
RepoCleaner.bat                    # 交互模式
RepoCleaner.bat E:\Github          # 指定仓库根目录
RepoCleaner.bat 1                  # 直接启用全局 node_modules
RepoCleaner.bat -lang en           # 使用英文
RepoCleaner.bat 5 E:\Projects      # 为指定路径安装全局依赖
```

---

## 🌐 GitHub 仓库信息（推荐设置）

为了让项目页面更丰富、更易被发现，建议在本仓库 **Settings → About** 中填写以下内容（本工具无法自动写入，请手动复制）：

### 仓库描述（Description）

```
🧰 一键清理前端项目冗余：统一管理全局 node_modules、禁用 .next/.vite 缓存、清理 dist/build 垃圾。支持 Next.js / React / Vue / Vite，中文/英文双语菜单，单仓库可节省数 GB 磁盘空间。
```

### 主页（Website）

可留空，或填写仓库地址：`https://github.com/sutchan/RepoCleaner`

### 话题标签（Topics）

复制以下标签到 **About → Topics**（每行一个，小写、连字符分隔，最多 20 个）：

```
windows
batch
batchfile
node-modules
cleanup-tool
disk-cleaner
nextjs
react
vuejs
vite
npm
cache-cleaner
developer-tools
cli
repository-maintenance
frontend
```

---

## ❓ 常见问题 FAQ

**Q：功能 1 会影响现有项目运行吗？**
A：启用全局 node_modules 后，项目不再生成本地 `node_modules`，统一从全局读取。请确保已先运行功能 5 安装所需依赖，否则项目会因找不到模块而报错。

**Q：功能 3 禁用 `.next` 缓存后热更新变慢？**
A：是的，禁用缓存会略微影响 Next.js 热更新速度。开发完成后可运行功能 4 恢复。

**Q：功能 6 / 7 删除后如何恢复？**
A：清理是不可逆的。删除 `node_modules` 后需重新 `npm install`；建议仅在确认不再开发或已备份时执行。

**Q：默认扫描哪个目录？**
A：默认扫描 `E:\Github\`，可通过交互菜单、命令行参数或 `config.ini` 的 `REPO_ROOT` 修改。

**Q：支持含空格的仓库路径吗？**
A：支持。所有目录扫描均已做引号包裹处理，可正确识别含空格的路径。

---

## ⚠️ 注意事项

1. **首次使用建议先运行 `[5]` 安装全局依赖**
2. **功能 3 会影响 Next.js 热更新速度**，开发完成后可恢复
3. **功能 6 删除后需要重新 `npm install`**，建议在确认不再开发时使用
4. **默认扫描 `E:\Github\` 目录**，可通过交互菜单或配置文件修改
5. **删除操作不可逆**，执行前请确认已备份重要文件

---

## 🎨 原型与设计规范 (prototype/)

应用代码已提升到仓库根（`src/` + 根级 `package.json` / `next.config.ts` / `tsconfig.json`），是与本工具对齐的 Next.js + React + TypeScript + Tailwind v4 参考界面（OPENSPEC 规范的可视化落地）。

- **根目录应用** 即原型实现：`src/app/page.tsx` 单页 dashboard，数据统一经 `@/lib/data` 导出，文案统一经 `src/lib/i18n.ts` 双语切换
- **设计系统** `prototype/DESIGN-SYSTEM.md` 与 Tokens `prototype/tokens.json`（令牌与 `src/app/globals.css` 的 CSS 变量一致）
- **语义化 id** 所有 UI 容器/区块均带 `id`（如 `repo-cleaner-app`、`main-tabs`、`repo-card-<name>`）

本地预览（在仓库根目录）：

```bash
npm install
npm run dev      # 开发预览
npm run build    # 生产构建（需先 npm install 拉全依赖）
npm run type-check
```

> `prototype/` 目录仅保留设计规范与令牌等参考文件（`DESIGN-SYSTEM.md` / `tokens.json` / `README.md`），不再包含可运行代码。

---

## 📂 目录结构

```
RepoCleaner/
├── RepoCleaner.bat      # 主程序脚本（Windows 批处理）
├── config.ini           # 配置文件（REPO_ROOT / LANG）
├── lang.ini             # 语言配置（LANG 值 + [en]/[zh] 文案节）
├── README.MD            # 项目说明文档
├── CHANGELOG.MD         # 更新日志
├── OPENSPEC.md          # 开放规格说明书
├── RepoCleaner.code-workspace  # VS Code 工作区配置
├── package.json         # 根应用依赖与脚本（dev/build/type-check）
├── next.config.ts       # Next.js 配置
├── tsconfig.json        # TypeScript 配置（@/* → ./src/*）
├── postcss.config.mjs   # Tailwind v4 PostCSS 配置
├── eslint.config.mjs    # ESLint 配置
├── public/              # 静态资源
├── src/                 # 参考界面应用代码（App Router）
│   ├── app/             # layout.tsx / page.tsx / globals.css
│   ├── components/      # icons.tsx + ui/ 基础组件
│   └── lib/             # i18n.ts + data/（单一数据源）
└── prototype/           # 仅保留设计规范与令牌
    ├── DESIGN-SYSTEM.md
    ├── tokens.json
    └── README.md
```

---

## 📄 许可证

[MIT License](./LICENSE)

---

## 📝 更新日志

详细更新内容请查看 [CHANGELOG.md](./CHANGELOG.md)
