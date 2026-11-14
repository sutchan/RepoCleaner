# Repo Manager 落地页目录 (landing/) v2.3.2

> **定位**：保存 Web 展示项目的落地页面（landing page）源码与静态资源，是页面内容的唯一落点。
> 本轮只建立目录与用途说明，**尚未创建任何页面文件**（规划，当前未实现）；规范正文一律放 [`docs/`](../docs/README.md)。

---

## 1. 用途与边界

| 目录 | 放什么 | 不放什么 |
|------|--------|----------|
| [`docs/`](../docs/README.md) | 规格、规范、排障与维护手册 | 页面代码 |
| [`prototype/`](../prototype/README.md) | 原型说明与原型稿 | 可运行代码、规范正文 |
| `src/` | 已弃用的 Next.js 参考界面（v1.x 历史） | 新增展示页面 |
| `public/` | 参考界面静态资源 | 落地页资源 |
| `landing/`（本目录） | 落地页 HTML / CSS / JS、图片与字体、页面级说明 | 产品业务逻辑、规范正文、临时产物 |

---

## 2. 内容约定

1. 页面入口建议 `landing/index.html`，静态资源归入 `landing/assets/`，使用相对路径引用，便于静态托管直接预览。
2. 页面文案须与 [`OPENSPEC.md`](../docs/OPENSPEC.md) 的功能事实一致，禁止虚构容量、耗时、用户量等数据；产品名用 `Repo Manager`，技术标识（二进制 `repo-manager`、远程 `sutchan/RepoManager`）按实况书写，对照见 [`MAINTAINERS.md` §1](../docs/MAINTAINERS.md)。
3. 视觉对齐 [`DESIGN-SYSTEM.md`](../docs/DESIGN-SYSTEM.md) 与 [`tokens.json`](../docs/tokens.json)；主要容器与区块加语义化 kebab-case `id`（如 `landing-hero`、`landing-features`、`landing-footer`），便于可访问性关联与测试定位。
4. 多语言按「简中 → 繁中 → 英文」优先级规划；当前产品仅 `zh` / `en`，页面若先行须显式标注差距（见 [`OPENSPEC.md` §7.3](../docs/OPENSPEC.md) L7），不得写成已支持。
5. 页脚须展示版本与 build 编号，格式与构建期注入规则按 [`CONTRIBUTING.md` §8.2](../docs/CONTRIBUTING.md) 执行；禁止硬编码版本号或运行时现场生成。
6. 源代码文件单文件 ≤ 200 行（按职责拆分），文档文件不拆分；新增 JS / TS 须通过 `npm run lint` 与 `npm run type-check`，改动文档后跑 `npm run check:docs`。
7. 构建产物写 `landing/dist/` 或 `landing/out/`（已被 `.gitignore` 忽略）；一次性脚本、探针与中间产物写仓库根 `.tmp/`，不得散落在仓库根或本目录。

---

## 3. 门禁登记

本目录已登记 `scripts/check-root-clean.mjs` 的 `ALLOWED_DIRS` 白名单，保证 `npm run check:root-clean` 与 CI 首步通过；今后新增根目录条目须同步白名单与 `.gitignore`，见 [`MAINTAINERS.md` §5](../docs/MAINTAINERS.md)。

> 部署平台与上线方式尚未确定，属规划项；确定后同步本文与 [`MAINTAINERS.md`](../docs/MAINTAINERS.md)，不得写成既成事实。
