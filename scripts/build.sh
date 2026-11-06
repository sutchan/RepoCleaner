#!/usr/bin/env bash
# scripts/build.sh v2.0.1
# Go 交叉编译脚本：产出各平台单文件可执行程序到 bin/。
# 构建期注入 10 位 build 编号（yymmddhhmm），见 docs/CONTRIBUTING.md §8.2。
#
# 工具链要求（Fyne GUI 依赖 OpenGL，需 C 编译器）：
#   - macOS 本机: Xcode (darwin 目标)
#   - linux 目标: gcc + libgl-dev
#   - windows 目标: mingw-w64 (x86_64-w64-mingw32-gcc)
# 无 C 工具链环境可先跑校验构建: go build -tags nogui ./...

set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p bin

# 2 位年份 + 月日时分（24 小时制），同一批次构建共用同一编号
build_id="$(date +%y%m%d%H%M)"

targets=(
  "windows amd64 .exe"
  "windows arm64 .exe"
  "darwin  amd64"
  "darwin  arm64"
  "linux   amd64"
  "linux   arm64"
)

for t in "${targets[@]}"; do
  read -r os arch ext <<<"$t"
  name="repo-cleaner-${os}-${arch}${ext:-}"
  echo "build $name (build $build_id)"
  GOOS="$os" GOARCH="$arch" go build -trimpath -ldflags "-s -w -X repo-cleaner/internal/buildinfo.BuildID=$build_id" -o "bin/$name" ./cmd/repo-cleaner
done

echo "done -> bin/"
