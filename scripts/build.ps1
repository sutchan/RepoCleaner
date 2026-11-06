# scripts/build.ps1 v2.0.1
# Go 交叉编译脚本：产出各平台单文件可执行程序到 bin/。
# 构建期注入 10 位 build 编号（yymmddhhmm），见 docs/CONTRIBUTING.md §8.2。
#
# 工具链要求（Fyne GUI 依赖 OpenGL，需 C 编译器）：
#   - Windows 目标: mingw-w64 (gcc)
#   - macOS 目标: 需在 macOS 上用 Xcode 构建（跨平台交叉编译 darwin 需 osxcross）
#   - Linux 目标: gcc + libgl-dev
# 无 C 工具链环境可先跑校验构建: go build -tags nogui ./...

$ErrorActionPreference = "Stop"
$out = Join-Path $PSScriptRoot "..\bin"
New-Item -ItemType Directory -Force -Path $out | Out-Null

# 2 位年份 + 月日时分（24 小时制），同一批次构建共用同一编号
$buildID = (Get-Date).ToString("yyMMddHHmm")

$targets = @(
    @{ os = "windows"; arch = "amd64" },
    @{ os = "windows"; arch = "arm64" },
    @{ os = "darwin";  arch = "amd64" },
    @{ os = "darwin";  arch = "arm64" },
    @{ os = "linux";   arch = "amd64" },
    @{ os = "linux";   arch = "arm64" }
)

foreach ($t in $targets) {
    $env:GOOS = $t.os
    $env:GOARCH = $t.arch
    $ext = if ($t.os -eq "windows") { ".exe" } else { "" }
    $name = "repo-cleaner-$($t.os)-$($t.arch)$ext"
    Write-Host "build $name (build $buildID)"
    go build -trimpath -ldflags "-s -w -X repo-cleaner/internal/buildinfo.BuildID=$buildID" -o (Join-Path $out $name) ./cmd/repo-cleaner
    if ($LASTEXITCODE -ne 0) { exit 1 }
}

Remove-Item Env:GOOS, Env:GOARCH -ErrorAction SilentlyContinue
Write-Host "done -> $out"
