# Repo Manager 维护者手册 (Maintainers) v2.3.3

> **定位**：仓库维护者的操作与决策记录——命名现状、GitHub 仓库设置、发版校验、文档治理、门禁与目录卫生。
> 使用者请从 [README.MD](../README.MD) 开始；功能行为以 [OPENSPEC.md](./OPENSPEC.md) 为准；编码 / 提交 / 审查规范以 [CONTRIBUTING.md](./CONTRIBUTING.md) 为准；未完成的任务以 [tasks.md](./tasks.md) 为唯一来源。
> 本文只记录维护动作与决策，不复制功能规格或规范正文。

---

## 1. 命名与仓库现状

仓库在不同层次使用了不同名称，维护时必须区分，避免误以为「改名已完成」：

| 层次 | 当前值 | 依据 |
|------|--------|------|
| 产品展示名 | `Repo Manager` | 根 `README.MD`、全部 `docs/` 文档标题与术语表 |
| Go 模块 / 可执行文件 / 日志 | `repo-manager` / `repo-manager.log` | `go.mod`、`cmd/repo-manager/`、`scripts/build.*`、`internal/logs` |
| Git 远程 | `https://github.com/sutchan/RepoManager.git` | `git remote -v` |
| 本地目录 / 工作区文件 | 目录名 `RepoManager`、文件 `RepoManager.code-workspace` | 本机磁盘实况，不进入代码与远程 |

**决策（2026-11-14，v2.3.3 完成）**：项目显示名 `Repo Manager` 与技术标识 `repo-manager`、远程 `sutchan/RepoManager` 已全量统一；应用内文案（GUI 标题、CLI 帮助、i18n 字典、Web 参考界面）与 Release 资产名同步。文档中的命令、路径与链接一律按实况书写；CHANGELOG 与 tasks-done 历史条目保留原貌。

技术标识同步清单（v2.3.3 已全部执行，后续改名须按同一清单整体推进）：`go.mod` 模块路径与全部 import、`cmd/` 目录、二进制与日志文件名、构建脚本与 Release 资产名、Fyne app ID、`package.json` 的 `name` / `homepage` / `repository` / `bugs`、README 徽章与链接、`.github/` 模板、应用内 i18n 文案与全部 `docs/` 文档标题与术语表，并重新打 tag。当前按上表书写，历史归档不追溯改写。

> `*.code-workspace` 是工作区配置，**禁止删除**；`scripts/check-root-clean.mjs` 按文件后缀放行，不列入白名单文件名。

---

## 2. GitHub 仓库设置（About）

以下为维护者在 **Settings → About** 手工填写的建议值，工具不会自动写入。描述与话题基于当前产品实况（Go 跨平台单二进制 + Fyne GUI + CLI）。

### 仓库描述（Description）

```
🧰 一键清理前端项目冗余：统一管理全局 node_modules、禁用 .next/.vite 缓存、清理 dist/build 垃圾。Go 编译的跨平台单文件工具（Windows / macOS / Linux，Fyne GUI + 命令行），支持 Next.js / React / Vue / Vite / Angular / Svelte，中文/英文双语界面。
```

### 主页（Website）

可留空，或填写仓库地址：`https://github.com/sutchan/RepoManager`

### 话题标签（Topics）

复制以下标签到 **About → Topics**（每行一个，小写、连字符分隔；GitHub 上限 20 个）：

```
go
cross-platform
gui
node-modules
cleanup-tool
disk-cleaner
nextjs
react
vuejs
vite
angular
svelte
npm
cache-cleaner
developer-tools
cli
repository-maintenance
frontend
```

---

## 3. 发布与版本同步

| 项 | 约定与实况 |
|----|------------|
| 版本权威源 | `internal/buildinfo/version.go` 的 `Version` 常量；同步清单见 [CONTRIBUTING.md §8.2](./CONTRIBUTING.md) |
| 发布流程 | 见 [CONTRIBUTING.md §8.4](./CONTRIBUTING.md)；工作流实现为 [`.github/workflows/ci.yml`](../.github/workflows/ci.yml) 与 [`.github/workflows/release.yml`](../.github/workflows/release.yml) |
| 触发与校验 | push `vX.Y.Z` tag；发布前校验 `tag == "v" + Version == package.json.version`，且 [CHANGELOG.MD](../CHANGELOG.MD) 存在 `## [X.Y.Z]` 小节，任一不满足即失败 |
| Release 资产 | `RepoManager-v<版本>-<os>-<arch>.zip / .tar.gz` 归档（Windows/macOS zip、Linux tar.gz）与 `SHA256SUMS.txt`；`windows/arm64` 因托管 runner 缺交叉工具链、`linux/arm64` 因缺 arm64 架构 GUI 开发库，均暂不产出（发布矩阵仅 4 目标，本地 `scripts/build.ps1` 可补齐） |
| Release 说明 | 从 `CHANGELOG.MD` 对应版本小节提取，标题为 tag 名 |
| 本机交叉编译 | `scripts/build.ps1` / `scripts/build.sh`，6 目标输出到 `bin/`；缺交叉编译器时跳过并警告 |

**发版前实查**（禁止按会话记忆写版本）：

```bash
git log -1 --oneline
git status --porcelain
node -p "require('./package.json').version"
grep 'const Version' internal/buildinfo/version.go
```

**当前门禁边界**：`release.yml` 的 `release` job 以 `needs: [gate, build]` 保证发布工作流自身全绿；`ci.yml` 由同一 tag push 独立触发，**不阻断** `release`。tag 推送后应确认两条工作流均为绿色，再对外宣布发布。

---

## 4. 文档治理

单一来源原则：同一事实只维护一处，其他位置用链接引用。

| 变更类型 | 必须同步 |
|----------|----------|
| 改功能行为 | [OPENSPEC.md](./OPENSPEC.md) §3 / §4 / §7 + 根 `README.MD` 功能详解 + `CHANGELOG.MD` |
| 改 UI / 令牌 | [DESIGN-SYSTEM.md](./DESIGN-SYSTEM.md) + [tokens.json](./tokens.json) |
| 改流程 / 命令 / CI | [CONTRIBUTING.md](./CONTRIBUTING.md) + [docs/README.md](./README.md) §4 |
| 新增 / 删除文档 | [docs/README.md](./README.md) §1 文档地图与 §3 目录树、根 `README.MD` 文档索引、本文 §4 |
| 改落地页 / 新增根目录子目录 | [`landing/README.md`](../landing/README.md) + 四处文档登记 + `scripts/check-root-clean.mjs` 白名单（见 §5） |
| 新增 / 完成任务 | [tasks.md](./tasks.md)（唯一未完成来源）/ [tasks-done.md](./tasks-done.md)（归档） |
| 维护者信息 / 仓库设置 | 本文 |
| 版本 bump | [CONTRIBUTING.md §8.2](./CONTRIBUTING.md) 全部展示位 + `CHANGELOG.MD` 新小节 |

**校验与写法**：

1. 改动文档后执行 `npm run check:docs`（相对链接存在且大小写精确）与 `npm run check:root-clean`（根目录无散落条目）。
2. 根文档是 `README.MD` / `CHANGELOG.MD`（大写扩展名），`docs/` 下为小写 `.md`；链接按实际文件名书写。
3. 路径、命令、文件名、版本号必须与仓库实况一致；数据型描述（数量、行数、命令数）须实算或引用唯一来源，禁止估算或虚构。
4. 「规划但未实现」必须显式标注，或记入 [DESIGN-SYSTEM.md §7](./DESIGN-SYSTEM.md) 差距清单；`CHANGELOG.MD` 历史小节不追溯改写。
5. 文档文件不拆分（即使超过 200 行）；源代码文件单文件 ≤ 200 行须按职责拆分。

---

## 5. 门禁与目录卫生

- 自动门禁命令以 [ci.yml](../.github/workflows/ci.yml) 为准（含 `check:root-clean`、`gofmt`、`go vet` / `go test` / `go build` 的 `nogui` 标签、`npm ci`、`check:docs`），完整说明见 [CONTRIBUTING.md §6.3](./CONTRIBUTING.md) 与 [OPENSPEC.md §9.1](./OPENSPEC.md)。发布侧另有 `scripts/package-release.sh`（归档）与 `npm run landing:version`（落地页版本注入，每次 bump 版本须重跑并同步 `landing/assets/version.js`）。
- 一次性脚本、探针、日志、中间产物写 `.tmp/`；工具正式产物写 `reports/` / `coverage/` / `test-results/`；`*.code-workspace` 按后缀放行。
- 新增根目录工程文件 / 子目录必须同时登记 `scripts/check-root-clean.mjs` 的 `ALLOWED_DIRS` 白名单，并核对 `.gitignore` 的忽略规则，二者不可脱节；`landing/` 已按此登记。落地页构建产物写 `landing/dist/` 或 `landing/out/`（由 `.gitignore` 忽略）。
- 产品运行时生成的 `config.ini` / `repo-manager.log` 位于可执行文件同目录，不属仓库根目录卫生约束。

---

## 6. 已知差距与决策记录

以下为 **v2.3.3 核查后的现状**（含已修复项，不是新的待办）；未完成任务统一见 [tasks.md](./tasks.md)：

| 编号 | 现状 | 目标 / 决策 |
|------|------|-------------|
| M1 | ~~已修复~~ v2.3.3：Release 只上传 `RepoManager-v<版本>-<os>-<arch>.zip / .tar.gz` 归档 | 独立 `version` job + `scripts/package-release.sh`，见 §3 |
| M2 | ~~已修复~~ v2.3.3：`gate` 含 `check:root-clean`，`release` job 校验同 SHA 的其它工作流全绿 | 跨工作流阻断见 §3 |
| M3 | ~~已修复~~ v2.3.3：帮助、日志、代码与测试注释统一为 3 层，帮助固定英文 | 见 [OPENSPEC.md §7.3](./OPENSPEC.md) L6 |
| M4 | ~~已修复~~ v2.3.3：i18n 三语齐全（zh / zh-Hant / en），切换 UI 与持久化覆盖全部语言 | 繁中别名与回退见 `internal/i18n` |
| M5 | 部分收敛：`internal/ui` 新增 `progress_parse_test.go`，`cmd/` 仍无自动化测试 | 发版前按 [CONTRIBUTING.md §9.3](./CONTRIBUTING.md) 手动冒烟（清单口径已归档 D5） |
| M6 | ~~已修复~~ v2.3.3：产品名与技术标识、远程、应用内文案全量统一 | 同步清单见 §1 |

> 修复任一项时同步更新本表、[tasks.md](./tasks.md) 与 [CHANGELOG.MD](../CHANGELOG.MD)；本表只记录决策，不替代任务跟踪。
