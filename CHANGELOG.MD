# Changelog

All notable changes to this project will be documented in this file.

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
