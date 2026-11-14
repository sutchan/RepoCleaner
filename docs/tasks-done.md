# Repo Manager 已完成任务 (docs/tasks-done.md)

> 最后更新：**2026-11-14 17:58**
> 本文件归档已完成的任务（来源：用户初始 8 项综合审查任务 + 后续 v1.2.3–v1.2.6 演进 + 本轮剩余任务 + v2.0.0 Go 迁移清单 + v2.3.1 文档体系重新规划 + v2.3.2 任务核查）。
> 未完成事项只在 [tasks.md](./tasks.md) 维护。

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
| R6 | dry-run 预演模式 + L4 修复 | 高 | 已完成 | v2.1.0：`engine.Options.DryRun` 贯穿功能 1–7，CLI `-dry-run` / GUI 复选框（会话级），预演跳过确认；功能 5 逐组检查退出码并双语汇总失败组、返回错误；`dryrun_test.go` 覆盖，nogui 门禁全绿，GUI 编译需 C 工具链留待 CI |
| V2 | 支持更多框架缓存清理（Angular / Svelte） | 低 | 已完成 | v2.2.0：功能 7 白名单新增 `.angular`、`.svelte-kit`，`WarningVite` 双语文案与 `OPENSPEC.md` §2 / §3.7 同步；新增 `TestCleanAngularSvelteCache`；顺带修复 `config.ini` 跟踪问题（.gitignore + `docs/config.example.ini`） |
| V6 | 落地 GitHub Actions：门禁工作流 + 全绿后自动发布 Release | 中 | 已完成 | v2.2.0：新增 `.github/workflows/ci.yml`（nogui 全套门禁 + Linux 默认标签 GUI 编译）与 `release.yml`（tag `vX.Y.Z` 触发，gate 重跑门禁 → 三平台矩阵构建 → tag 与版本权威源一致性校验 → SHA256SUMS + CHANGELOG 提取说明 → 自动 Release）；规范落地于 `CONTRIBUTING.md` §8.4；windows/arm64 因 runner 缺交叉工具链暂不产出（已在 §8.4 标注） |

## 四、v2.0.0 Go 迁移清单 G01–G21（2026-11-05 ~ 2026-11-06 全部完成）

| ID | 任务 | 完成版本 | 证据 / 结果 |
|----|------|----------|-------------|
| G01 | 初始化 Go module | 2.0.0 | `go.mod`（module `repo-cleaner`，go 1.25，Fyne v2）+ `go.sum` |
| G02 | `internal/buildinfo/version.go` | 2.0.1 | `Version` 常量 + `Footer()` / `BuildInfo()` + 构建期 `BuildID` 注入与回退；`version_test.go` 覆盖格式 |
| G03 | `internal/platform` 接口定义 | 2.0.0 | `Platform` 12 个方法 + `HasFile` / `IsNodeProject` / `IsNextProject` / `NormalizeRoot` / `ValidateRoot` |
| G04 | Windows 平台实现 | 2.0.2 | `platform_windows.go`：`attrib` / `setx` + `reg` / `E:\Github` / `npm root -g`；v2.0.2 消除吞错 |
| G05 | Unix 平台实现 | 2.0.2 | `platform_unix.go`：`~/.profile` 持久化 / `~/Github`；v2.0.2 补齐 macOS `chflags hidden/nohidden` |
| G06 | `internal/config` | 2.0.0 | `Load` / `Save`（读旧 → 过滤 → 补写 → 临时文件 → 回写）+ 默认值，`config_test.go` |
| G07 | `internal/i18n` zh/en 字典 | 2.0.2 | `i18n.go` 界面字典 + `log.go` 执行日志字典；GUI / CLI / engine / nogui 全部引用，无硬编码文案 |
| G08 | 功能 1 `EnableGlobalNodeModules` | 2.0.0 | `.npmrc` 模板 `tmp=%q`（修复 L1），`engine_test.go` 断言引号 |
| G09 | 功能 2 `ResetNodeConfig` | 2.0.0 | 删除 `NODE_PATH` 与项目 `.npmrc` |
| G10 | 功能 3/4 `.next` 缓存 | 2.0.0 | `clean.go`：写同名文件 + `SetHidden`；功能 4 先 `ClearHidden` 再删除 |
| G11 | 功能 5 `InstallGlobalDeps` | 2.0.2 | 5 组 `npm install -g`，失败组记录双语日志、不中断后续组 |
| G12 | 功能 6/7 白名单清理 | 2.0.2 | `clean.go` `cleanByWhitelist`；v2.0.2 修正功能 7 保留 `node_modules` 的测试断言 |
| G13 | 命令行模式 | 2.0.2 | `main.go`：`-h` / `-lang` / 功能号 / 路径；危险操作 CLI `Y/N`，提示走 i18n |
| G14 | Fyne 主题 | 2.0.0 | `internal/ui/theme.go` 对齐 DESIGN-SYSTEM §2.1 设计令牌 |
| G15 | 主窗口 | 2.0.2 | `ui.go`：目录输入/选择/保存、语言切换、功能按钮 1–7、结果日志区、页脚；v2.0.2 补「关于」弹窗 |
| G16 | 危险操作二次确认 | 2.0.0 | `ui.go` `dialog.ShowConfirm`（功能 1/6/7）+ CLI `Y/N` |
| G17 | `cmd/repo-cleaner/main.go` 入口 | 2.0.2 | 参数解析 → CLI / GUI 分流；`gui.go`（默认）/ `headless.go`（nogui）按 build tag 分发 |
| G18 | 跨平台交叉编译脚本 | 2.0.2 | `scripts/build.ps1` / `build.sh`：6 目标 + `-ldflags -X` 注入 10 位 build 号；v2.0.2 显式 CGO 与交叉编译器检查，缺工具链跳过并警告 |
| G19 | `.gitignore` 忽略 Go 产物 | 2.0.0 | `bin/` 已忽略 |
| G20 | 全量门禁通过 | 2.0.2 | `gofmt -l .`（空）、`go vet -tags nogui ./cmd/... ./internal/...`、`go test -tags nogui ./cmd/... ./internal/...`、`go build -tags nogui ./...` 全绿；`npm run check:docs` / `type-check` / `lint` / `build` 通过。本机无 OpenGL C 工具链，默认 GUI 构建须在具备工具链的机器或 CI 验证 |
| G21 | 提交并推送 | 2.0.2 | 收尾变更提交至 `main` 并推送 `origin/main` |

## 五、v2.3.0（2026-11-08 完成）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| V1 | 支持递归扫描 monorepo 深层项目 | 中 | 已完成 | `-recursive` 参数 + GUI 复选框；`platform.ListProjectDirsDeep` 深度 3 层（覆盖 REPO_ROOT/monorepo/packages/app）、跳过依赖/产物目录、排序返回；`config.ini` 支持 `RECURSIVE`；文档与帮助同步 |
| V3 | 支持自定义清理规则 | 低 | 已完成 | `config.ini` 新增 `EXTRA_CLEAN_DIRS`（逗号分隔），并入功能 6 / 7 白名单，仅限项目目录内部匹配 |
| V4 | 添加日志记录功能 | 低 | 已完成 | 新增 `internal/logs` 文件日志器，CLI / GUI 执行日志同步追加写入 `repo-cleaner.log` |
| R0 | v2.3.0 收尾 | 高 | 已完成 | 版本单一来源 bump 至 2.3.0 并同步全部展示位；CHANGELOG 补 `[2.3.0]`；`gofmt` / `go vet` / `go test` / `go build`（`-tags nogui`）门禁复跑通过；提交并推送 |

## 六、v2.3.1 文档体系重新规划（2026-11-14 完成）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| DOC1 | 重新规划文档信息架构 | 高 | 已完成 | 新增 [`docs/TROUBLESHOOTING.md`](./TROUBLESHOOTING.md)（症状 → 原因 → 处理）与 [`docs/MAINTAINERS.md`](./MAINTAINERS.md)（命名 / 仓库设置 / 发布 / 治理 / 差距）；[`docs/README.md`](./README.md) §1 文档地图补齐全部文档并明确单一来源；根 `README.MD` 移除维护者专属的 GitHub 设置章节 |
| DOC2 | 修正文档与代码失真 | 高 | 已完成 | 递归扫描（默认直接子目录、开启后 3 层）、`internal/` 7 个包、`config.ini` 4 键与 `repo-cleaner.log`、单独路径被忽略、帮助参数与深度 2/3、`release.yml` 的 `needs` 关系、测试覆盖（10 文件 / 31 函数 / 6 包）、门禁命令与 `ci.yml` 对齐、目录树与工作区文件名 |
| DOC3 | 数据真实性与可读性 | 中 | 已完成 | 移除无依据的 `Saves 10GB+` / `Auto-Clean` 徽章；补 Release 资产命名与 `SHA256SUMS.txt` 校验、预演与日志 FAQ、排障入口；根 README 命令块改为引用文档中心 |
| DOC4 | 治理规则闭环 | 中 | 已完成 | `CONTRIBUTING.md` §3.2 / §9.2 / §10 与 `docs/README.md` §5 新增「新增文档四处登记」「任务单一来源」「数据真实性」「`check:docs` + `check:root-clean` 双门禁」；§8.2 版本同步清单补新增文档 |
| DOC5 | 差距与待办登记 | 中 | 已完成 | `MAINTAINERS.md` §6 记录 M1–M6 现状与决策，`tasks.md` 对应登记 D1–D6 代码侧待办，`CHANGELOG.MD` 补 Docs 条目 |

> 本节只归档**文档侧**成果；D1–D6 涉及代码 / workflow 改动，仍保留在 [tasks.md](./tasks.md)。

## 七、v2.3.2 落地页目录（2026-11-14 完成）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| DOC6 | 建立 Web 展示项目落地页目录 | 中 | 已完成 | 新建 `landing/` 与 `landing/README.md`（用途 / 边界 / 内容约定 / 门禁登记）；根目录白名单、四处文档登记与版本展示位同步；页面文件仍为规划项，见 [tasks.md](./tasks.md) D7 |

## 八、v2.3.2 任务核查归档（2026-11-14 17:58）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| D5 | 为 `cmd/` 与 `internal/ui` 补自动化测试，或固化手动冒烟清单 | 低 | 已完成（清单口径） | [`CONTRIBUTING.md` §9.3](./CONTRIBUTING.md) 已固化 GUI 手动冒烟清单（帮助 / `-lang` / 语言切换 / 关于 / 功能 1 / 3→4 / 6 / 7 / 路径异常），`ci.yml` 以 `-tags ci` 编译 GUI 代码层；`OPENSPEC.md` §9.1 引用该清单 |
| D2a | `release.yml` 的 `gate` 补 `check:root-clean` | 高 | 已完成 | v2.3.2：`gate` 首步与 `ci.yml` 对齐；D2 剩余「跨工作流阻断」仍在 [tasks.md](./tasks.md) |
| D3a | `-h` 帮助与执行日志递归深度统一为 3 层 | 中 | 已完成 | v2.3.2：`cmd/repo-cleaner/main.go` 的 `-recursive` 说明与帮助、`internal/i18n/log.go` 双语提示、`internal/ui/ui.go` 注释均为 3 层；D3 剩余「残留注释 / 规格 L6」仍在 tasks.md |
| W1 | CI GUI 构建补齐 Wayland 开发依赖 | 高 | 已完成 | v2.3.2：`ci.yml` / `release.yml` 安装 `libwayland-dev` / `libxkbcommon-dev`，解决 `wayland-client-core.h` 缺失；排障与规格文档同步 |

> 本节按 2026-11-14 17:58 的代码与 workflow 实况核查归档；「规划写成既成事实」仍为禁止项，未完成事项只在 [tasks.md](./tasks.md) 维护。

## 九、v2.3.3 剩余任务收口（2026-11-14 18:21）

| ID | 任务 | 优先级 | 状态 | 结果 |
|----|------|--------|------|------|
| V5 | 添加详细进度显示（GUI 进度条 + CLI 分步进度） | 中 | 已完成 | `internal/ui/output.go` 的 `progressView` 按执行日志 `[n/m]` 推进进度条与步骤文案；`progress_parse.go` 抽为纯函数并新增 `progress_parse_test.go`；CLI 原有 `[1/4]` / `[1/5]` 分步日志保留 |
| D1 | Release 资产按全局规范打包归档 | 中 | 已完成 | `release.yml` 新增独立 `version` job 读取 `internal/buildinfo/version.go`；`scripts/package-release.sh` 产出 `RepoManager-v<版本>-<os>-<arch>.zip / .tar.gz`（Windows/macOS zip、Linux tar.gz）与 `SHA256SUMS.txt`，Release 只上传归档 |
| D2 | 让同 SHA 的其它工作流失败阻断 release job | 高 | 已完成 | `release` job 以 `gh api` 校验同 SHA 全部非 release 工作流为绿，运行中轮询、失败或超时退出；`permissions` 补 `actions: read`；`gate` 的 `check:root-clean` 见第八节 D2a |
| D3 | 统一残留递归深度文本、明确帮助固定英文 | 低 | 已完成 | `options.go` / `recursive_extra_test.go` 注释改为 3 层；`OPENSPEC.md` L6、`TROUBLESHOOTING.md`、`DESIGN-SYSTEM.md` G9、根 README 同步；CLI 帮助固定英文，不随 `-lang` 切换 |
| D4 | 补齐繁体中文 | 中 | 已完成 | `internal/i18n/hant.go` 繁中界面 + 执行日志字典；`Normalize` 支持 `zh-Hant` / `tw` / `zh-TW`；GUI 下拉框、CLI `-lang`、`config.ini` 枚举、参考界面三语数据与落地页切换 UI 全覆盖；localStorage 持久化 + 枚举校验 + `<html lang>` 同步；`i18n_test.go` 新增断言 |
| D6 | 产品名同步到应用内文案与技术标识 | 高 | 已完成 | `go.mod` / import / `cmd/repo-manager/` / 二进制与日志名 / Fyne app ID / `package.json` / GitHub 链接 / Release 资产 / i18n 与 GUI/CLI 文案统一；CHANGELOG 与归档历史条目按 §8.3 保留原貌 |
| D7 | 实现 Web 展示落地页 | 低 | 已完成 | `landing/index.html` + `assets/styles.css` / `app.js` / `version.js`（`npm run landing:version` 构建期注入 `V{semver}· build yymmddhhmm`）；语义化 kebab-case id、三语切换、响应式与深色偏好、设计令牌对齐；部署 / 上线方式仍为规划项 |
| R7 | 版本与文档同步 | 中 | 已完成 | bump 至 2.3.3 并同步 `package.json` / `package-lock.json` / README 徽章与页脚 / `tokens.json` / 被改文件头注释；`CHANGELOG.MD` 新增 `[2.3.3]`；OPENSPEC / CONTRIBUTING / DESIGN-SYSTEM / MAINTAINERS / TROUBLESHOOTING / docs/README / landing/README 四处登记同步 |

> 本节按 2026-11-14 18:21 的代码与 workflow 实况归档；`docs/tasks.md` 已更新为无待办状态。已知限制（`windows/arm64` 无交叉工具链、`linux/arm64` 缺 arm64 架构开发库）登记于 [MAINTAINERS.md §6](./MAINTAINERS.md) 与 [CONTRIBUTING.md §8.4](./CONTRIBUTING.md)，不属待办。
