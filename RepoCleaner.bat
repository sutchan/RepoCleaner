@echo off
chcp 65001 >nul 2>&1
setlocal enabledelayedexpansion
title RepoCleaner - Github Project Cleaner

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

:INTERACTIVE_INPUT
cls
echo.
echo ======================================================================
echo               Select Repository Root Directory
echo ======================================================================
echo.
echo   [1] Use Current Directory: %CD%
echo   [2] Use Parent Directory: %~dp0..\
echo   [3] Manual Input Path
echo   [4] Use Default: %DEFAULT_REPO_ROOT%
echo.
set /p "PATH_CHOICE=Enter choice [1-4]: "

if "%PATH_CHOICE%"=="1" set "REPO_ROOT=%CD%"
if "%PATH_CHOICE%"=="2" set "REPO_ROOT=%~dp0..\"
if "%PATH_CHOICE%"=="3" (
    set /p "REPO_ROOT=Enter full path: "
)
if "%PATH_CHOICE%"=="4" set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not defined REPO_ROOT set "REPO_ROOT=%DEFAULT_REPO_ROOT%"

if not exist "!REPO_ROOT!" (
    echo.
    echo    ERROR: Path "!REPO_ROOT!" does not exist!
    echo    Please try again...
    echo.
    pause
    goto :INTERACTIVE_INPUT
)

echo.
set /p "SAVE_CONFIG=Save as default? [Y/N]: "
if /i "!SAVE_CONFIG!"=="Y" (
    echo REPO_ROOT=!REPO_ROOT! > "%~dp0config.ini"
    echo.
    echo    Config saved to config.ini
    timeout /t 1 >nul 2>&1
)

:ROOT_SET
if "!REPO_ROOT:~-1!"=="\" set "REPO_ROOT=!REPO_ROOT:~0,-1!"

:: ====================== Main Menu ======================
:MENU
cls
echo.
echo         ======================================================================
echo                        RepoCleaner
echo                   Github Project Cleaner
echo         ======================================================================
echo.
echo            Target: %REPO_ROOT%
echo.
echo         [1] Enable Global node_modules     - No local node_modules
echo         [2] Reset Node Config             - Restore default settings
echo.
echo         [3] Clean .next Cache             - Disable NextJS cache
echo         [4] Restore NextJS Cache          - Enable normal compilation
echo.
echo         [5] Install Global Dependencies   - React/Vue/Next/axios/tools
echo         [6] Clean All Project Cache        - Remove node_modules/.next/dist
echo.
echo         [0] Exit
echo.
echo         ======================================================================
echo.
set /p "CHOICE=Enter choice [0-6]: "

if "%CHOICE%"=="1" goto GLOBAL_NODE
if "%CHOICE%"=="2" goto RESET_NODE
if "%CHOICE%"=="3" goto CLEAN_NEXT
if "%CHOICE%"=="4" goto RESET_NEXT
if "%CHOICE%"=="5" goto INSTALL_DEPS
if "%CHOICE%"=="6" goto CLEAN_ALL_CACHE
if "%CHOICE%"=="0" exit
goto MENU

:: ====================== Function 1: Enable Global node_modules ======================
:GLOBAL_NODE
cls
echo.
echo [1/4] Getting global node_modules path...
for /f "delims=" %%i in ('npm root -g') do set "GLOBAL_NODE_MODULES=%%i"
echo Path: !GLOBAL_NODE_MODULES!

echo.
echo [2/4] Setting NODE_PATH environment variable...
setx NODE_PATH "!GLOBAL_NODE_MODULES!" >nul 2>&1

echo.
echo [3/4] Disabling local node_modules for all projects...
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
echo [OK] Global node_modules Enabled!
echo [INFO] Run [5] to install dependencies first time
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 2: Reset Node Config ======================
:RESET_NODE
cls
echo.
echo Removing NODE_PATH environment variable...
setx NODE_PATH "" >nul 2>&1
reg delete "HKCU\Environment" /v NODE_PATH /f >nul 2>&1

echo.
echo Cleaning project .npmrc configs...
for /d %%d in (!REPO_ROOT!\*) do (
    del /f /q "%%d\.npmrc" >nul 2>&1
)

echo.
echo ======================================================================
echo [OK] Node Config Reset!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 3: Clean .next Cache ======================
:CLEAN_NEXT
cls
echo.
echo Scanning and cleaning NextJS cache...
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
echo [OK] .next Cache Cleaned and Disabled!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 4: Restore .next ======================
:RESET_NEXT
cls
echo.
echo Restoring NextJS cache functionality...
for /d %%d in (!REPO_ROOT!\*) do (
    if exist "%%d\next.config.js" (
        echo Restoring: %%~nd
        del /f /q /a:s /a:h /a:r "%%d\.next" 2>nul
        rd /s /q "%%d\.next" 2>nul
    )
)

echo.
echo ======================================================================
echo [OK] NextJS Cache Restored!
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 5: Install Global Dependencies ======================
:INSTALL_DEPS
cls
echo.
echo ======================================================================
echo        Installing Common Global Dependencies
echo ======================================================================
echo.
echo Install List: React, Vue, Next, Axios, Express, pnpm, yarn, rimraf, etc.
echo Please wait...
echo.

echo.
echo [1/5] Installing React ecosystem...
npm install -g react react-dom next

echo.
echo [2/5] Installing Vue ecosystem...
npm install -g vue @vue/cli

echo.
echo [3/5] Installing HTTP libraries...
npm install -g axios express

echo.
echo [4/5] Installing package managers...
npm install -g pnpm yarn

echo.
echo [5/5] Installing dev tools...
npm install -g cross-env dotenv nodemon pm2 rimraf mkdirp

echo.
echo ======================================================================
echo [OK] All Dependencies Installed!
echo [OK] Projects can use them directly without local install
echo ======================================================================
echo.
pause
goto MENU

:: ====================== Function 6: Clean All Project Cache ======================
:CLEAN_ALL_CACHE
cls
echo.
echo ======================================================================
echo        Cleaning All Project Cache
echo ======================================================================
echo.
echo WARNING: This will permanently delete all cache files!
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
echo [OK] All Cache Cleaned!
echo ======================================================================
echo.
pause
goto MENU
