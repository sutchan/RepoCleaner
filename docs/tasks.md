# Repo Manager 任务清单 (docs/tasks.md)

> 最后更新：**2026-11-14**
> 维护约定：已完成任务归档至 [tasks-done.md](./tasks-done.md)；本文件仅保留进行中 / 待办任务。
> 全局版本唯一来源：`internal/buildinfo/version.go` 的 `Version`（当前 **2.3.1**）。

## 进行中 / 待办

| ID | 任务 | 优先级 | 状态 | 更新时间 | 备注 |
|----|------|--------|------|----------|------|
| V5 | 添加详细进度显示 | 低 | 待办 | 2026-11-07 | GUI 进度条 / CLI 分步输出 |
| D1 | Release 资产按全局规范打包为 `<产品名>-v<版本>-<os>-<arch>.<扩展名>` 归档后上传 | 中 | 待办 | 2026-11-14 | 对应 [MAINTAINERS.md §6](./MAINTAINERS.md) M1；当前直接上传 `dist/*` 二进制 + `SHA256SUMS.txt` |
| D2 | `release.yml` 的 `gate` 补 `check:root-clean`，并让新增工作流失败可阻断 `release` job | 高 | 待办 | 2026-11-14 | 对应 M2 / [CONTRIBUTING.md §8.4](./CONTRIBUTING.md)「目标」；当前 ci 与 release 无跨工作流阻断 |
| D3 | 统一递归深度文本为 3 层（`-h` 帮助、`options.go` / `ui.go` / 测试注释），明确帮助固定英文 | 中 | 待办 | 2026-11-14 | 对应 M3 / [OPENSPEC.md §7.3](./OPENSPEC.md) L6 |
| D4 | 评估补齐繁体中文（语言优先级：简中 → 繁中 → 英文） | 低 | 待办 | 2026-11-14 | 对应 M4 / [OPENSPEC.md §7.3](./OPENSPEC.md) L7 |
| D5 | 为 `cmd/repo-cleaner` 与 `internal/ui` 补自动化测试，或固化手动冒烟清单 | 低 | 待办 | 2026-11-14 | 对应 M5 / [CONTRIBUTING.md §9.3](./CONTRIBUTING.md) |
| D6 | 把产品名 `Repo Manager` 同步到应用内文案与技术标识（i18n、GUI 标题、`-h` 帮助、`go.mod`、二进制 / 日志名、Release 资产、远程） | 中 | 待办 | 2026-11-14 | 对应 M6 / [MAINTAINERS.md §1](./MAINTAINERS.md)；按清单一次性改动，勿分批改名 |

> v2.0.0 Go 迁移 G01–G21 已于 2026-11-06 全部完成（归档至 [tasks-done.md](./tasks-done.md) 第四节）；
> v2.1.0 的 dry-run 预演模式与 L4 修复已完成（归档至 [tasks-done.md](./tasks-done.md) 第三节 R6）；
> v2.2.0 的 Angular / Svelte 缓存清理、`config.ini` 跟踪修复与 GitHub Actions 发布自动化已完成（归档至 [tasks-done.md](./tasks-done.md) V2 / V6 行）；
> v2.3.0 的递归扫描 monorepo（V1）、自定义清理目录（V3）、执行日志落盘（V4）已于 2026-11-08 完成（归档至 [tasks-done.md](./tasks-done.md) 第五节）；
> v2.3.1 文档体系重新规划已完成（归档至 [tasks-done.md](./tasks-done.md) 第六节），上表 D1–D6 为审计后登记的**代码侧待办**。

## 若新增任务

按表格格式追加到上方「进行中 / 待办」，并标注优先级（高/中/低）、状态（进行中/待办）、更新时间（YYYY-MM-DD）；完成后移入 tasks-done.md。
