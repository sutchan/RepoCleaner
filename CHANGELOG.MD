# Changelog

All notable changes to this project will be documented in this file.

---

## [2.3.2] - 2026-11-14

### Added
- **落地页目录** - 新建 `landing/` 与 `landing/README.md`，用于保存 Web 展示项目的落地页面；明确与 `docs/` / `prototype/` / `src/` / `public/` 的边界、内容约定与门禁登记。页面文件尚未创建（规划项，登记为 `docs/tasks.md` D7）

### Changed
- **根目录白名单** - `scripts/check-root-clean.mjs` 的 `ALLOWED_DIRS` 登记 `landing/`，保证 `npm run check:root-clean` 与 CI 首步通过

### Docs
- **四处登记同步** - `docs/README.md` §1 / §3、根 `README.MD` 索引与目录树、`CONTRIBUTING.md` §3.1 / §3.2、`MAINTAINERS.md` §4 / §5 同步；`prototype/README.md` 补落地页边界
- **版本号同步** - 权威源 `internal/buildinfo/version.go` 及 §8.2 全部展示位同步至 `2.3.2`

---

## [2.3.1] - 2026-11-14

### Added
- **固定临时目录** - 建立 `.tmp/`（一次性脚本 / 探针 / 日志 / 中间产物）与 `reports/`（工具正式产物），由 `.gitkeep` 固定目录、内容全部忽略
- **根目录清洁度门禁** - 新增 `scripts/check-root-clean.mjs` 与 `npm run check:root-clean`：白名单校验根目录 git 可见条目，拦截散落的临时脚本 / 日志 / 产物；纳入 `.github/workflows/ci.yml` 首步

### Changed
- **忽略与门禁排除** - `.gitignore` 补齐 `.tmp/` / `reports/` / `coverage/` / `test-results/` 与 `*.tmp` / `*.out` / `/*.txt` / `audit-*.json` / `*-result.json` / `.tmp-*` 等兜底；ESLint、`tsconfig.json` 与 `scripts/check-doc-links.mjs` 同步排除临时目录
- **构建产物不落根目录** - README 本机快速构建与 CI 的 GUI 编译验证统一输出至 `.tmp/repo-cleaner`
- **版本号同步** - 权威源 `internal/buildinfo/version.go` 及 §8.2 全部展示位同步至 `2.3.1`

### Docs
- **开发规范补充** - `CONTRIBUTING.md` §3.2 新增临时文件统一目录约定，§6.3 / §9.2 补门禁检查；`README.MD` / `docs/README.md` 目录树与常用命令同步
- **项目更名** - 产品显示名统一为「Repo Manager」（旧名 RepoCleaner），全部文档标题、正文、术语表与维护者手册 §1 决策同步；GitHub 仓库地址、Go 模块 / 二进制 / 命令路径按实况保留，待同步清单见 [MAINTAINERS.md §1](./docs/MAINTAINERS.md)
- **文档体系重新规划** - 新增 `docs/TROUBLESHOOTING.md`（症状 → 原因 → 处理）与 `docs/MAINTAINERS.md`（命名 / 仓库设置 / 发布 / 治理 / 已知差距）；`docs/README.md` 文档地图、四处登记规则、版本同步清单与审查清单补齐；根 `README.MD` 移除维护者专属章节并补排障入口
- **文档事实校验** - 修正递归扫描（默认直接子目录、开启后 3 层）、`internal/` 7 个包、`config.ini` 4 键与日志落盘、CLI 单独路径、发布 `needs` 关系、测试覆盖（10 文件 / 31 函数 / 6 包）、门禁命令与 `ci.yml` 对齐；移除无依据的 `Saves 10GB+` / `Auto-Clean` 徽章；代码侧差距登记为 `docs/tasks.md` D1–D6，完成项归档至 `docs/tasks-done.md` 第六节

---

## [2.3.0] - 2026-11-08

### Added
- **递归扫描 monorepo（V1）** - 新增 `-recursive` 命令行参数与 GUI「递归扫描」复选框；`platform.ListProjectDirsDeep` 以 3 层深度枚举深层项目（覆盖 `REPO_ROOT/monorepo/packages/app` 典型结构），跳过 `node_modules` / `.next` / `.git` 等依赖与产物目录，结果排序返回；`config.ini` 支持 `RECURSIVE=true|1|yes` 持久化
- **自定义清理目录（V3）** - `config.ini` 新增 `EXTRA_CLEAN_DIRS`（逗号分隔），并入功能 6 / 7 清理白名单，仅限项目目录内部匹配
- **执行日志落盘（V4）** - 新增 `internal/logs` 文件日志器：CLI 与 GUI 的执行日志同步追加写入可执行文件同目录的 `repo-cleaner.log`
- **文案与帮助** - i18n 新增递归扫描中英文案（`RecursiveLabel` / `RecursiveOn`）；`-h` 帮助补充 `-recursive` 说明

### Changed
- **选项优先级统一** - `DryRun` / `Recursive` / `ExtraCleanDirs` 按「命令行参数 > `config.ini` > 默认关闭」生效（对齐 `OPENSPEC.md` §4.5 / §5.2）
- **文档同步至 v2.3.0** - README 徽章与页脚、`package.json` / `package-lock.json`、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.3.0`；被改 Go 文件头注释同步 bump

---

## [2.2.0] - 2026-11-07

### Added
- **Angular / Svelte 缓存清理（V2）** - 功能 7 白名单新增 `.angular`、`.svelte-kit`（v2.2.0 起适合 Vite / Turborepo / Angular / Svelte 项目）；`WarningVite` 双语警告文案同步；新增 `TestCleanAngularSvelteCache`
- **配置示例模板** - 新增 `docs/config.example.ini`

### Changed
- **`config.ini` 不再入库** - 该文件含开发者本机 `REPO_ROOT` 路径，属运行时个人配置：`.gitignore` 排除、移出 Git 跟踪；`OPENSPEC.md` §5.2 / §5.3 / 新增 §9.1a 与 `CONTRIBUTING.md` 目录树同步标注
- **版本号同步** - `internal/buildinfo/version.go`、`package.json` / `package-lock.json`、README 徽章与页脚、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.2.0`

### Docs
- **发布规范与落地（`CONTRIBUTING.md` §8.4）** - 新增「发布流程与 Release 自动化」规范并同步落地：GitHub Actions 所有工作流全绿后自动发布 Release；新增 `.github/workflows/ci.yml`（门禁 + GUI 编译验证）与 `.github/workflows/release.yml`（tag `vX.Y.Z` 驱动，gate 重跑门禁 → 三平台矩阵构建 5 目标二进制 → tag 与版本权威源一致性校验 → SHA256SUMS + CHANGELOG 提取说明 → 自动创建 Release）；`docs/README.md` 摘要与 §9.2 审查清单同步；windows/arm64 因 runner 缺交叉工具链暂不产出（§8.4 已标注）
- **全量文档审计** - `SECURITY.md` 从「Windows 批处理工具」改写为 Go 跨平台产品实况（版本表、报告内容、安全边界 `os.ReadDir` 枚举 / dry-run 行为）；`SUPPORT.md` 与 Issue 模板（bug_report / feature_request）版本头同步 2.2.0、报告内容改为跨平台（OS 字段、版本获取方式、复现示例 `RepoCleaner.bat` → `repo-cleaner`）；README 补 Angular / Svelte 适配类型、功能 7 新清理项、`-dry-run` 参数与示例、workflows 目录、构建步骤细化；`CONTRIBUTING.md` §3.1 目录树补 `workflows/`；`OPENSPEC.md` §4.4 菜单示意标注 v1.x 历史来源；`prototype/README.md` 修正「正式」与已弃用的措辞冲突；`docs/README.md` §3 目录树从 v1.x 快照重写为 Go 产品实况（补 `cmd/` / `internal/` / `workflows/` / build 脚本 / 配置示例，`src/` 标注已弃用）、§4 补 Go 门禁与构建命令、§1 文档地图补 `config.example.ini`

---

## [2.1.0] - 2026-11-07

### Added
- **dry-run 预演模式** - `engine.Options.DryRun` 贯穿功能 1–7：仅输出「将写入 / 将删除 / 将设置环境变量 / 将安装」预演日志，不修改任何文件与环境变量；CLI 新增 `-dry-run` 参数，GUI 新增「预演模式」复选框（会话级选择，不写入 `config.ini`）；预演下危险操作跳过二次确认（对齐 `OPENSPEC.md` §4.7）
- **预演模式单元测试** - `internal/engine/dryrun_test.go` 覆盖功能 1 / 5 / 6 预演行为（不删除文件、不写配置、不执行 npm）与功能 5 失败汇总

### Changed
- **L4 修复** - 功能 5 检查每组 `npm install -g` 退出码，结束后按双语文案汇总失败组并返回错误（CLI 退出码非 0）；单组失败仍不中断后续组
- **版本号同步** - `internal/buildinfo/version.go`、`package.json` / `package-lock.json`、README 徽章与页脚、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.1.0`

---

## [2.0.2] - 2026-11-06

### Added
- **执行日志国际化** - `internal/i18n/log.go` 新增 zh / en 执行日志字典，`engine` / CLI / `nogui` 提示统一走 `i18n.Logs()`，切换语言后日志同步生效
- **GUI 关于区** - 主窗口顶部新增「关于」按钮弹窗，版本与构建信息仍只在页脚展示
- **macOS 隐藏属性** - `internal/platform` 的 Unix 实现补齐 `chflags hidden/nohidden`，命令缺失时降级为点前缀隐藏

### Changed
- **迁移收尾规范同步** - `OPENSPEC.md` 按 Go 实现重写功能 / 交互 / 兼容性描述；`CONTRIBUTING.md` §6.3 与 PR 模板改为 Go 门禁；README 修正唯一版本源与二次确认表述
- **交叉编译脚本** - `scripts/build.ps1` / `build.sh` 显式启用 CGO 并检查目标交叉编译器，缺工具链时跳过并警告，不静默产出不可运行文件
- **测试装置拆分** - `internal/engine/engine_test.go` 拆出 `testkit_test.go`，两个文件均 ≤200 行

### Fixed
- **删除类测试隔离** - `stubPlatform` 覆盖 `DefaultRepoRoot()` 为 `t.TempDir()`，测试不再触碰真实 `E:\Github`
- **功能 7 白名单断言写反** - 修正 `TestCleanViteCache` 对 `node_modules` 的保留断言，Go 门禁恢复全绿
- **吞错处理** - GUI 语言 / 配置保存、CLI 保存配置与 Windows `setx` 清空不再静默忽略错误

### Docs
- **迁移任务归档** - `docs/tasks.md` 的 G01–G21 全部回填完成并归档至 `docs/tasks-done.md`
- **版本号同步** - `internal/buildinfo/version.go`、`package.json` / `package-lock.json`、README 徽章与页脚、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.0.2`

---

## [2.0.1] - 2026-11-06

### Added
- **应用页脚构建标识** - Go Fyne 主窗口与已弃用 Web 参考界面的页脚统一展示 `vX.Y.Z · build yymmddhhmm`（10 位）；Go 构建脚本以 `-ldflags -X` 注入 `internal/buildinfo.BuildID`，未注入时从 VCS 构建时间推导，仍缺失则仅显示版本号
- **页脚格式单元测试** - `internal/buildinfo/version_test.go` 校验 `Footer()` 格式、回退与 `yymmddhhmm` 布局

### Changed
- **build 编号格式** - `next.config.ts` 由 12 位 `yyyymmddhhmm` 改为 10 位 `yymmddhhmm`；`internal/buildinfo.Footer()` 成为页脚文本唯一来源，CLI `-h` 改用「页脚文本 + VCS 短提交号」
- **规范同步** - `CONTRIBUTING.md` §8.2、`DESIGN-SYSTEM.md` §3.0 / §3.2 / §8、`OPENSPEC.md` §4.6、README 明确页脚为唯一展示位置、10 位格式、构建期注入与回退规则
- **版本号同步** - `internal/buildinfo/version.go`、`package.json` / `package-lock.json`、README 徽章与页脚、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.0.1`

### Fixed
- **页脚版本重复前缀** - `internal/ui/ui.go` 原 `"v"+buildVersion()` 输出 `vv2.0.0 · ...`，改为直接渲染 `buildinfo.Footer()`

---

## [2.0.0] - 2026-11-05

### Added
- **产品迁移至 Go（重大版本）** - 新增 Go 模块（`go.mod` + `cmd/repo-cleaner` + `internal/`），产品本体由 `RepoCleaner.bat` 迁移为 Go 编译的跨平台可执行文件 `repo-cleaner`（Windows / macOS / Linux 单二进制，无第三方运行时依赖）
- **Go 分层架构** - `internal/platform`（OS 差异抽象：目录扫描、删除、隐藏属性、环境变量持久化，按 `build tag` 分 Windows / Unix 实现）/ `internal/engine`（功能 1～7 平台无关业务）/ `internal/config`（`config.ini` 读写）/ `internal/i18n`（中英文案）/ `internal/ui`（Fyne GUI）
- **Fyne 桌面 GUI** - 主窗口功能按钮 1–7、语言实时切换、危险操作二次确认弹窗、结果日志区、关于区（版本 + 构建信息）
- **跨平台能力** - 功能 3 / 4 的隐藏属性按平台实现（Windows `attrib`、macOS `chflags`、Linux 点前缀 + 只读位）；`NODE_PATH` 持久化按平台实现；默认目录 Windows `E:\Github`、Unix `~/Github`
- **单元测试** - `internal/engine` / `internal/config` / `internal/i18n` 纯逻辑测试与 `internal/platform` 临时目录集成测试（`go test ./...`）
- **交叉编译脚本** - `scripts/build.ps1` / `scripts/build.sh` 一键产出 Windows（amd64 / arm64）、macOS（amd64 / arm64）、Linux（amd64 / arm64）单文件可执行程序

### Changed
- **L1 修复** - `.npmrc` 的 `tmp=` 值加引号，兼容含空格的 `TEMP` 路径（v1.x 已知限制，见 `OPENSPEC.md` §7.3）
- **版本单一来源** - 迁移至 `internal/buildinfo/version.go` 的 `Version` 常量；`package.json` 保持同步
- **弃用旧版** - `RepoCleaner.bat` 标记为已弃用（保留兼容）；`src/` Next.js 参考界面标记为已弃用（纯展示，不执行真实清理）
- **文档同步至 v2.0.0** - `OPENSPEC`（架构 / 功能 / 交互 / 技术 / 兼容 / 验证）、`CONTRIBUTING`（Go 环境 / 目录 / 风格 / 版本 / 审查 / 安全）、`DESIGN-SYSTEM`（Fyne 界面章节）、根 `README`（跨平台使用 / 构建 / 结构）
- **版本号同步** - README 徽章与页脚、`package.json` / `package-lock.json`、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `2.0.0`

### Security
- 删除目标仍只来自目录枚举（`internal/platform` 层），用户输入仅用于指定 `REPO_ROOT`；危险操作保留二次确认

---

## [1.2.6] - 2026-11-04

### Added
- **Community Health Files 全套（`.github/`）**
  - `CODE_OF_CONDUCT.md` — 基于 Contributor Covenant v2.1 的中文行为准则（承诺/标准/执行责任/范围/举报/四级执行指南/归属）
  - `SECURITY.md` — 安全策略：支持版本、GitHub Security Advisories 私密报告渠道、删除类工具的安全边界、不属于漏洞的情形、不承诺固定 SLA
  - `SUPPORT.md` — 提问前自查表、有效 Issue 应含信息、不适用渠道与免责
  - `FUNDING.yml` — 赞助配置占位（全注释，未填写任何账号）
  - `CODEOWNERS` — 默认审查人（用户名取自 remote 实查的 `sutchan`）
  - `ISSUE_TEMPLATE/config.yml` + `bug_report.yml` + `feature_request.yml` — Issue Forms，含脱敏提示、隔离目录复现确认与「安全漏洞勿公开」引导
- **PR 模板完善** — 补充行为准则 / 安全策略入口与贡献者确认项

### Changed
- **不复制贡献规范** — GitHub 已识别 `docs/CONTRIBUTING.md`，故 `.github/` 不再创建 `CONTRIBUTING.md` 副本（避免两处维护及优先级遮盖）；同时遵守「禁止创建 `.github/readme.md`」约定。README / LICENSE / CHANGELOG 继续留在根目录
- **文档同步** — `docs/README.md` 新增 §1.1 Community Health Files 归属表与维护约定，`docs/CONTRIBUTING.md`、`docs/OPENSPEC.md`、根 `README.MD` 的文档地图与目录树补齐 `.github/` 清单；README 新增「社区与贡献」章节
- **版本号同步** — `package.json`、`package-lock.json`、README 徽章与页脚、`docs/tokens.json`、各 `docs/*.md` 与 `prototype/README.md` 标题同步至 `1.2.6`

### Security
- `SECURITY.md` 明确「路径穿越 / 绕过二次确认 / 删除越界」值得私密报告，并列出既有的路径来源限制与二次确认机制
- `FUNDING.yml` 保留全注释占位，不填写未经核实的账号、邮箱或承诺响应时间；所有模板均提示日志脱敏与隔离目录复现

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
