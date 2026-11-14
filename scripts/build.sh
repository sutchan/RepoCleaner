#!/usr/bin/env bash
# scripts/build.sh v2.0.2
# Go 交叉编译脚本：产出各平台单文件可执行程序到 bin/。
# 构建期注入 10 位 build 编号（yymmddhhmm），见 docs/CONTRIBUTING.md §8.2。
#
# Fyne GUI 依赖 cgo / OpenGL：本机目标使用默认 C 编译器；
# 交叉目标需对应交叉编译器，缺失时跳过并给出警告，不静默产出不可运行文件。
# 无 C 工具链环境可先跑校验构建: go build -tags nogui ./...

set -euo pipefail
cd "$(dirname "$0")/.."
mkdir -p bin

# 2 位年份 + 月日时分（24 小时制），同一批次构建共用同一编号
build_id="$(date +%y%m%d%H%M)"
host_os="$(go env GOHOSTOS)"
host_arch="$(go env GOHOSTARCH)"

# os|arch|ext|cross-compiler
targets=(
  "windows|amd64|.exe|x86_64-w64-mingw32-gcc"
  "windows|arm64|.exe|aarch64-w64-mingw32-gcc"
  "darwin|amd64||o64-clang"
  "darwin|arm64||oa64-clang"
  "linux|amd64||x86_64-linux-gnu-gcc"
  "linux|arm64||aarch64-linux-gnu-gcc"
)

built=0
skipped=0
for t in "${targets[@]}"; do
  IFS='|' read -r os arch ext cc <<<"$t"
  if [[ "$os" != "$host_os" || "$arch" != "$host_arch" ]]; then
    if ! command -v "$cc" >/dev/null 2>&1; then
      echo "WARN: skip $os/$arch: cross compiler $cc not found" >&2
      skipped=$((skipped + 1))
      continue
    fi
    export CC="$cc"
  else
    unset CC || true
  fi
  name="repo-manager-${os}-${arch}${ext}"
  echo "build $name (build $build_id)"
  CGO_ENABLED=1 GOOS="$os" GOARCH="$arch" go build -trimpath \
    -ldflags "-s -w -X repo-manager/internal/buildinfo.BuildID=$build_id" \
    -o "bin/$name" ./cmd/repo-manager
  built=$((built + 1))
done

echo "done -> bin/ (built $built, skipped $skipped)"
