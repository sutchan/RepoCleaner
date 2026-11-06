# RepoCleaner 文档中心 v2.0.1

> 本目录是 RepoCleaner **所有规范与规格文档的唯一归档位置**。
> 仓库根只保留面向使用者的入口 [README.MD](../README.MD)、变更历史 [CHANGELOG.MD](../CHANGELOG.MD) 与 [LICENSE](../LICENSE)。
> `prototype/` 仅保存原型说明（见 [prototype/README.md](../prototype/README.md)），不再存放可运行代码或规范正文。

---

## 1. 文档地图

| 文档 | 定位 | 写什么 | 不写什么 |
|------|------|--------|----------|
| [README.MD](../README.MD) | 使用者入口 | 功能介绍、快速开始、参数、FAQ、文档索引 | 开发规范、版本历史明细 |
| [docs/OPENSPEC.md](./OPENSPEC.md) | 产品功能规格 | 7 项功能的行为契约、交互、环境变量、兼容性、已知限制 | 编码风格、提交/分支规则 |
| [docs/CONTRIBUTING.md](./CONTRIBUTING.md) | 开发规范（**唯一来源**） | 编码风格、命名约定、目录组织、提交格式、分支管理、代码审查、版本规则、术语表、安全红线 | 功能介绍、设计令牌 |
| [docs/DESIGN-SYSTEM.md](./DESIGN-SYSTEM.md) | UI 设计规范（**唯一来源**） | 设计令牌、组件库、交互、响应式、a11y、规范与实现差距清单 | 业务功能、开发流程 |
| [docs/tokens.json](./tokens.json) | 机器可读设计令牌 | 与 DESIGN-SYSTEM.md §2 对齐的颜色/字体/间距/圆角/动效值 | 不作为产品版本号来源 |
| [CHANGELOG.MD](../CHANGELOG.MD) | 唯一版本历史 | 按版本倒序的变更明细 | 功能规格、规范正文 |
| [prototype/README.md](../prototype/README.md) | 原型目录说明 | 原型现状与本目录用途 | 不复述设计规范正文 |
| [.github/PULL_REQUEST_TEMPLATE.md](../.github/PULL_REQUEST_TEMPLATE.md) | PR 审查清单 | 提交 PR 时逐项勾选 | — |

**单一来源原则**：同一事实只在一处维护。需要版本历史 → `CHANGELOG.MD`；需要功能行为 → `OPENSPEC.md`；需要协作流程 → `CONTRIBUTING.md`；需要视觉/组件规范 → `DESIGN-SYSTEM.md`。禁止跨文档复制正文。

### 1.1 Community Health Files（`.github/`）

GitHub 会自动识别仓库根、`docs/` 与 `.github/` 下的特定文件。本仓库的归属如下：

| 文件 | 位置 | 说明 |
|------|------|------|
| README | 根 `README.MD` | 使用者入口；**禁止**创建 `.github/readme.md`（与根 README 重复） |
| LICENSE | 根 `LICENSE` | MIT |
| CHANGELOG | 根 `CHANGELOG.MD` | 唯一版本历史 |
| CONTRIBUTING | `docs/CONTRIBUTING.md` | **唯一来源**，GitHub 已识别 `docs/` 下的 `CONTRIBUTING.md`，因此 `.github/` 不再放副本 |
| CODE_OF_CONDUCT | [`.github/CODE_OF_CONDUCT.md`](../.github/CODE_OF_CONDUCT.md) | 社区行为标准与举报方式 |
| SECURITY | [`.github/SECURITY.md`](../.github/SECURITY.md) | 安全策略与私密漏洞报告渠道（删除类工具必读） |
| SUPPORT | [`.github/SUPPORT.md`](../.github/SUPPORT.md) | 提问前自查、Issue 要求与不适用渠道 |
| FUNDING | `.github/FUNDING.yml` | 赞助配置；当前全注释占位，未填写任何账号 |
| ISSUE_TEMPLATE | `.github/ISSUE_TEMPLATE/` | `config.yml` + 缺陷报告 / 功能建议两份 Issue Form |
| PULL_REQUEST_TEMPLATE | [`.github/PULL_REQUEST_TEMPLATE.md`](../.github/PULL_REQUEST_TEMPLATE.md) | PR 审查清单 |
| CODEOWNERS | `.github/CODEOWNERS` | 默认审查人 |

**决策**：规范正文只留在 `docs/`，`.github/` 只放 GitHub 平台需要的入口与模板，避免同一规范两处维护。查看文件清单见 [docs/CONTRIBUTING.md §3.1](./CONTRIBUTING.md)。

---

## 2. 按角色阅读

| 你是… | 建议路径 |
|-------|----------|
| 使用者，想快速用起来 | [README.MD](../README.MD)「快速开始」→「功能详解」→「常见问题 FAQ」 |
| 贡献者，准备改代码 | [CONTRIBUTING.md](./CONTRIBUTING.md) §2 环境 → §3 目录 → §4～§5 编码与命名 → §6～§9 提交/分支/审查 |
| 想确认某项功能的行为 | [OPENSPEC.md](./OPENSPEC.md) §3 功能规格 → §7 已知限制 |
| 改 UI / 加组件 | [DESIGN-SYSTEM.md](./DESIGN-SYSTEM.md) §2 令牌 → §3 组件库 → §7 差距清单；令牌值见 [tokens.json](./tokens.json) |
| 发版 / bump 版本 | [CONTRIBUTING.md](./CONTRIBUTING.md) §8 → [CHANGELOG.MD](../CHANGELOG.MD) |
| 审查 PR | [.github/PULL_REQUEST_TEMPLATE.md](../.github/PULL_REQUEST_TEMPLATE.md) + [CONTRIBUTING.md](./CONTRIBUTING.md) §9.2 |
| 报缺陷 / 提建议 | [SUPPORT.md](../.github/SUPPORT.md) → 用 [Issue 模板](https://github.com/sutchan/RepoCleaner/issues/new/choose) 提交 |
| 报告安全漏洞 | [SECURITY.md](../.github/SECURITY.md)，走私密渠道，**不要公开提 Issue** |
| 社区行为问题 | [CODE_OF_CONDUCT.md](../.github/CODE_OF_CONDUCT.md) |

---

## 3. 仓库结构速览

```
RepoCleaner/
├── README.MD / CHANGELOG.MD / LICENSE          # 入口、变更历史、许可证
├── RepoCleaner.bat / config.ini / lang.ini      # 产品本体与运行配置
├── docs/                                       # 本目录：所有规范与规格
│   ├── README.md          # 文档中心（本文）
│   ├── OPENSPEC.md        # 产品功能规格
│   ├── CONTRIBUTING.md    # 开发规范
│   ├── DESIGN-SYSTEM.md   # UI 设计规范
│   └── tokens.json        # 设计令牌
├── prototype/README.md                         # 原型目录说明（无规范正文、无代码）
├── src/                                        # 参考界面应用（Next.js App Router）
│   ├── app/               # layout.tsx / page.tsx / globals.css
│   ├── components/        # icons.tsx + ui/
│   └── lib/               # i18n.ts / utils.ts / data/
├── public/                                     # 静态资源
├── scripts/check-doc-links.mjs                 # 文档链接校验门禁
├── .github/                                    # Community Health Files
│   ├── CODE_OF_CONDUCT.md                     # 行为准则
│   ├── SECURITY.md                            # 安全策略
│   ├── SUPPORT.md                             # 支持指南
│   ├── FUNDING.yml                            # 赞助配置（全注释占位）
│   ├── CODEOWNERS                             # 默认审查人
│   ├── PULL_REQUEST_TEMPLATE.md               # PR 审查模板
│   └── ISSUE_TEMPLATE/
│       ├── config.yml                         # 模板配置与联系入口
│       ├── bug_report.yml                     # 缺陷报告表单
│       └── feature_request.yml                # 功能建议表单
└── package.json / tsconfig.json / next.config.ts / ...  # 工程配置
```

完整目录规则见 [CONTRIBUTING.md §3](./CONTRIBUTING.md)。

---

## 4. 常用命令

在**仓库根目录**执行：

```bash
npm install            # 安装依赖
npm run dev            # 开发预览
npm run type-check     # TypeScript 类型检查
npm run lint           # ESLint
npm run check:docs     # 校验所有 Markdown 相对链接（含文件名大小写）
npm run build          # 生产构建
npm run start          # 启动构建产物
```

`RepoCleaner.bat` 无自动化测试，改动后按 [CONTRIBUTING.md §9.3](./CONTRIBUTING.md) 在隔离目录手动冒烟验证。

---

## 5. 维护约定

1. **改功能** → 同步 `OPENSPEC.md` + 根 `README.MD`；**改 UI** → 同步 `DESIGN-SYSTEM.md` + `tokens.json`；**改流程** → 同步 `CONTRIBUTING.md`。
2. 文档中的路径、命令、文件名、版本号必须与仓库实况一致。写文档前先实查 `git log -1` 与 `node -p "require('./package.json').version"`，禁止凭记忆填写。
3. **链接按实际文件名大小写书写**（根文档是 `README.MD` / `CHANGELOG.MD` 大写，`docs/` 下为小写）；GitHub / Linux 区分大小写。每次改动文档后必须执行 `npm run check:docs`（脚本：`scripts/check-doc-links.mjs`），逐段比对真实文件名，失效链接返回退出码 1。
4. 「规划但未实现」的内容必须显式标注，或记入 `DESIGN-SYSTEM.md` §7 差距清单，不得写成既成事实。
5. 文档文件**不拆分**（即使超过 200 行也保持单文件完整）；源代码文件单文件 ≤ 200 行须按职责拆分。
6. 每次文件改动至少 bump 一次 patch 版本，并同步 [CONTRIBUTING.md §8.2](./CONTRIBUTING.md) 列出的全部展示位。
7. 禁止创建 `.github/readme.md`（与根 README 重复）；`RepoCleaner.code-workspace` 禁止删除。
8. GitHub 平台入口（行为准则、安全、支持、Issue/PR 模板、CODEOWNERS）放 `.github/`，规范正文放 `docs/`；改动社区健康文件时同步本页 §1.1 与 [CONTRIBUTING.md §3.1](./CONTRIBUTING.md) 的文件清单。禁止在 `.github/` 复制 `docs/CONTRIBUTING.md` 正文。
9. `.github/*.yml`（FUNDING / Issue Forms / config）改动后须验证 YAML 可解析，且不得填写未经核实的账号、邮箱或承诺响应时间。

---

## 6. 术语速查

| 统一写法 | 禁止写法 |
|----------|----------|
| `Next.js` / `Nuxt` | NextJS / Nextjs / NuxtJS |
| 仓库根目录（`REPO_ROOT`，含多个项目的父目录） | 与「项目根目录」混用 |
| 项目根目录（单个项目，`REPO_ROOT` 的直接子目录） | — |
| 参考界面应用（根 `src/`） | prototype 应用、前端工程 |
| 批处理脚本（`RepoCleaner.bat`） | bat 脚本、主程序 |
| 功能 1 ～ 功能 7（名称见 OPENSPEC §3.0） | 功能一、F1 |

完整术语表见 [CONTRIBUTING.md §5](./CONTRIBUTING.md)。
