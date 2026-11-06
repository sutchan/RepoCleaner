# RepoCleaner 已完成任务 (docs/tasks-done.md)

> 最后更新：2026-11-06
> 本文件归档已完成的任务（来源：用户初始 8 项综合审查任务 + 后续 v1.2.3–v1.2.6 演进 + 本轮剩余任务）。

## 一、初始综合审查任务（用户首条指令）

| ID | 任务 | 完成版本 | 证据 / 结果 |
|----|------|----------|-------------|
| C1 | 审查并修复代码 bug / 语法 / 安全 / 性能 | v1.2.3 | 修复 page.tsx 缺 `"use client"`（否则 `next build` 报错）；layout.tsx 移除 `next/font/google` 联网依赖；消除散落硬编码颜色；bat 功能4 恢复缓存失效修复 |
| C2 | 代码对齐规范文档与原型 | v1.2.3 | 颜色改用 design token 工具类，字体改系统字体栈，语义化 id 对齐 DESIGN-SYSTEM §3.3 |
| C3 | 原型与项目代码加语义化 id | v1.2.3 | page.tsx 为 header/overview/security-note/main-tabs/repo-card-*/tweak-card-*/app-footer 等配置 id |
| C4 | 拆分 >200 行源代码文件 | — | 经核查 `src/**/*.{ts,tsx}` 均 ≤200 行，无需拆分；`.bat` 不在源码 200 行拆分规则内 |
| C5 | 测试功能完整性与鲁棒性 | v1.2.3 | `npm install` → `type-check` → `next build`（编译+预渲染）→ `eslint` 全部通过 |
| C6 | 更新项目文档（功能/使用/变更/记忆） | v1.2.3–v1.2.6 | README 原型章节、OPENSPEC §4.2、DESIGN-SYSTEM、CHANGELOG 同步；新增项目记忆；规范统一归档至 `docs/`（v1.2.5/1.2.6） |
| C7 | 修改过的文件版本号 bump 至最新 | v1.2.3 | 各被改文件头注释与全局展示位同步 bump |
| C8 | 提交并推送所有修改 | v1.2.3 | 已推送 `origin/main`（含忽略 `tsconfig.tsbuildinfo` 清理提交） |

## 二、后续演进（v1.2.4–v1.2.6，由规范归档与文档完善提交完成）

| ID | 任务 | 版本 | 结果 |
|----|------|------|------|
| C9 | 完善 README 项目信息、补充 LICENSE | v1.2.4 | README 重写 + MIT LICENSE |
| C10 | 规范文档统一归档至 `docs/` | v1.2.5/1.2.6 | DESIGN-SYSTEM/tokens/OPENSPEC/CONTRIBUTING 迁入 `docs/`；`prototype/` 仅留说明；DESIGN-SYSTEM §7 差距清单维护 |
| C11 | 版本展示位与构建注入对齐 | v1.2.6 | `next.config.ts` 从 `package.json` 注入 `NEXT_PUBLIC_APP_VERSION` 与 `BUILD_ID`，页脚展示版本+构建号；`package.json`/`tokens.json`/README 徽章同步 1.2.6 |

## 三、本轮剩余任务完成（2026-11-05）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| R1 | 同步版本展示位至 1.2.6（lang.ini、RepoCleaner.bat 头注释） | 中 | 已完成 | 编辑 `lang.ini`（`; Version: 1.2.6`）与 `RepoCleaner.bat`（`v1.2.6`）；`page.tsx` 页脚由 `next.config.ts` 注入，无需改。已提交推送 |
| R2 | 复核 RepoCleaner.bat 健壮性与安全性 | 低 | 已完成（核验） | 删除/枚举路径均引号包裹、来源为 `for /d` 枚举非用户输入，无命令注入；F4 此前已修复；无关键缺陷 |
| R3 | 全量核验 src 应用语义化 ID 覆盖交互元素 | 低 | 已完成（核验） | page.tsx 已为 header/overview/security-note/main-tabs/repo-card-*/tweak-card-*/app-footer 等配置 id，符合 DESIGN-SYSTEM §3.3 |
| R4 | 确认 DESIGN-SYSTEM §7 差距清单与代码一致 | 低 | 已完成（核验） | §7 G1–G8 含「规划/已实现」标记，亮色主题、Tabs ARIA、清理按钮 onClick 等明确列为未实现，文档与实现一致 |
| R5 | 页脚版本与 10 位 build 编号规范落地 | 高 | 已完成 | v2.0.1：`internal/buildinfo.Footer()` 统一生成 `vX.Y.Z · build yymmddhhmm`，构建脚本 `-ldflags -X` 注入，Web 侧 `next.config.ts` 改 10 位；规范写入 `CONTRIBUTING.md` §8.2 / `DESIGN-SYSTEM.md` §3.0 / `OPENSPEC.md` §4.6 |
