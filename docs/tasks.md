# RepoCleaner 任务清单 (docs/tasks.md)

> 最后更新：**2026-11-06**
> 维护约定：已完成任务归档至 [tasks-done.md](./tasks-done.md)；本文件仅保留进行中 / 待办任务。
> 全局版本唯一来源：`internal/buildinfo/version.go` 的 `Version`（当前 **2.0.1**）。

## 进行中 / 待办

### v2.0.0 Go 迁移实施（产品本体：`repo-cleaner`）

| ID | 任务 | 优先级 | 状态 | 更新时间 | 备注 |
|----|------|--------|------|----------|------|
| G01 | 初始化 Go module（`go.mod`：module `repo-cleaner`，go 1.25） | 高 | 待办 | 2026-11-05 | |
| G02 | `internal/buildinfo/version.go`：`Version` 常量 + 构建信息 | 高 | 待办 | 2026-11-05 | 版本唯一来源 |
| G03 | `internal/platform` 接口定义（`Platform`：枚举项目目录 / 删除 / 隐藏属性 / 环境变量持久化 / 临时目录） | 高 | 待办 | 2026-11-05 | |
| G04 | `internal/platform/platform_windows.go`（`attrib` / `setx` + `reg` / `E:\Github` 默认值） | 高 | 待办 | 2026-11-05 | |
| G05 | `internal/platform/platform_unix.go`（`chflags` / shell profile / `~/Github` 默认值） | 高 | 待办 | 2026-11-05 | |
| G06 | `internal/config`：Load / Save（读旧 → 过滤目标键 → 补写 → 回写）+ 默认值 | 高 | 待办 | 2026-11-05 | |
| G07 | `internal/i18n`：zh / en 文案字典（菜单、提示、确认、结果） | 高 | 待办 | 2026-11-05 | |
| G08 | `internal/engine`：`EnableGlobalNodeModules`（`.npmrc` 模板，`tmp=` 加引号 → 修 L1） | 高 | 待办 | 2026-11-05 | |
| G09 | `internal/engine`：`ResetNodeConfig` | 高 | 待办 | 2026-11-05 | |
| G10 | `internal/engine`：`CleanNextCache` / `RestoreNextCache`（隐藏属性跨平台） | 高 | 待办 | 2026-11-05 | |
| G11 | `internal/engine`：`InstallGlobalDeps`（5 组 `npm install -g`） | 高 | 待办 | 2026-11-05 | |
| G12 | `internal/engine`：`CleanAllCaches` / `CleanViteCache` | 高 | 待办 | 2026-11-05 | |
| G13 | `internal/engine`：命令行模式（`-h` / `-lang` / 数字 / 路径） | 中 | 待办 | 2026-11-05 | |
| G14 | `internal/ui`：Fyne 主题（对齐 DESIGN-SYSTEM §2.1 配色） | 中 | 待办 | 2026-11-05 | |
| G15 | `internal/ui`：主窗口（目录选择 / 语言按钮 / 功能按钮 1–7 / 结果日志区 / 关于区） | 中 | 待办 | 2026-11-05 | |
| G16 | `internal/ui`：危险操作二次确认对话框（功能 1 / 6 / 7） | 中 | 待办 | 2026-11-05 | |
| G17 | `cmd/repo-cleaner/main.go` 入口组装（参数解析 → 命令行模式或 GUI） | 高 | 待办 | 2026-11-05 | |
| G18 | `scripts/build.ps1` + `scripts/build.sh` 跨平台交叉编译脚本 | 中 | 待办 | 2026-11-05 | Win/macOS/Linux × amd64/arm64 |
| G19 | `.gitignore` 补充 Go 构建产物（`bin/`） | 中 | 待办 | 2026-11-05 | |
| G20 | 门禁：`gofmt` + `go vet ./...` + `go test ./...` + `go build ./...` + `npm run check:docs` 全部通过 | 高 | 待办 | 2026-11-05 | |
| G21 | 提交并推送 | 高 | 待办 | 2026-11-05 | |

### 完成动作（执行后回填）

- G01–G20 完成后逐项更新状态；G21 推送后归档至 tasks-done.md。

## 若新增任务

按以下格式追加到上方「进行中 / 待办」表，并标注优先级（高/中/低）、状态（进行中/待办）、更新时间（YYYY-MM-DD）；完成后移入 tasks-done.md。
