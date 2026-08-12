# Changelog

All notable changes to this project will be documented in this file.

---

## [1.2.1] - 2026-08-12

### Fixed
- **config.ini 去重** - 删除重复的 `LANG` 行与尾部多余空行/空格，避免 bat 读取 `if exist` 路径误判
- **prototype 页面重构** - 重写 `page.tsx`，接入真实 7 项功能数据与 `src/lib/data.ts` 单一数据源，移除 Next.js 默认模板占位内容
- **prototype 图标零依赖** - 用内联 SVG 图标组件替换外部 `lucide-react` 依赖，消除未安装依赖导致的构建失败，符合"无外部依赖"理念
- **prototype 版本号** - 修正 package.json/DESIGN-SYSTEM 中散落的 v0.1.0，统一为 v1.2.1

### Added
- **prototype i18n 双语切换** - 新建 `src/lib/i18n.ts`，顶部语言开关（中/EN）实时切换全部文案，翻译键全覆盖，对齐 bat 的 i18n 行为
- **语义化容器 id** - prototype 所有 UI 容器/区块加 `id`（如 `repo-cleaner-app`、`overview`、`main-tabs`、`repo-card-<name>`、`security-note`），便于调试定位与无障碍访问
- **根 `.gitignore`** - 忽略 npm 误存缓存（`%USERPROFILE%/`、`.npm-cache/`）、日志与构建产物
- **prototype type-check 脚本** - package.json 新增 `type-check`（tsc --noEmit）

### Changed
- **版本号统一** - README 徽章、RepoCleaner.bat 头、lang.ini 头、prototype（package.json/layout/data 等头注释）全部同步至 v1.2.1

### Security
- bat 删除操作路径均带引号且来自 `for` 目录枚举变量，未引入任意路径执行风险
- `%USERPROFILE%/` 误存缓存已通过 `.gitignore` 忽略，避免将 npm 缓存误提交

---

## [1.1.0] - 2026-08-12

### Fixed
- **消除命令行参数解析重复代码** - 将 7 个 `ARG_RUN_1~7` 统一为单标签 `RUN_ARG_FUNCTION` 分发，对齐 OPENSPEC 2.2
- **i18n 缺口补全** - `lang.ini` 英文节补齐 `MENU_OPTION7`，修正 `MENU_CHOOSE` 范围 `[0-6]`→`[0-7]`
- **交互式路径输入国际化** - 硬编码 `Enter full path: ` 改为 `PATH_INPUT` 变量，支持中英文切换（OPENSPEC 2.3）
- **config.ini 修复** - 去除 `REPO_ROOT=E:\Github` 值末尾多余空格，避免 `if exist` 误判路径不存在
- **版本号统一** - README 徽章、lang.ini 头、CHANGELOG 同步至 v1.1.0

### Security
- 路径校验沿用 `if exist` 白名单，未引入任意路径执行风险

---

## [1.0.1] - 2025-03-19

### Added
- **多语言支持** - 支持中文/英文切换，可通过 `lang.ini` 配置
- **功能7: 清理 Vite/构建缓存** - 新增清理 `.vite`、`dist`、`.turbo` 等构建缓存
- **智能目录选择** - 支持命令行参数、配置文件、交互式选择三种方式
- **首次运行语言选择** - 启动时显示语言选择菜单
- **自动保存配置** - 可将目录配置保存到 `config.ini`

### Improved
- **功能1增强** - 启用全局 node_modules 时自动清理现有本地 `node_modules` 目录
- **.npmrc 配置优化** - 添加 cache 和 tmp 配置项
- **用户交互** - 改进提示信息和错误处理

---

## [1.0.0] - 2025-03-18

### Added
- **功能1: 启用全局 node_modules** - 项目永不生成本地依赖文件夹
- **功能2: 还原默认 Node 配置** - 恢复系统原始状态
- **功能3: 清理 & 禁用 .next 缓存** - 禁止 NextJS 生成缓存目录
- **功能4: 还原 NextJS 缓存** - 恢复正常编译
- **功能5: 一键安装常用全局依赖** - React/Vue/Next/axios 等开发库
- **功能6: 一键清理所有项目缓存** - 删除 node_modules/.next/dist/build
