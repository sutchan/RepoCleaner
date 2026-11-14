# Repo Manager 使用与构建排障 (Troubleshooting) v2.3.2

> **定位**：面向使用者与贡献者的「症状 → 原因 → 处理」清单，内容以当前代码实况为准。
> 功能行为契约见 [OPENSPEC.md](./OPENSPEC.md)，使用入口见 [README.MD](../README.MD)，安全边界见 [SECURITY.md](../.github/SECURITY.md)。
> 报缺陷前先看 [SUPPORT.md](../.github/SUPPORT.md)；涉及删除越界、绕过二次确认或路径穿越，按 [SECURITY.md](../.github/SECURITY.md) 走私密渠道，不要公开提 Issue。

---

## 1. 排查前三步

1. **核对版本与构建**：执行 `repo-cleaner -h`，输出含 `Version`（页脚文本 + 7 位 VCS 短提交号）；GUI 版本显示在窗口页脚。
2. **查看执行日志**：可执行文件同目录下的 `repo-cleaner.log`，追加写入并带时间戳，CLI 与 GUI 均会落盘。
3. **用预演复现**：加 `-dry-run`（GUI 勾选「预演模式」）可查看将要写入 / 删除 / 安装的操作，不修改任何文件与环境变量，也不执行 `npm install`；预演下危险操作跳过二次确认。

---

## 2. 安装与启动

| 症状 | 原因 | 处理 |
|------|------|------|
| 本机 `go build` 报 OpenGL / C 工具链相关错误 | Fyne GUI 需要 CGO、系统 OpenGL、X11 与 Wayland 开发库 | 完整 GUI 构建需 C 工具链；无工具链时用 `go build -tags nogui ./...` 只编译非 GUI 部分（**不会产出可用 GUI**）。CI 的 `gui-build` job 在 Ubuntu 安装开发库后验证完整编译 |
| Linux 编译报 `wayland-client-core.h: No such file or directory` | `go-gl/glfw v3.4` 在 Linux 默认标签同时编译 X11 与 Wayland 后端，只装了 X11 开发库 | 安装 `libwayland-dev` 与 `libxkbcommon-dev`（CI 的 `gui-build` 已内置）；或改用 `-tags x11` 只编 X11 后端（产物不支持 Wayland，**不用于发布**） |
| 找不到 Release 资产 | 资产按 `repo-cleaner-<goos>-<goarch>[.exe]` 命名（Windows 带 `.exe`），另附 `SHA256SUMS.txt` | 按 OS / 架构选择对应文件并用校验和核对；`windows/arm64` 暂不产出（托管 runner 缺交叉工具链），可在本机用 `scripts/build.ps1` 补齐 |
| 启动后是 GUI 而不是预期的命令行操作 | 只有给出功能号 `1`–`7` 才进入非交互 CLI | 写成 `repo-cleaner <功能号> <路径>`；单独给路径会启动 GUI 并忽略该路径 |

> 前置要求：Windows / macOS / Linux；功能 1 / 5 需要 Node.js ≥ 18 与 npm；源码构建需 Go ≥ 1.25。

---

## 3. 目录与配置

| 症状 | 原因 | 处理 |
|------|------|------|
| 改了 `config.ini` 却没生效 | 配置文件读取自**可执行文件同目录**，不是仓库根目录 | 复制 [config.example.ini](./config.example.ini) 到 exe 同目录；该文件含本机路径，不入库 |
| 语言设置不生效 | 语言仅 `zh` / `en`，除 `en` 外的任何值（含空值、非法值）均回落 `zh` | 用 `LANG=en` / `-lang en`；`-lang` 会写回 `config.ini` |
| 目录没切换 | 优先级为 CLI 位置路径 > `config.ini` 的 `REPO_ROOT` > 平台默认（Windows `E:\Github`、Unix `~/Github`）；无功能号时路径被忽略 | 在 GUI 中选择并「保存」，或与功能号同用 CLI 位置路径 |
| 提示「路径不存在」/「不是目录」 | `platform.ValidateRoot` 校验失败 | CLI 报错并以非 0 退出；GUI 弹窗后返回目录选择。路径含空格时用引号包裹参数 |

`config.ini` 共 4 个键：`REPO_ROOT`、`LANG`、`RECURSIVE`、`EXTRA_CLEAN_DIRS`。保存时程序会去重同名键、保留注释与未知行；CLI 执行完成后会把本次根目录写回。

---

## 4. 清理结果不符预期

| 症状 | 原因 | 处理 |
|------|------|------|
| monorepo 深层项目没被清理 | 默认只枚举 `REPO_ROOT` 的直接子目录 | 启用 `-recursive`、GUI「递归扫描」复选框或 `RECURSIVE=true|1|yes`；递归深度 3 层，跳过 `node_modules` / `.next` / `.git` 等依赖与产物目录 |
| 帮助里的递归深度与规格不一致 | `-h` 帮助文本写「depth 2」，实际调用为 3 层 | **以 3 层为准**（见 [OPENSPEC.md §3.0](./OPENSPEC.md)）；文本与注释待统一，已登记 [tasks.md](./tasks.md) |
| 自定义目录没有被清理 | `EXTRA_CLEAN_DIRS` 只追加**目录名**（逗号分隔），且仅匹配项目目录内部 | 不要写路径或 `REPO_ROOT` 顶层目录名；留空表示无自定义项 |
| 功能 6 / 7 清掉的项不完全一致 | 两者白名单不同且有重叠（`dist`、`build`、`.nuxt`） | 属预期设计：功能 6 面向全量清理、功能 7 面向构建缓存，清单见 [OPENSPEC.md §3.6 / §3.7](./OPENSPEC.md) |
| Next.js 热更新变慢 | 功能 3 会生成带隐藏属性的 `.next` 文件阻止缓存生成 | 开发完成后运行功能 4 恢复 |

> 删除不可逆：功能 1 / 6 / 7 执行前需二次确认（GUI 弹窗 / CLI 输入 `Y`）；建议先在隔离目录用 `-dry-run` 预演。

---

## 5. 依赖与环境变量

| 症状 | 原因 | 处理 |
|------|------|------|
| 启用全局 `node_modules` 后项目找不到模块 | 项目改为从全局读取依赖 | 先运行功能 5 安装常用全局依赖，再运行功能 1 |
| 功能 5 报错但其他组仍继续 | 5 组 `npm install -g` 逐组检查退出码，单组失败不中断，结束后汇总失败组并返回错误 | 按输出定位失败组；CLI 退出码非 0。需要网络与 npm |
| `NODE_PATH` 在新终端不生效 | 功能 1 写入用户环境变量 | 重开终端后再验证；功能 2 可清除该变量并恢复 npm 默认行为 |
| `npm root -g` 相关报错 | 平台层获取全局模块路径失败 | 确认 npm 可用且在 PATH 中 |

---

## 6. 门禁与发布失败

| 门禁 | 失败时 |
|------|--------|
| `npm run check:root-clean` | 根目录出现未登记条目：临时文件移入 `.tmp/`、产物移入 `reports/`；确属工程清单的登记 `scripts/check-root-clean.mjs` 白名单并同步 `.gitignore` |
| `npm run check:docs` | Markdown 相对链接不存在或大小写不符：根文档是 `README.MD` / `CHANGELOG.MD`（大写），`docs/` 下为小写 `.md` |
| `gofmt -l ./cmd ./internal` | 有输出即格式不合规，对列出的文件执行 `gofmt -w` |
| `go vet` / `go test` / `go build`（`-tags nogui ./...`） | 按报错包定位；GUI 完整编译另需 C 工具链与 OpenGL / X11 / Wayland 开发库 |
| 发布校验失败 | tag 必须等于 `v` + `internal/buildinfo/version.go` 的 `Version`，且等于 `package.json.version`；`CHANGELOG.MD` 必须有 `## [X.Y.Z]` 小节。见 [MAINTAINERS.md §3](./MAINTAINERS.md) |

CI 命令清单见 [OPENSPEC.md §9.1](./OPENSPEC.md) 与 [`.github/workflows/ci.yml`](../.github/workflows/ci.yml)；发版顺序见 [CONTRIBUTING.md §8.4](./CONTRIBUTING.md)。

---

## 7. 报问题前

请在 [SUPPORT.md](../.github/SUPPORT.md) 的指引下提供：OS 与架构、`repo-cleaner -h` 的版本与 build 信息、复现命令、**脱敏后**的 `config.ini` 键与 `repo-cleaner.log` 片段。日志可能包含本机绝对路径，公开提交前务必脱敏。

安全漏洞（删除越界、绕过二次确认、路径穿越）走 [SECURITY.md](../.github/SECURITY.md) 的私密报告渠道；一般缺陷与建议走 [Issue 模板](https://github.com/sutchan/RepoCleaner/issues/new/choose)。
