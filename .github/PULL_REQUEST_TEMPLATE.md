<!--
PR 标题格式与提交信息一致：<type>: <description>（如 fix: 修复功能4恢复 .next 缓存失效）
完整规范见 docs/CONTRIBUTING.md（§6 提交信息 / §9 代码审查流程），文档导航见 docs/README.md。
-->

## 变更总结

<!-- 一两句话说明改了什么，以及为什么改 -->

## 变更类型

- [ ] feat 新功能
- [ ] fix 缺陷修复
- [ ] docs 文档
- [ ] refactor 重构
- [ ] style 代码风格
- [ ] test 测试
- [ ] chore 构建 / 依赖 / 配置
- [ ] perf 性能
- [ ] ci CI/CD
- [ ] revert 回滚

## 关联与版本

- 关联 issue / PR：
- 本次版本号：`vX.Y.Z`（提交信息中已标注）
- 是否为破坏性变更：否 / 是（说明见下）

## 测试说明

<!-- 实际执行过的命令与结果，未执行的项目须说明原因 -->

- [ ] `npm run type-check`
- [ ] `npm run lint`
- [ ] `npm run check:docs`
- [ ] `npm run build`
- [ ] `RepoCleaner.bat` 手动冒烟（按 docs/CONTRIBUTING.md §9.3，仅改批处理时必填）

## 审查清单

- [ ] 行为与 `docs/OPENSPEC.md` 一致，变更功能已同步规格文档
- [ ] UI 对齐 `docs/DESIGN-SYSTEM.md` / `docs/tokens.json`，无硬编码颜色
- [ ] 主要容器与关键元素已加语义化 `id`
- [ ] 无 `console.log` / `debugger` / `any` / 未使用 import
- [ ] 代码文件单文件 ≤ 200 行，超出已拆分且导出契约不变
- [ ] 批处理：路径带引号、延迟展开正确、危险操作有二次确认
- [ ] 版本号已 bump，`package.json` / README 徽章 / CHANGELOG / 展示位一致
- [ ] `CHANGELOG.MD` 已新增版本小节；相关文档已同步
- [ ] 新增文档归入 `docs/`，Markdown 相对链接按实际大小写可达
- [ ] 无密钥 / Token / 真实用户数据；`node_modules`、`.next`、`*.log`、`*.tsbuildinfo` 未被跟踪

## 破坏性变更 / 回滚方案

<!-- 无则写「无」；有则说明影响范围、兼容性与回滚步骤 -->
