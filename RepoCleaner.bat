@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion
title 🛠️ Github 项目清洁工具箱 - 无node_modules / 无.next / 全自动
mode con: cols=90 lines=38

:: ====================== 你的仓库根目录 ======================
set "REPO_ROOT=E:\Github"

:: ====================== 主菜单 ======================
:MENU
cls
echo.
echo         ======================================================================
echo                               🧰 Github 项目清洁工具箱
echo         ======================================================================
echo.
echo            📂 适用目录：%REPO_ROOT%
echo.
echo         [1] 启用全局 node_modules    ✅  项目永不生成依赖文件夹
echo         [2] 还原默认 Node 配置        ⚙️  恢复系统原始状态
echo.
echo         [3] 清理 & 禁用 .next 缓存    🚀  禁止生成缓存目录
echo         [4] 还原 NextJS 缓存          📁  恢复正常编译
echo.
echo         [5] 一键安装常用全局依赖      📦  React/Vue/Next/axios/工具
echo         [6] 一键清理所有项目缓存      🧹  删除 node_modules/.next/dist
echo.
echo         [0] 退出工具箱
echo.
echo         ======================================================================
echo.
set /p "CHOICE=👉 请输入数字选择功能："

if "%CHOICE%"=="1" goto GLOBAL_NODE
if "%CHOICE%"=="2" goto RESET_NODE
if "%CHOICE%"=="3" goto CLEAN_NEXT
if "%CHOICE%"=="4" goto RESET_NEXT
if "%CHOICE%"=="5" goto INSTALL_DEPS
if "%CHOICE%"=="6" goto CLEAN_ALL_CACHE
if "%CHOICE%"=="0" exit
goto MENU

:: ====================== 功能1：启用全局 node_modules ======================
:GLOBAL_NODE
cls
echo.
echo [1/4] 正在自动获取全局 node_modules 路径...
for /f "delims=" %%i in ('npm root -g') do set "GLOBAL_NODE_MODULES=%%i"
echo 路径：!GLOBAL_NODE_MODULES!

echo.
echo [2/4] 正在配置环境变量 NODE_PATH...
setx NODE_PATH "!GLOBAL_NODE_MODULES!" >nul 2>&1

echo.
echo [3/4] 正在为所有项目禁用本地 node_modules...
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\package.json" (
        echo 已配置：%%~nd
        (
            echo # 禁止生成本地依赖
            echo global=true
            echo prefix=!GLOBAL_NODE_MODULES!
            echo no-package-lock=true
            echo silent=true
        ) > "%%d\.npmrc"
    )
)

echo.
echo ======================================================================
echo ✅ 全局 node_modules 启用成功！
echo ℹ️ 首次使用请执行 菜单[5] 一键安装依赖
echo ======================================================================
echo.
pause
goto MENU

:: ====================== 功能2：还原默认 Node ======================
:RESET_NODE
cls
echo.
echo 正在删除 NODE_PATH 环境变量...
setx NODE_PATH "" >nul 2>&1
reg delete "HKCU\Environment" /v NODE_PATH /f >nul 2>&1

echo.
echo 正在清理项目 .npmrc 禁用配置...
for /d %%d in (!REPO_ROOT!\*) do (
    del /f /q "%%d\.npmrc" >nul 2>&1
)

echo.
echo ======================================================================
echo ✅ Node 已恢复默认设置！
echo ======================================================================
echo.
pause
goto MENU

:: ====================== 功能3：清理 & 禁用 .next ======================
:CLEAN_NEXT
cls
echo.
echo 正在扫描并清理 NextJS 项目缓存...
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\next.config.js" (
        echo 处理项目：%%~nd
        rd /s /q "%%d\.next" 2>nul
        echo. > "%%d\.next" 2>nul
        attrib +h +s +r "%%d\.next" 2>nul
    )
)

echo.
echo ======================================================================
echo ✅ .next 缓存已清理 + 禁用成功！
echo ======================================================================
echo.
pause
goto MENU

:: ====================== 功能4：还原 .next ======================
:RESET_NEXT
cls
echo.
echo 正在恢复 NextJS 正常缓存功能...
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\next.config.js" (
        echo 还原项目：%%~nd
        del /f /q /a:s /a:h /a:r "%%d\.next" 2>nul
        rd /s /q "%%d\.next" 2>nul
    )
)

echo.
echo ======================================================================
echo ✅ NextJS 已恢复正常编译！
echo ======================================================================
echo.
pause
goto MENU

:: ====================== 功能5：一键安装常用全局依赖 ======================
:INSTALL_DEPS
cls
echo.
echo ======================================================================
echo          📦 正在一键安装常用全局开发依赖
echo ======================================================================
echo.
echo 安装列表：React、Vue、Next、Axios、Express、pnpm、yarn、rimraf 等
echo 等待安装完成...
echo.

npm install -g ^
react react-dom ^
vue @vue/cli ^
next ^
axios express ^
pnpm yarn ^
cross-env dotenv ^
nodemon pm2 ^
rimraf mkdirp

echo.
echo ======================================================================
echo ✅ 所有常用开发依赖已全局安装完成！
echo ✅ 所有项目可直接使用，无需本地安装
echo ======================================================================
echo.
pause
goto MENU

:: ====================== 功能6：一键清理所有项目缓存 ======================
:CLEAN_ALL_CACHE
cls
echo.
echo ======================================================================
echo          🧹 正在清理所有项目缓存（node_modules/.next/dist）
echo ======================================================================
echo.
echo 警告：将永久删除所有项目的缓存文件，释放大量空间！
echo.
pause
echo.

for /d %%d in (!REPO_ROOT!\*) do (
    echo 正在清理：%%~nd
    :: 清理 node_modules
    rd /s /q "%%d\node_modules" 2>nul
    :: 清理 .next
    rd /s /q "%%d\.next" 2>nul
    :: 清理 dist
    rd /s /q "%%d\dist" 2>nul
    :: 清理 build
    rd /s /q "%%d\build" 2>nul
    :: 清理 .nuxt
    rd /s /q "%%d\.nuxt" 2>nul
    :: 清理缓存文件夹
    del /f /q "%%d\package-lock.json" 2>nul
    del /f /q "%%d\yarn.lock" 2>nul
    del /f /q "%%d\pnpm-lock.yaml" 2>nul
)

echo.
echo ======================================================================
echo ✅ 所有项目缓存清理完成！磁盘空间已大幅释放
echo ======================================================================
echo.
pause
goto MENU