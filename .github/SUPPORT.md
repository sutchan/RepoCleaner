# 支持指南 (Support) v2.2.0

RepoCleaner 由个人维护，**不提供实时支持，也不承诺响应时间或服务等级**。多数问题可先从文档中找到答案。

## 提问前请先查阅

| 你的问题 | 优先阅读 |
|----------|----------|
| 怎么安装和使用 | [README.MD](../README.MD)「快速开始」 |
| 某个功能是做什么的 | [README.MD](../README.MD)「功能详解」或 [OPENSPEC.md](../docs/OPENSPEC.md) §3 |
| 参数、语言、目录怎么配置 | [README.MD](../README.MD)「智能配置」 |
| 使用中的常见疑问 | [README.MD](../README.MD)「常见问题 FAQ」与「注意事项」 |
| 为什么某个行为不符合预期 | [OPENSPEC.md](../docs/OPENSPEC.md) §7「已知限制」 |
| 某个版本改了什么 | [CHANGELOG.MD](../CHANGELOG.MD) |
| 想参与开发 | [文档中心](../docs/README.md) → [CONTRIBUTING.md](../docs/CONTRIBUTING.md) |
| 界面 / 设计规范 | [DESIGN-SYSTEM.md](../docs/DESIGN-SYSTEM.md) |

## 提交 Issue

若文档未能解答，请在 [Issues](https://github.com/sutchan/RepoCleaner/issues) 使用对应模板提交：

- 缺陷报告（功能异常、路径问题、误删等）
- 功能建议（请先查看 [OPENSPEC.md](../docs/OPENSPEC.md) §10「扩展计划」，避免重复）

一个有效的报告应包含：

- RepoCleaner 版本号（GUI 页脚 / `repo-cleaner -h` 输出 / [CHANGELOG.MD](../CHANGELOG.MD)）
- 操作系统与版本（Windows / macOS / Linux）；涉及功能 1 / 5 时附 Node.js 与 npm 版本
- 运行方式（双击交互 / 命令行参数，给出完整命令）
- `REPO_ROOT` 的**目录结构示意**（可用假项目名，不要贴真实路径或私有仓库名）
- 复现步骤、预期行为、实际行为
- 日志或截图（**请脱敏**：移除真实用户名、绝对路径、密钥与代码内容）

> 所有删除类问题的复现请在**隔离的临时目录**进行，严禁在真实仓库根目录验证。

## 不适用的渠道

| 情况 | 正确做法 |
|------|----------|
| 安全漏洞（路径穿越、绕过二次确认、误删越界等） | **不要公开提 Issue**，按 [SECURITY.md](SECURITY.md) 私密报告 |
| 社区行为 / 骚扰问题 | 按 [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 私密举报 |
| 误删数据后的恢复 | 本项目无法恢复不可逆删除；请从备份或版本控制恢复 |
| npm / Node.js / Next.js 上游问题 | 向对应上游项目反馈 |

## 许可与免责

本项目按 [MIT License](../LICENSE) 提供，「按原样」分发，不附带任何担保。使用删除类功能前请自行确认备份。
