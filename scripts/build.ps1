# scripts/build.ps1 v2.0.2
# Go 交叉编译脚本：产出各平台单文件可执行程序到 bin/。
# 构建期注入 10 位 build 编号（yymmddhhmm），见 docs/CONTRIBUTING.md §8.2。
#
# Fyne GUI 依赖 cgo / OpenGL：本机目标使用默认 C 编译器；
# 交叉目标需对应交叉编译器，缺失时跳过并给出警告，不静默产出不可运行文件。
# 无 C 工具链环境可先跑校验构建: go build -tags nogui ./...

$ErrorActionPreference = "Stop"
$out = Join-Path $PSScriptRoot "..\bin"
New-Item -ItemType Directory -Force -Path $out | Out-Null

# 2 位年份 + 月日时分（24 小时制），同一批次构建共用同一编号
$buildID = (Get-Date).ToString("yyMMddHHmm")
$hostOS = (go env GOHOSTOS).Trim()
$hostArch = (go env GOHOSTARCH).Trim()

$targets = @(
    @{ os = "windows"; arch = "amd64"; cc = "x86_64-w64-mingw32-gcc" },
    @{ os = "windows"; arch = "arm64"; cc = "aarch64-w64-mingw32-gcc" },
    @{ os = "darwin";  arch = "amd64"; cc = "o64-clang" },
    @{ os = "darwin";  arch = "arm64"; cc = "oa64-clang" },
    @{ os = "linux";   arch = "amd64"; cc = "x86_64-linux-gnu-gcc" },
    @{ os = "linux";   arch = "arm64"; cc = "aarch64-linux-gnu-gcc" }
)

$built = 0
$skipped = 0
foreach ($t in $targets) {
    $native = ($t.os -eq $hostOS -and $t.arch -eq $hostArch)
    if (-not $native -and -not (Get-Command $t.cc -ErrorAction SilentlyContinue)) {
        Write-Warning "跳过 $($t.os)/$($t.arch)：缺少交叉编译器 $($t.cc)"
        $skipped++
        continue
    }
    $env:GOOS = $t.os
    $env:GOARCH = $t.arch
    $env:CGO_ENABLED = "1"
    if ($native) {
        Remove-Item Env:CC -ErrorAction SilentlyContinue
    } else {
        $env:CC = $t.cc
    }
    $ext = if ($t.os -eq "windows") { ".exe" } else { "" }
    $name = "repo-cleaner-$($t.os)-$($t.arch)$ext"
    Write-Host "build $name (build $buildID)"
    go build -trimpath -ldflags "-s -w -X repo-cleaner/internal/buildinfo.BuildID=$buildID" -o (Join-Path $out $name) ./cmd/repo-cleaner
    if ($LASTEXITCODE -ne 0) { exit 1 }
    $built++
}

Remove-Item Env:GOOS, Env:GOARCH, Env:CC -ErrorAction SilentlyContinue
Write-Host "done -> $out (built $built, skipped $skipped)"
