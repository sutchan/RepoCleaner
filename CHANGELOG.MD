# Changelog

All notable changes to this project will be documented in this file.

---

## [1.2.5] - 2026-11-04

### Added
- **建立 `docs/` 文档体系** - 将 `OPENSPEC.md`、`CONTRIBUTING.md` 从根，`DESIGN-SYSTEM.md`、`tokens.json` 从 `prototype/` 统一迁入 `docs/`；新增 `docs/README.md` 文档中心（文档地图、职责边界、按角色阅读路径、仓库结构速览、维护约定、术语速查）
- **完整开发规范 `docs/CONTRIBUTING.md`** - 明确编码风格、命名约定、目录组织、语义化 id、提交信息格式（type 九类）、分支管理、代码审查流程与清单、版本号规则、文档维护约定、术语表、安全红线
- **PR 审查模板** - 新增 `.github/PULL_REQUEST_TEMPLATE.md`，把审查清单与质量门禁落地为可勾选模板
- **文档链接校验门禁** - 新增 `scripts/check-doc-links.mjs` 与 `npm run check:docs`，逐段比对真实文件名校验所有 Markdown 相对链接（含大小写），失效返回退出码 1；已纳入提交前检查与 PR 审查清单
- **设计规范差距清单** - `docs/DESIGN-SYSTEM.md` §7 明确列出 G1～G8 规划未实现项，避免设计愿景被当作既成事实

### Changed
- **规格与实现对齐** - 重写 `docs/OPENSPEC.md`：删除易失真的源码行号表与重复版本历史（统一指向 `CHANGELOG.MD`）、补「已知限制」L1～L5、把虚构的「功能全部 ✅ 通过」表改为可执行的门禁命令与手动冒烟清单
- **设计规范与实现对齐** - 修正图标描边（1.5px → 实际 `strokeWidth={2}`）、移除不存在的第三方 Spinner 描述、标注亮色主题/间距/圆角/动效/响应式断点/`a11y` 的实际落地状态
- **统一术语与功能名称** - 全仓库统一 `Next.js` / `Nuxt`、功能 1～7 名称（对齐 `docs/OPENSPEC.md` §3.0）、「仓库根目录 vs 项目根目录」「参考界面应用」「规范文档目录」等表述；同步替换 `RepoCleaner.bat` 中英菜单、`lang.ini` 参考节、`src/lib/data/features.ts` 与 `project-types.ts` 的旧写法（`NextJS`、还原默认/还原、一键安装等），并修正菜单对齐空格；`CHANGELOG.MD` 历史条目按规范保留原貌不改
- **消除冗余** - `prototype/README.md` 精简为原型目录说明并指向 `docs/`；删除 Next.js 模板遗留的 `prototype/.gitignore`（规则已被根 `.gitignore` 覆盖）与误生成的 `prototype/next-env.d.ts`，并在根 `.gitignore` 忽略 `next-env.d.ts` 防止再次误提交
- **版本号单一来源** - `next.config.ts` 从 `package.json` 读取版本并注入 `NEXT_PUBLIC_APP_VERSION` 与 `NEXT_PUBLIC_BUILD_ID`（`yyyymmddhhmm`），`src/app/page.tsx` 页脚展示版本号 + build 编号，注入缺失时回退仅显示版本号
- **演示数据标注** - `src/lib/data/repos.ts` 头注释明确以下为 mock 数据、容量数字为示意，不可作为事实引用
- **版本号同步** - `package.json`、`package-lock.json`（历史停在 1.2.3，一并修正）、README 徽章与页脚、`docs/tokens.json`、各规范文档标题同步至 `1.2.5`

### Fixed
- **README 链接大小写 bug** - README 曾把链接写成小写扩展名的 `CHANGELOG.md`，实际文件为 `CHANGELOG.MD`，在 Linux / GitHub 上会 404；按实际大小写修正
- **产品事实矛盾** - `src/lib/i18n.ts` 页脚文案「跨平台仓库清理工具」与纯 Windows 批处理产品不符，改为「RepoCleaner 项目清洁助手」/「RepoCleaner Project Cleaner」
- **文档路径失效** - 迁移后全量修正各文档的相对链接与目录树，新增链接大小写/层级校验要求
- **繁简混用** - 修正 `OPENSPEC.md` 中的「現有的」等用字

---

## [1.2.4] - 2026-11-04

### Docs
- **丰富项目页面信息** - 重写 README，新增目录、功能亮点特性、常见问题 FAQ、GitHub 仓库信息（推荐描述 / 主页 / 话题标签）、目录结构等章节，使项目页面更完整
- **补充 LICENSE** - 仓库声明 MIT 许可但缺失 LICENSE 文件，新增标准 MIT 许可证文本（含 `./LICENSE` 链接）

### Changed
- **版本号同步** - package.json 与 README 徽章统一至 `1.2.4`

---

## [1.2.3] - 2026-11-04

### Changed
- **应用提升至仓库根** - 将 Next.js 参考界面从 `prototype/` 提升到根目录（`src/` + 根级 `package.json` / `next.config.ts` / `tsconfig.json` / `postcss.config.mjs` / `eslint.config.mjs` / `public/`），`prototype/` 仅保留设计规范与令牌（`DESIGN-SYSTEM.md` / `tokens.json` / `README.md`）
- **package.json 更名** - `name` 由 `prototype` 改为 `repo-cleaner`，版本同步至 `1.2.3`
- **文档对齐新结构** - README「原型」章节、OPENSPEC §4.2 文件结构同步根级应用布局

### Fixed
- **prototype 客户端边界（P0）** - `src/app/page.tsx` 补充 `"use client"`，否则 App Router 下 `useState` 会在 `next build` 阶段报错，应用无法运行
- **构建期联网依赖** - `layout.tsx` 移除 `next/font/google` 的 `Inter`（离线/沙箱环境拉取字体会导致构建失败），改用 `font-sans` 系统字体令牌，对齐 DESIGN-SYSTEM 字体规范
- **违反设计令牌规范** - 页面与组件中的硬编码颜色（`bg-[#0d1117]` / `text-gray-100/400/500` / `bg-[#26262a]` / `bg-[#0d0d0f]`）全部替换为 `bg-background` / `text-foreground` / `text-muted-foreground` / `bg-muted` 等令牌工具类，消除散落硬编码
- **功能4 恢复 NextJS 缓存失效（P1）** - `RepoCleaner.bat` 删除被设为 隐藏+系统+只读 的 `.next` 文件时，`del` 多 `/a:` 过滤仅末项生效且默认不删隐藏/系统文件，导致恢复静默失败；改为先 `attrib -h -s -r` 清除属性再 `del /f /q`
- **版本号统一** - README 徽章、package.json、lang.ini 头、bat 头、各源文件头注释、`tokens.json`、`DESIGN-SYSTEM.md`、`page.tsx` 页脚同步至 `1.2.3`

### Security
- bat 删除操作路径仍仅来自 `for` 目录枚举变量并带引号，未引入任意路径执行风险

---

## [1.2.2] - 2026-08-13

### Fixed
- **prototype 编译失败（P0）** - `page.tsx` 引用不存在的 `scanResults`/`tweaks`、`AlertTitle`/`AlertDescription`、`variant="default"`、`Tabs defaultValue` 等导致 20 个 TS 错误；新增 `src/lib/data/repos.ts`（扫描结果/优化建议）并补全 `alert/button/tabs` 组件 API，`tsc --noEmit` 恢复通过
- **数据层拆分** - `src/lib/data.ts`（236 行）按职责拆为 `src/lib/data/`（features/project-types/steps/repos + barrel），消除与 `src/lib/i18n.ts` 的双 i18n 冲突，保持 `@/lib/data` 导出契约不变
- **`page.tsx` 超限拆分** - 212 行 → 158 行，内联 SVG 图标抽离至 `src/components/icons.tsx`
- **空格路径支持** - 全部 7 处 `for /d` 目录扫描加引号，真正支持含空格的仓库根目录
- **`next.config.ts` 检测** - 功能 3/4 同时识别 `next.config.js` 与 `next.config.ts`（Next.js 15+/16 默认）
- **功能 1 二次确认** - 执行删除前要求输入 Y 确认，避免误删本地 node_modules
- **config.ini 写入修复** - `-lang` 参数与路径保存不再产生重复 `LANG=` 行、不再丢失另一配置项（实测验证）
- **lang.ini 保护** - `-lang` 不再整文件覆盖 lang.ini，避免清空 `[en]/[zh]` 文案节
- **globals.css 版本对齐** - 头注释 v1.2.0 → v1.2.2

### Changed
- **产品名统一** - bat 菜单/`lang.ini`/prototype `i18n.ts` 统一为「RepoCleaner 项目清洁助手」
- **文档对齐实现** - `DESIGN-SYSTEM.md` 更新实际组件/图标/交互；README 修正资源表述（`data/` 目录与 `tokens.json`）；OPENSPEC 更新功能规格与文件结构；重写 prototype/README.md（移除 Next.js 默认模板）
- **新增 `tokens.json`** - 设计令牌与 `DESIGN-SYSTEM.md` §2 及 `globals.css` CSS 变量一致
- **仓库卫生** - 删除 prototype 残留的 `%USERPROFILE%/` npm 缓存、日志与 `tsconfig.tsbuildinfo`；根 `.gitignore` 忽略 `.codebuddy/`

### Security
- 路径校验沿用 `for` 目录枚举 + `if exist` 白名单，本次未引入任意路径执行风险

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
