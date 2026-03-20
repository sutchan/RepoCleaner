# 🧰 项目清洁助手 (RepoCleaner)

<div align="center">

[![Windows](https://img.shields.io/badge/Windows-0078D6?style=flat&logo=windows&logoColor=white)](https://github.com/sutchan/RepoCleaner)
[![Batch](https://img.shields.io/badge/Batch-.bat-blue?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Version](https://img.shields.io/badge/Version-1.0.1-green?style=flat)](https://github.com/sutchan/RepoCleaner)
[![License](https://img.shields.io/badge/License-MIT-orange?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Node](https://img.shields.io/badge/Node.js-339933?style=flat&logo=node.js&logoColor=white)](https://nodejs.org)
[![Disk](https://img.shields.io/badge/Saves-10GB%2B-red?style=flat)](https://github.com/sutchan/RepoCleaner)
[![Auto](https://img.shields.io/badge/Auto-Clean-4CAF50?style=flat)](https://github.com/sutchan/RepoCleaner)

</div>

> 一键让所有前端项目保持干净：无冗余 node_modules、无 .next 缓存、无打包垃圾

适配项目类型：NextJS / React / Vue / 常规 Node 项目

---

## ✨ 功能介绍

本助手用于管理项目目录，自动解决以下痛点：

| 问题 | 解决方案 |
|------|----------|
| 每个项目重复生成巨大的 `node_modules` 占用空间 | 全局 node_modules，项目永不生成本地依赖 |
| NextJS 项目产生大量 `.next` 缓存文件 | 禁用/启用 .next 缓存一键切换 |
| 项目打包产生 `dist`/`build` 垃圾文件 | 一键清理所有缓存 |
| 依赖混乱、无法统一管理 | 统一安装到全局，告别重复安装 |

---

## 🚀 快速开始

### 前置要求

- Windows 系统
- 已安装 Node.js 和 npm

### 使用步骤

1. **将 `RepoCleaner.bat` 放在项目根目录**
2. **双击运行**
3. **选择仓库根目录**（支持自动检测、手动输入、保存配置）
4. **按数字菜单选择功能**
5. **一次配置，所有项目永久干净**

### 🌟 推荐使用流程

| 场景 | 操作 |
|------|------|
| 首次配置 | 运行 [1] 启用全局 node_modules → 运行 [5] 安装所有常用依赖 → 运行 [3] 禁用 .next 缓存 |
| 空间不足 | 运行 [6] 清理所有缓存 |
| Vite 项目 | 运行 [7] 清理 Vite/构建缓存 |

---

## 📖 功能详解

### 功能1：启用全局 node_modules

- 自动获取全局 npm 模块路径
- 配置系统环境变量 `NODE_PATH`
- 为所有包含 `package.json` 的项目生成 `.npmrc`，禁用本地依赖生成

> 首次使用后，项目将共享全局依赖，不再生成 `node_modules` 文件夹

### 功能2：还原默认 Node 配置

- 清除 `NODE_PATH` 环境变量
- 删除所有项目的 `.npmrc` 配置文件
- 恢复 npm 默认行为

### 功能3：清理 & 禁用 .next 缓存

- 扫描所有包含 `next.config.js` 的项目
- 删除现有的 `.next` 缓存目录
- 创建不可见的 `.next` 文件阻止缓存生成

> 适合磁盘空间紧张时使用，会影响 NextJS 热更新速度

### 功能4：还原 NextJS 缓存

- 删除禁止缓存的 `.next` 文件
- 恢复 NextJS 正常编译和热更新功能

### 功能5：一键安装常用全局依赖

一次性安装以下常用开发库：

```
react react-dom next
vue @vue/cli
axios express
pnpm yarn
cross-env dotenv nodemon pm2 rimraf mkdirp
```

### 功能6：一键清理所有项目缓存

清理以下内容：

| 清理项 | 说明 |
|--------|------|
| `node_modules` | 项目依赖目录 |
| `.next` | NextJS 缓存 |
| `.nuxt` | NuxtJS 缓存 |
| `dist` | 打包输出目录 |
| `build` | 构建输出目录 |
| `package-lock.json` | npm 锁文件 |
| `yarn.lock` | yarn 锁文件 |
| `pnpm-lock.yaml` | pnpm 锁文件 |

### 功能7：清理 Vite/构建缓存

清理以下内容：

| 清理项 | 说明 |
|--------|------|
| `.vite` | Vite 缓存目录 |
| `.turbo` | Turborepo 缓存 |
| `.nuxt` | NuxtJS 缓存 |
| `dist` | 打包输出目录 |
| `build` | 构建输出目录 |

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
| 交互选择 | 运行后按 [L] 键切换语言 |
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

---

## ⚠️ 注意事项

1. **首次使用建议先运行 [5] 安装全局依赖**
2. **功能3 会影响 NextJS 热更新速度**，开发完成后可恢复
3. **功能6 删除后需要重新 `npm install`**，建议在确认不再开发时使用
4. **默认扫描 `E:\Github\` 目录**，可通过交互菜单或配置文件修改

---

## 📄 License

MIT License

---

## 📝 更新日志

详细更新内容请查看 [CHANGELOG.md](./CHANGELOG.md)
