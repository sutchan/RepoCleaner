@echo off
chcp 65001 >nul 2>&1
setlocal enabledelayedexpansion
title RepoCleaner - Github Project Cleaner

:: ====================== Language Config ======================
set "LANG_FILE=%~dp0lang.ini"
set "LANG_SECTION=zh"

if exist "%LANG_FILE%" (
    for /f "usebackq tokens=1,2 delims==" %%a in ("%LANG_FILE%") do (
        if /i "%%a"=="LANG" set "LANG_SECTION=%%b"
    )
)

if "!LANG_SECTION!"=="en" goto :LOAD_EN
goto :LOAD_ZH

:: ====================== Load Chinese ======================
:LOAD_ZH
set "LANG_TITLE=选择语言 / Select Language"
set "LANG_1=English"
set "LANG_2=中文"
set "LANG_PROMPT=请输入选项 [1-2]:"

set "PATH_TITLE=选择代码库根目录"
set "PATH_1=使用当前目录:"
set "PATH_2=使用上级目录:"
set "PATH_3=手动输入路径"
set "PATH_4=使用默认:"
set "PATH_PROMPT=请输入选项 [1-4]:"
set "PATH_ERROR=错误: 路径不存在! 请重试..."
set "PATH_SAVE=是否保存为默认? [Y/N]:"
set "PATH_SAVED=配置已保存到 config.ini"

set "MENU_TITLE=RepoCleaner - GitHub 项目清理工具"
set "MENU_TARGET=目标目录:"
set "MENU_OPTION1=启用全局 node_modules     - 无需本地 node_modules"
set "MENU_OPTION2=重置 Node 配置            - 恢复默认设置"
set "MENU_OPTION3=清理 .next 缓存           - 禁用 NextJS 缓存"
set "MENU_OPTION4=恢复 NextJS 缓存          - 启用正常编译"
set "MENU_OPTION5=安装全局依赖              - React/Vue/Next/axios/工具"
set "MENU_OPTION6=清理所有项目缓存          - 删除 node_modules/.next/dist"
set "MENU_EXIT=退出"
set "MENU_CHOOSE=请输入选项 [0-6]:"

set "F1_TITLE=启用全局 node_modules"
set "F1_STEP1=[1/4] 获取全局 node_modules 路径..."
set "F1_STEP2=[2/4] 设置 NODE_PATH 环境变量..."
set "F1_STEP3=[3/4] 禁用所有项目的本地 node_modules..."
set "F1_SUCCESS=全局 node_modules 已启用!"
set "F1_INFO=请先运行 [5] 安装依赖"

set "F2_TITLE=重置 Node 配置"
set "F2_STEP1=删除 NODE_PATH 环境变量..."
set "F2_STEP2=清理项目 .npmrc 配置..."
set "F2_SUCCESS=Node 配置已重置!"

set "F3_TITLE=清理 .next 缓存"
set "F3_STEP1=扫描并清理 NextJS 缓存..."
set "F3_SUCCESS=.next 缓存已清理并禁用!"

set "F4_TITLE=恢复 NextJS 缓存"
set "F4_STEP1=恢复 NextJS 缓存功能..."
set "F4_SUCCESS=NextJS 缓存已恢复!"

set "F5_TITLE=安装常用全局依赖"
set "F5_LIST=安装列表: React, Vue, Next, Axios, Express, pnpm, yarn, rimraf 等"
set "F5_WAIT=请稍候..."
set "F5_STEP1=[1/5] 安装 React 生态..."
set "F5_STEP2=[2/5] 安装 Vue 生态..."
set "F5_STEP3=[3/5] 安装 HTTP 库..."
set "F5_STEP4=[4/5] 安装包管理器..."
set "F5_STEP5=[5/5] 安装开发工具..."
set "F5_SUCCESS=所有依赖已安装!"
set "F5_INFO=项目可直接使用，无需本地安装"

set "F6_TITLE=清理所有项目缓存"
set "F6_WARNING=警告: 这将永久删除所有缓存文件!"
set "F6_SUCCESS=所有缓存已清理!"

set "OK=[完成]"
set "ERROR=[错误]"
set "PAUSE=按任意键继续..."

goto :LANG_LOADED

:: ====================== Load English ======================
:LOAD_EN
set "LANG_TITLE=Select Language / 选择语言"
set "LANG_1=English"
set "LANG_2=中文"
set "LANG_PROMPT=Enter choice [1-2]:"

set "PATH_TITLE=Select Repository Root Directory"
set "PATH_1=Use Current Directory:"
set "PATH_2=Use Parent Directory:"
set "PATH_3=Manual Input Path"
set "PATH_4=Use Default:"
set "PATH_PROMPT=Enter choice [1-4]:"
set "PATH_ERROR=ERROR: Path does not exist! Please try again..."
set "PATH_SAVE=Save as default? [Y/N]:"
set "PATH_SAVED=Config saved to config.ini"

set "MENU_TITLE=RepoCleaner - Github Project Cleaner"
set "MENU_TARGET=Target:"
set "MENU_OPTION1=Enable Global node_modules     - No local node_modules"
set "MENU_OPTION2=Reset Node Config             - Restore default settings"
set "MENU_OPTION3=Clean .next Cache             - Disable NextJS cache"
set "MENU_OPTION4=Restore NextJS Cache          - Enable normal compilation"
set "MENU_OPTION5=Install Global Dependencies   - React/Vue/Next/axios/tools"
set "MENU_OPTION6=Clean All Project Cache       - Remove node_modules/.next/dist"
set "MENU_EXIT=Exit"
set "MENU_CHOOSE=Enter choice [0-6]:"

set "F1_TITLE=Enable Global node_modules"
set "F1_STEP1=[1/4] Getting global node_modules path..."
set "F1_STEP2=[2/4] Setting NODE_PATH environment variable..."
set "F1_STEP3=[3/4] Disabling local node_modules for all projects..."
set "F1_SUCCESS=Global node_modules Enabled!"
set "F1_INFO=Run [5] to install dependencies first time"

set "F2_TITLE=Reset Node Config"
set "F2_STEP1=Removing NODE_PATH environment variable..."
set "F2_STEP2=Cleaning project .npmrc configs..."
set "F2_SUCCESS=Node Config Reset!"

set "F3_TITLE=Clean .next Cache"
set "F3_STEP1=Scanning and cleaning NextJS cache..."
set "F3_SUCCESS=.next Cache Cleaned and Disabled!"

set "F4_TITLE=Restore NextJS Cache"
set "F4_STEP1=Restoring NextJS cache functionality..."
set "F4_SUCCESS=NextJS Cache Restored!"

set "F5_TITLE=Installing Common Global Dependencies"
set "F5_LIST=Install List: React, Vue, Next, Axios, Express, pnpm, yarn, rimraf, etc."
set "F5_WAIT=Please wait..."
set "F5_STEP1=[1/5] Installing React ecosystem..."
set "F5_STEP2=[2/5] Installing Vue ecosystem..."
set "F5_STEP3=[3/5] Installing HTTP libraries..."
set "F5_STEP4=[4/5] Installing package managers..."
set "F5_STEP5=[5/5] Installing dev tools..."
set "F5_SUCCESS=All Dependencies Installed!"
set "F5_INFO=Projects can use them directly without local install"

set "F6_TITLE=Cleaning All Project Cache"
set "F6_WARNING=WARNING: This will permanently delete all cache files!"
set "F6_SUCCESS=All Cache Cleaned!"

set "OK=[OK]"
set "ERROR=[ERROR]"
set "PAUSE=Press any key to continue..."

:LANG_LOADED

:: ====================== Repository Root Config ======================
set "DEFAULT_REPO_ROOT=E:\Github"

if not "%~1"=="" (
    set "REPO_ROOT=%~1"
    goto :ROOT_SET
)

if exist "%~dp0config.ini" (
    for /f "usebackq tokens=1,2 delims==" %%a in ("%~dp0config.ini") do (
        if /i "%%a"=="REPO_ROOT" set "REPO_ROOT=%%b"
    )
    if defined REPO_ROOT goto :ROOT_SET
)

:: ====================== Language Selection ======================
:LANG_MENU
cls
echo.
echo ======================================================================
echo                !LANG_TITLE!
echo ======================================================================
echo.
echo    [1] !LANG_1!
echo    [2] !LANG_2!
echo.
set /p "LANG_CHOICE=!LANG_PROMPT!"

if "!LANG_CHOICE!"=="1" (
    set "LANG_SECTION=en"
    goto :LOAD_EN_SKIP
)
if "!LANG_CHOICE!"=="2" (
    set "LANG_SECTION=zh"
    goto :LOAD_ZH_SKIP
)
if not defined LANG_CHOICE set "LANG_SECTION=zh"

:LOAD_EN_SKIP
if "!LANG_SECTION!"=="en" (
    set "LANG_TITLE=Select Language / 选择语言"
    set "LANG_1=English"
    set "LANG_2=中文"
    set "LANG_PROMPT=Enter choice [1-2]:"

    set "PATH_TITLE=Select Repository Root Directory"
    set "PATH_1=Use Current Directory:"
    set "PATH_2=Use Parent Directory:"
    set "PATH_3=Manual Input Path"
    set "PATH_4=Use Default:"
    set "PATH_PROMPT=Enter choice [1-4]:"
    set "PATH_ERROR=ERROR: Path does not exist! Please try again..."
    set "PATH_SAVE=Save as default? [Y/N]:"
    set "PATH_SAVED=Config saved to config.ini"

    set "MENU_TITLE=RepoCleaner - Github Project Cleaner"
    set "MENU_TARGET=Target:"
    set "MENU_OPTION1=Enable Global node_modules     - No local node_modules"
    set "MENU_OPTION2=Reset Node Config             - Restore default settings"
    set "MENU_OPTION3=Clean .next Cache             - Disable NextJS cache"
    set "MENU_OPTION4=Restore NextJS Cache          - Enable normal compilation"
    set "MENU_OPTION5=Install Global Dependencies   - React/Vue/Next/axios/tools"
    set "MENU_OPTION6=Clean All Project Cache       - Remove node_modules/.next/dist"
    set "MENU_EXIT=Exit"
    set "MENU_CHOOSE=Enter choice [0-6]:"

    set "F1_TITLE=Enable Global node_modules"
    set "F1_STEP1=[1/4] Getting global node_modules path..."
    set "F1_STEP2=[2/4] Setting NODE_PATH environment variable..."
    set "F1_STEP3=[3/4] Disabling local node_modules for all projects..."
    set "F1_SUCCESS=Global node_modules Enabled!"
    set "F1_INFO=Run [5] to install dependencies first time"

    set "F2_TITLE=Reset Node Config"
    set "F2_STEP1=Removing NODE_PATH environment variable..."
    set "F2_STEP2=Cleaning project .npmrc configs..."
    set "F2_SUCCESS=Node Config Reset!"

    set "F3_TITLE=Clean .next Cache"
    set "F3_STEP1=Scanning and cleaning NextJS cache..."
    set "F3_SUCCESS=.next Cache Cleaned and Disabled!"

    set "F4_TITLE=Restore NextJS Cache"
    set "F4_STEP1=Restoring NextJS cache functionality..."
    set "F4_SUCCESS=NextJS Cache Restored!"

    set "F5_TITLE=Installing Common Global Dependencies"
    set "F5_LIST=Install List: React, Vue, Next, Axios, Express, pnpm, yarn, rimraf, etc."
    set "F5_WAIT=Please wait..."
    set "F5_STEP1=[1/5] Installing React ecosystem..."
    set "F5_STEP2=[2/5] Installing Vue ecosystem..."
    set "F5_STEP3=[3/5] Installing HTTP libraries..."
    set "F5_STEP4=[4/5] Installing package managers..."
    set "F5_STEP5=[5/5] Installing dev tools..."
    set "F5_SUCCESS=All Dependencies Installed!"
    set "F5_INFO=Projects can use them directly without local install"

    set "F6_TITLE=Cleaning All Project Cache"
    set "F6_WARNING=WARNING: This will permanently delete all cache files!"
    set "F6_SUCCESS=All Cache Cleaned!"

    set "OK=[OK]"
    set "ERROR=[ERROR]"
    set "PAUSE=Press any key to continue..."

    goto :INTERACTIVE_INPUT
)

:LOAD_ZH_SKIP
set "LANG_TITLE=选择语言 / Select Language"
set "LANG_1=English"
set "LANG_2=中文"
set "LANG_PROMPT=请输入选项 [1-2]:"

set "PATH_TITLE=选择代码库根目录"
set "PATH_1=使用当前目录:"
set "PATH_2=使用上级目录:"
set "PATH_3=手动输入路径"
set "PATH_4=使用默认:"
set "PATH_PROMPT=请输入选项 [1-4]:"
set "PATH_ERROR=错误: 路径不存在! 请重试..."
set "PATH_SAVE=是否保存为默认? [Y/N]:"
set "PATH_SAVED=配置已保存到 config.ini"

set "MENU_TITLE=RepoCleaner - GitHub 项目清理工具"
set "MENU_TARGET=目标目录:"
set "MENU_OPTION1=启用全局 node_modules     - 无需本地 node_modules"
set "MENU_OPTION2=重置 Node 配置            - 恢复默认设置"
set "MENU_OPTION3=清理 .next 缓存           - 禁用 NextJS 缓存"
set "MENU_OPTION4=恢复 NextJS 缓存          - 启用正常编译"
set "MENU_OPTION5=安装全局依赖              - React/Vue/Next/axios/工具"
set "MENU_OPTION6=清理所有项目缓存          - 删除 node_modules/.next/dist"
set "MENU_EXIT=退出"
set "MENU_CHOOSE=请输入选项 [0-6]:"

set "F1_TITLE=启用全局 node_modules"
set "F1_STEP1=[1/4] 获取全局 node_modules 路径..."
set "F1_STEP2=[2/4] 设置 NODE_PATH 环境变量..."
set "F1_STEP3=[3/4] 禁用所有项目的本地 node_modules..."
set "F1_SUCCESS=全局 node_modules 已启用!"
set "F1_INFO=请先运行 [5] 安装依赖"

set "F2_TITLE=重置 Node 配置"
set "F2_STEP1=删除 NODE_PATH 环境变量..."
set "F2_STEP2=清理项目 .npmrc 配置..."
set "F2_SUCCESS=Node 配置已重置!"

set "F3_TITLE=清理 .next 缓存"
set "F3_STEP1=扫描并清理 NextJS 缓存..."
set "F3_SUCCESS=.next 缓存已清理并禁用!"

set "F4_TITLE=恢复 NextJS 缓存"
set "F4_STEP1=恢复 NextJS 缓存功能..."
set "F4_SUCCESS=NextJS 缓存已恢复!"

set "F5_TITLE=安装常用全局依赖"
set "F5_LIST=安装列表: React, Vue, Next, Axios, Express, pnpm, yarn, rimraf 等"
set "F5_WAIT=请稍候..."
set "F5_STEP1=[1/5] 安装 React 生态..."
set "F5_STEP2=[2/5] 安装 Vue 生态..."
set "F5_STEP3=[3/5] 安装 HTTP 库..."
set "F5_STEP4=[4/5] 安装包管理器..."
set "F5_STEP5=[5/5] 安装开发工具..."
set "F5_SUCCESS=所有依赖已安装!"
set "F5_INFO=项目可直接使用，无需本地安装"

set "F6_TITLE=清理所有项目缓存"
set "F6_WARNING=警告: 这将永久删除所有缓存文件!"
set "F6_SUCCESS=所有缓存已清理!"

set "OK=[完成]"
set "ERROR=[错误]"
set "PAUSE=按任意键继续..."

goto :INTERACTIVE_INPUT

:INTERACTIVE_INPUT
cls
echo.
echo ======================================================================
echo               !PATH_TITLE!
echo ======================================================================
echo.
echo    [1] !PATH_1! %CD%
echo    [2] !PATH_2! %~dp0..\
echo    [3] !PATH_3!
echo    [4] !PATH_4! %DEFAULT_REPO_ROOT%
echo.
set /p "PATH_CHOICE=!PATH_PROMPT!"

if "%PATH_CHOICE%"=="1" set "REPO_ROOT=%CD%"
if "%PATH_CHOICE%"=="2" set "REPO_ROOT=%~dp0..\"
if "%PATH_CHOICE%"=="3" (
    set /p "REPO_ROOT=Enter full path: "
)
if "%PATH_CHOICE%"=="4" set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not defined REPO_ROOT set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not exist "!REPO_ROOT!" (
    echo.
    echo    !PATH_ERROR!
    echo.
    pause
    goto :INTERACTIVE_INPUT
)

echo.
set /p "SAVE_CONFIG=!PATH_SAVE!"
if /i "!SAVE_CONFIG!"=="Y" (
    echo REPO_ROOT=!REPO_ROOT! > "%~dp0config.ini"
    echo.
    echo    !PATH_SAVED!
    timeout /t 1 >nul 2>&1
)

:ROOT_SET
if "!REPO_ROOT:~-1!"=="\" set "REPO_ROOT=!REPO_ROOT:~0,-1!"

:: ====================== Main Menu ======================
:MENU
cls
echo.
echo         ======================================================================
echo                        !MENU_TITLE!
echo         ======================================================================
echo.
echo            !MENU_TARGET! !REPO_ROOT!
echo.
echo         [1] !MENU_OPTION1!
echo         [2] !MENU_OPTION2!
echo.
echo         [3] !MENU_OPTION3!
echo         [4] !MENU_OPTION4!
echo.
echo         [5] !MENU_OPTION5!
echo         [6] !MENU_OPTION6!
echo.
echo         [L] !LANG_1! / !LANG_2!
echo         [0] !MENU_EXIT!
echo.
echo         ======================================================================
echo.
set /p "CHOICE=!MENU_CHOOSE!"

if "%CHOICE%"=="1" goto GLOBAL_NODE
if "%CHOICE%"=="2" goto RESET_NODE
if "%CHOICE%"=="3" goto CLEAN_NEXT
if "%CHOICE%"=="4" goto RESET_NEXT
if "%CHOICE%"=="5" goto INSTALL_DEPS
if "%CHOICE%"=="6" goto CLEAN_ALL_CACHE
if /i "%CHOICE%"=="L" goto :LANG_MENU
if "%CHOICE%"=="0" exit
goto MENU

:: ====================== Function 1: Enable Global node_modules ======================
:GLOBAL_NODE
cls
echo.
echo !F1_STEP1!
for /f "delims=" %%i in ('npm root -g') do set "GLOBAL_NODE_MODULES=%%i"
echo Path: !GLOBAL_NODE_MODULES!

echo.
echo !F1_STEP2!
setx NODE_PATH "!GLOBAL_NODE_MODULES!" >nul 2>&1

echo.
echo !F1_STEP3!
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\package.json" (
        echo Configured: %%~nd
        (
            echo # Disable local dependencies
            echo global=true
            echo prefix=!GLOBAL_NODE_MODULES!
            echo no-package-lock=true
            echo silent=true
        ) > "%%d\.npmrc"
    )
)

echo.
echo ======================================================================
echo !OK! !F1_SUCCESS!
echo !F1_INFO!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 2: Reset Node Config ======================
:RESET_NODE
cls
echo.
echo !F2_STEP1!
setx NODE_PATH "" >nul 2>&1
reg delete "HKCU\Environment" /v NODE_PATH /f >nul 2>&1

echo.
echo !F2_STEP2!
for /d %%d in (!REPO_ROOT!\*) do (
    del /f /q "%%d\.npmrc" >nul 2>&1
)

echo.
echo ======================================================================
echo !OK! !F2_SUCCESS!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 3: Clean .next Cache ======================
:CLEAN_NEXT
cls
echo.
echo !F3_STEP1!
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\next.config.js" (
        echo Processing: %%~nd
        rd /s /q "%%d\.next" 2>nul
        echo. > "%%d\.next" 2>nul
        attrib +h +s +r "%%d\.next" 2>nul
    )
)

echo.
echo ======================================================================
echo !OK! !F3_SUCCESS!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 4: Restore .next ======================
:RESET_NEXT
cls
echo.
echo !F4_STEP1!
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\next.config.js" (
        echo Restoring: %%~nd
        del /f /q /a:s /a:h /a:r "%%d\.next" 2>nul
        rd /s /q "%%d\.next" 2>nul
    )
)

echo.
echo ======================================================================
echo !OK! !F4_SUCCESS!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 5: Install Global Dependencies ======================
:INSTALL_DEPS
cls
echo.
echo ======================================================================
echo        !F5_TITLE!
echo ======================================================================
echo.
echo !F5_LIST!
echo !F5_WAIT!
echo.

echo.
echo !F5_STEP1!
npm install -g react react-dom next

echo.
echo !F5_STEP2!
npm install -g vue @vue/cli

echo.
echo !F5_STEP3!
npm install -g axios express

echo.
echo !F5_STEP4!
npm install -g pnpm yarn

echo.
echo !F5_STEP5!
npm install -g cross-env dotenv nodemon pm2 rimraf mkdirp

echo.
echo ======================================================================
echo !OK! !F5_SUCCESS!
echo !F5_INFO!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 6: Clean All Project Cache ======================
:CLEAN_ALL_CACHE
cls
echo.
echo ======================================================================
echo        !F6_TITLE!
echo ======================================================================
echo.
echo !F6_WARNING!
echo.
pause
echo.

for /d %%d in (!REPO_ROOT!\*) do (
    echo Cleaning: %%~nd
    rd /s /q "%%d\node_modules" 2>nul
    rd /s /q "%%d\.next" 2>nul
    rd /s /q "%%d\dist" 2>nul
    rd /s /q "%%d\build" 2>nul
    rd /s /q "%%d\.nuxt" 2>nul
    del /f /q "%%d\package-lock.json" 2>nul
    del /f /q "%%d\yarn.lock" 2>nul
    del /f /q "%%d\pnpm-lock.yaml" 2>nul
)

echo.
echo ======================================================================
echo !OK! !F6_SUCCESS!
echo ======================================================================
echo.
pause
goto MENU
