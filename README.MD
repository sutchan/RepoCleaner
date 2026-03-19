# 🧰 Github 项目清洁工具箱

> 一键让所有前端项目保持干净：无冗余 node_modules、无 .next 缓存、无打包垃圾
> 适配：NextJS / React / Vue / 常规 Node 项目

## ✨ 功能介绍

本工具箱用于管理 `E:\Github\` 目录下所有项目，自动解决：
- 每个项目重复生成巨大的 `node_modules` 占用空间
- NextJS 项目产生大量 `.next` 缓存文件
- 项目打包产生 `dist`/`build` 垃圾文件
- 依赖混乱、无法统一管理的问题

## 🚀 一键工具箱（最核心）

### 文件：`Github项目清洁工具箱.bat`

**功能清单**
1. 启用全局 node_modules —— 项目永不生成本地依赖
2. 还原默认 Node 配置 —— 恢复系统原始状态
3. 清理 & 禁用 .next 缓存 —— 禁止 NextJS 生成缓存目录
4. 还原 NextJS 缓存 —— 恢复正常编译
5. 一键安装常用全局依赖 —— React/Vue/Next/axios 等开发库
6. 一键清理所有项目缓存 —— 删除 node_modules/.next/dist/build/lock 文件

## 📦 预装依赖（功能5）