# RepoCleaner 任务清单 (docs/tasks.md)

> 最后更新：**2026-11-07**
> 维护约定：已完成任务归档至 [tasks-done.md](./tasks-done.md)；本文件仅保留进行中 / 待办任务。
> 全局版本唯一来源：`internal/buildinfo/version.go` 的 `Version`（当前 **2.2.0**）。

## 进行中 / 待办

| ID | 任务 | 优先级 | 状态 | 更新时间 | 备注 |
|----|------|--------|------|----------|------|
| V1 | 支持递归扫描 monorepo 深层项目 | 中 | 待办 | 2026-11-07 | 对应 `OPENSPEC.md` §7.3 L3，需重新评估作用域约定（§3.0） |
| V3 | 支持自定义清理规则 | 低 | 待办 | 2026-11-07 | 需扩展 `internal/config` 并保持默认值兼容 |
| V4 | 添加日志记录功能（执行日志写入本地文件） | 低 | 待办 | 2026-11-07 | 复用 `engine.Logger` 接口，注意 `.gitignore` 排除日志产物 |
| V5 | 添加详细进度显示 | 低 | 待办 | 2026-11-07 | GUI 进度条 / CLI 分步输出 |

> v2.0.0 Go 迁移 G01–G21 已于 2026-11-06 全部完成（归档至 [tasks-done.md](./tasks-done.md) 第四节）；
> v2.1.0 的 dry-run 预演模式与 L4 修复已完成（归档至 [tasks-done.md](./tasks-done.md) 第三节 R6）；
> v2.2.0 的 Angular / Svelte 缓存清理、`config.ini` 跟踪修复与 GitHub Actions 发布自动化已完成（归档至 [tasks-done.md](./tasks-done.md) V2 / V6 行）。

## 若新增任务

按表格格式追加到上方「进行中 / 待办」，并标注优先级（高/中/低）、状态（进行中/待办）、更新时间（YYYY-MM-DD）；完成后移入 tasks-done.md。
