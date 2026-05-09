@echo off
chcp 65001 >nul 2>&1
setlocal enabledelayedexpansion
title RepoCleaner - 项目清理工具

:: ========================================================================
:: 配置和初始化
:: ========================================================================
set "CONFIG_FILE=%~dp0config.ini"
set "LANG_FILE=%~dp0lang.ini"
set "DEFAULT_REPO_ROOT=E:\Github"

:: ========================================================================
:: 加载配置
:: ========================================================================
call :LOAD_CONFIG

:: ========================================================================
:: 加载语言
:: ========================================================================
call :LOAD_LANGUAGE

:: ========================================================================
:: 解析命令行参数
:: ========================================================================
if not "%~1"=="" call :PARSE_ARGUMENTS %*

:: ========================================================================
:: 设置根目录
:: ========================================================================
call :SET_REPO_ROOT

:: ========================================================================
:: 主循环
:: ========================================================================
:MENU
call :DISPLAY_MENU
call :GET_USER_CHOICE
call :PROCESS_CHOICE
goto :MENU

:: ========================================================================
:: 配置管理模块
:: ========================================================================
:LOAD_CONFIG
set "LANG_SECTION=zh"
set "REPO_ROOT="

if exist "%CONFIG_FILE%" (
    for /f "usebackq tokens=1,2 delims==" %%a in ("%CONFIG_FILE%") do (
        if /i "%%a"=="REPO_ROOT" set "REPO_ROOT=%%b"
        if /i "%%a"=="LANG" set "LANG_SECTION=%%b"
    )
)
goto :EOF

:SAVE_CONFIG
set "KEY=%~1"
set "VALUE=%~2"

if not defined VALUE goto :EOF

if /i "%KEY%"=="REPO_ROOT" (
    (
        echo REPO_ROOT=%VALUE%
        if exist "%CONFIG_FILE%" (
            for /f "usebackq tokens=1,* delims==" %%a in ("%CONFIG_FILE%") do (
                if /i "%%a" neq "REPO_ROOT" echo %%a=%%b
            )
        )
    ) > "%CONFIG_FILE%"
) else if /i "%KEY%"=="LANG" (
    (
        echo LANG=%VALUE%
        if exist "%CONFIG_FILE%" (
            for /f "usebackq tokens=1,* delims==" %%a in ("%CONFIG_FILE%") do (
                if /i "%%a" neq "LANG" echo %%a=%%b
            )
        )
    ) > "%CONFIG_FILE%"
)
goto :EOF

:SET_REPO_ROOT
if defined REPO_ROOT (
    if exist "!REPO_ROOT!" (
        if "!REPO_ROOT:~-1!"=="\" set "REPO_ROOT=!REPO_ROOT:~0,-1!"
        goto :EOF
    )
)

:INTERACTIVE_PATH
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
set /p "CHOICE=!PATH_PROMPT!"

if "!CHOICE!"=="1" set "REPO_ROOT=%CD%"
if "!CHOICE!"=="2" set "REPO_ROOT=%~dp0..\"
if "!CHOICE!"=="3" set /p "REPO_ROOT=!PATH_INPUT!"
if "!CHOICE!"=="4" set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not defined REPO_ROOT set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not exist "!REPO_ROOT!" (
    echo.
    echo    !PATH_ERROR!
    echo.
    pause
    goto :INTERACTIVE_PATH
)

if "!REPO_ROOT:~-1!"=="\" set "REPO_ROOT=!REPO_ROOT:~0,-1!"

echo.
set /p "SAVE=!PATH_SAVE!"
if /i "!SAVE!"=="Y" (
    call :SAVE_CONFIG REPO_ROOT "!REPO_ROOT!"
    echo    !PATH_SAVED!
    timeout /t 1 >nul 2>&1
)
goto :EOF

:: ========================================================================
:: 语言管理模块
:: ========================================================================
:LOAD_LANGUAGE
if "!LANG_SECTION!"=="en" (
    call :SET_LANG_EN
) else (
    call :SET_LANG_ZH
)
goto :EOF

:SET_LANG_ZH
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
set "PATH_INPUT=请输入完整路径: "

set "MENU_TITLE=RepoCleaner - 项目清理工具"
set "MENU_TARGET=目标目录:"
set "MENU_OPTION1=启用全局 node_modules     - 无需本地 node_modules"
set "MENU_OPTION2=重置 Node 配置            - 恢复默认设置"
set "MENU_OPTION3=清理 .next 缓存           - 禁用 NextJS 缓存"
set "MENU_OPTION4=恢复 NextJS 缓存          - 启用正常编译"
set "MENU_OPTION5=安装全局依赖              - React/Vue/Next/axios/工具"
set "MENU_OPTION6=清理所有项目缓存          - 删除 node_modules/.next/dist"
set "MENU_OPTION7=清理 Vite/构建缓存        - 删除 .vite/dist/.turbo"
set "MENU_EXIT=退出"
set "MENU_CHOOSE=请输入选项 [0-7]:"

set "F1_TITLE=启用全局 node_modules"
set "F1_STEP1=[1/4] 获取全局 node_modules 路径..."
set "F1_STEP2=[2/4] 设置 NODE_PATH 环境变量..."
set "F1_STEP3=[3/4] 禁用所有项目的本地 node_modules..."
set "F1_STEP4=[4/4] 清理现有本地 node_modules 目录..."
set "F1_SUCCESS=全局 node_modules 已启用!"
set "F1_INFO=请先运行 [5] 安装依赖"

set "F2_TITLE=重置 Node 配置"
set "F2_STEP1=[1/2] 删除 NODE_PATH 环境变量..."
set "F2_STEP2=[2/2] 清理项目 .npmrc 配置..."
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

set "F7_TITLE=清理 Vite/构建缓存"
set "F7_WARNING=警告: 这将删除 .vite/dist/.turbo 目录!"
set "F7_SUCCESS=Vite/构建缓存已清理!"

set "OK=[完成]"
set "ERROR=[错误]"
set "PAUSE=按任意键继续..."

set "HELP_TITLE=RepoCleaner - 项目清理工具"
set "HELP_USAGE=用法: RepoCleaner.bat [选项] [路径]"
set "HELP_OPTIONS=选项:"
set "HELP_EXAMPLES=示例:"
set "HELP_OPT_FUNC=  1-7       直接运行对应的菜单选项"
set "HELP_OPT_LANG=  -lang en  设置为英文"
set "HELP_OPT_LANG2=  -lang zh  设置为中文"
set "HELP_OPT_HELP=  -h        显示帮助信息"
set "HELP_EX1=  RepoCleaner.bat                    # 交互模式"
set "HELP_EX2=  RepoCleaner.bat E:\Github          # 设置项目根目录"
set "HELP_EX3=  RepoCleaner.bat 1                  # 启用全局 node_modules"
set "HELP_EX4=  RepoCleaner.bat -lang en           # 设置英文"
set "HELP_EX5=  RepoCleaner.bat 5 E:\Projects      # 为指定路径安装依赖"

goto :EOF

:SET_LANG_EN
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
set "PATH_INPUT=Enter full path: "

set "MENU_TITLE=RepoCleaner - Github Project Cleaner"
set "MENU_TARGET=Target:"
set "MENU_OPTION1=Enable Global node_modules     - No local node_modules"
set "MENU_OPTION2=Reset Node Config             - Restore default settings"
set "MENU_OPTION3=Clean .next Cache             - Disable NextJS cache"
set "MENU_OPTION4=Restore NextJS Cache          - Enable normal compilation"
set "MENU_OPTION5=Install Global Dependencies   - React/Vue/Next/axios/tools"
set "MENU_OPTION6=Clean All Project Cache       - Remove node_modules/.next/dist"
set "MENU_OPTION7=Clean Vite/Build Cache        - Remove .vite/dist/.turbo"
set "MENU_EXIT=Exit"
set "MENU_CHOOSE=Enter choice [0-7]:"

set "F1_TITLE=Enable Global node_modules"
set "F1_STEP1=[1/4] Getting global node_modules path..."
set "F1_STEP2=[2/4] Setting NODE_PATH environment variable..."
set "F1_STEP3=[3/4] Disabling local node_modules for all projects..."
set "F1_STEP4=[4/4] Cleaning existing local node_modules directories..."
set "F1_SUCCESS=Global node_modules Enabled!"
set "F1_INFO=Run [5] to install dependencies first time"

set "F2_TITLE=Reset Node Config"
set "F2_STEP1=[1/2] Removing NODE_PATH environment variable..."
set "F2_STEP2=[2/2] Cleaning project .npmrc configs..."
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

set "F7_TITLE=Cleaning Vite/Build Cache"
set "F7_WARNING=WARNING: This will delete .vite/dist/.turbo directories!"
set "F7_SUCCESS=Vite/Build Cache Cleaned!"

set "OK=[OK]"
set "ERROR=[ERROR]"
set "PAUSE=Press any key to continue..."

set "HELP_TITLE=RepoCleaner - Project Cleaner"
set "HELP_USAGE=Usage: RepoCleaner.bat [options] [path]"
set "HELP_OPTIONS=Options:"
set "HELP_EXAMPLES=Examples:"
set "HELP_OPT_FUNC=  1-7       Run corresponding menu option directly"
set "HELP_OPT_LANG=  -lang en  Set language to English"
set "HELP_OPT_LANG2=  -lang zh  Set language to Chinese"
set "HELP_OPT_HELP=  -h        Show this help"
set "HELP_EX1=  RepoCleaner.bat                    # Interactive mode"
set "HELP_EX2=  RepoCleaner.bat E:\Github          # Set project root"
set "HELP_EX3=  RepoCleaner.bat 1                  # Enable global node_modules"
set "HELP_EX4=  RepoCleaner.bat -lang en           # Set English"
set "HELP_EX5=  RepoCleaner.bat 5 E:\Projects      # Install deps for path"

goto :EOF

:: ========================================================================
:: 参数解析模块
:: ========================================================================
:PARSE_ARGUMENTS
:PARSE_LOOP
if "%~1"=="" goto :EOF

set "ARG=%~1"

if /i "%ARG%"=="-h" goto :SHOW_HELP
if /i "%ARG%"=="--help" goto :SHOW_HELP

if /i "%ARG%"=="-lang" (
    set "LANG_SECTION=%~2"
    call :SAVE_CONFIG LANG "%~2"
    shift
    shift
    goto :PARSE_LOOP
)

if /i "%ARG%"=="1" set "ARG_MODE=1"
if /i "%ARG%"=="2" set "ARG_MODE=2"
if /i "%ARG%"=="3" set "ARG_MODE=3"
if /i "%ARG%"=="4" set "ARG_MODE=4"
if /i "%ARG%"=="5" set "ARG_MODE=5"
if /i "%ARG%"=="6" set "ARG_MODE=6"
if /i "%ARG%"=="7" set "ARG_MODE=7"

if defined ARG_MODE (
    if not "%~2"=="" (
        if exist "%~2" set "REPO_ROOT=%~2"
    )
    goto :SET_ROOT_DIRECT
)

if exist "%ARG%" (
    set "REPO_ROOT=%ARG%"
    shift
    goto :PARSE_LOOP
)

shift
goto :PARSE_LOOP

:SHOW_HELP
echo.
echo !HELP_TITLE!
echo.
echo !HELP_USAGE!
echo.
echo !HELP_OPTIONS!
echo !HELP_OPT_FUNC!
echo !HELP_OPT_LANG!
echo !HELP_OPT_LANG2!
echo !HELP_OPT_HELP!
echo.
echo !HELP_EXAMPLES!
echo !HELP_EX1!
echo !HELP_EX2!
echo !HELP_EX3!
echo !HELP_EX4!
echo !HELP_EX5!
echo.
exit /b 0

:SET_ROOT_DIRECT
if exist "!REPO_ROOT!" (
    if "!REPO_ROOT:~-1!"=="\" set "REPO_ROOT=!REPO_ROOT:~0,-1!"
)
goto :RUN_ARG_FUNCTION

:RUN_ARG_FUNCTION
if "!ARG_MODE!"=="1" goto :FUNC_GLOBAL_NODE
if "!ARG_MODE!"=="2" goto :FUNC_RESET_NODE
if "!ARG_MODE!"=="3" goto :FUNC_CLEAN_NEXT
if "!ARG_MODE!"=="4" goto :FUNC_RESTORE_NEXT
if "!ARG_MODE!"=="5" goto :FUNC_INSTALL_DEPS
if "!ARG_MODE!"=="6" goto :FUNC_CLEAN_ALL
if "!ARG_MODE!"=="7" goto :FUNC_CLEAN_VITE
goto :EOF

:: ========================================================================
:: UI显示模块
:: ========================================================================
:DISPLAY_MENU
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
echo         [7] !MENU_OPTION7!
echo.
echo         [L] !LANG_1! / !LANG_2!
echo         [0] !MENU_EXIT!
echo.
echo         ======================================================================
echo.
goto :EOF

:GET_USER_CHOICE
set /p "CHOICE=!MENU_CHOOSE!"
goto :EOF

:PROCESS_CHOICE
if "%CHOICE%"=="1" goto :FUNC_GLOBAL_NODE
if "%CHOICE%"=="2" goto :FUNC_RESET_NODE
if "%CHOICE%"=="3" goto :FUNC_CLEAN_NEXT
if "%CHOICE%"=="4" goto :FUNC_RESTORE_NEXT
if "%CHOICE%"=="5" goto :FUNC_INSTALL_DEPS
if "%CHOICE%"=="6" goto :FUNC_CLEAN_ALL
if "%CHOICE%"=="7" goto :FUNC_CLEAN_VITE
if /i "%CHOICE%"=="L" (
    call :SHOW_LANG_MENU
    goto :EOF
)
if "%CHOICE%"=="0" exit
goto :EOF

:SHOW_LANG_MENU
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

if "!LANG_CHOICE!"=="1" set "LANG_SECTION=en"
if "!LANG_CHOICE!"=="2" set "LANG_SECTION=zh"

call :LOAD_LANGUAGE
goto :EOF

:: ========================================================================
:: 功能实现模块
:: ========================================================================
:FUNC_GLOBAL_NODE
cls
echo.
echo !F1_STEP1!
for /f "delims=" %%i in ('npm root -g') do set "GLOBAL_NODE_MODULES=%%i"
echo Path: !GLOBAL_NODE_MODULES!

echo.
echo !F1_STEP2!
setx NODE_PATH "!GLOBAL_NODE_MODULES!" >nul 2>&1
set "NODE_PATH=!GLOBAL_NODE_MODULES!"

echo.
echo !F1_STEP3!
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\package.json" (
        echo Configured: %%~nd
        (
            echo prefix=!GLOBAL_NODE_MODULES!
            echo cache=!GLOBAL_NODE_MODULES!\.npm-cache
            echo tmp="!TEMP!"
            echo global=true
            echo prefer-global=true
        ) > "%%d\.npmrc"
    )
)

echo.
echo !F1_STEP4!
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\package.json" (
        if exist "%%d\node_modules" (
            echo Cleaning: %%~nd
            rd /s /q "%%d\node_modules" 2>nul
        )
    )
)

echo.
echo ======================================================================
echo !OK! !F1_SUCCESS!
echo !F1_INFO!
echo.
echo Global Path: !GLOBAL_NODE_MODULES!
echo NODE_PATH: !NODE_PATH!
echo ======================================================================
echo.
pause
goto :EOF

:FUNC_RESET_NODE
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
goto :EOF

:FUNC_CLEAN_NEXT
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
goto :EOF

:FUNC_RESTORE_NEXT
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
goto :EOF

:FUNC_INSTALL_DEPS
cls
echo.
echo ======================================================================
echo        !F5_TITLE!
echo ======================================================================
echo.
echo !F5_LIST!
echo !F5_WAIT!
echo.

call :EXEC_NPM_INSTALL "react react-dom next" "!F5_STEP1!"
call :EXEC_NPM_INSTALL "vue @vue/cli" "!F5_STEP2!"
call :EXEC_NPM_INSTALL "axios express" "!F5_STEP3!"
call :EXEC_NPM_INSTALL "pnpm yarn" "!F5_STEP4!"
call :EXEC_NPM_INSTALL "cross-env dotenv nodemon pm2 rimraf mkdirp" "!F5_STEP5!"

echo.
echo ======================================================================
echo !OK! !F5_SUCCESS!
echo !F5_INFO!
echo ======================================================================
echo.
pause
goto :EOF

:EXEC_NPM_INSTALL
set "PKGS=%~1"
set "STEP_MSG=%~2"
echo.
echo !STEP_MSG!
npm install -g !PKGS!
goto :EOF

:FUNC_CLEAN_ALL
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
    rd /s /q "%%d\.nuxt" 2>nul
    rd /s /q "%%d\dist" 2>nul
    rd /s /q "%%d\build" 2>nul
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
goto :EOF

:FUNC_CLEAN_VITE
cls
echo.
echo ======================================================================
echo        !F7_TITLE!
echo ======================================================================
echo.
echo !F7_WARNING!
echo.
pause
echo.

for /d %%d in (!REPO_ROOT!\*) do (
    echo Cleaning: %%~nd
    rd /s /q "%%d\.vite" 2>nul
    rd /s /q "%%d\dist" 2>nul
    rd /s /q "%%d\.turbo" 2>nul
    rd /s /q "%%d\coverage" 2>nul
    rd /s /q "%%d\.nyc_output" 2>nul
)

echo.
echo ======================================================================
echo !OK! !F7_SUCCESS!
echo ======================================================================
echo.
pause
goto :EOF

:: ========================================================================
:: 通用显示模块 (已弃用，保留以兼容)
:: ========================================================================
:: SHOW_SUCCESS 函数已弃用，每个功能现在直接显示自己的成功消息
