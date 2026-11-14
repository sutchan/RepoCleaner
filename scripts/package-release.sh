#!/usr/bin/env bash
# scripts/package-release.sh v2.3.3
# 将 build job 产出的二进制打包为 <产品名>-v<版本>-<os>-<arch>.<扩展名>，
# 生成 SHA256SUMS.txt；仅归档文件进入 Release 资产。
set -euo pipefail
shopt -s nullglob

VERSION="${1:?usage: package-release.sh <version> [dist-dir]}"
DIST="${2:-dist}"
PRODUCT="RepoManager"

cd "$DIST"
for bin in repo-manager-*; do
  [ -f "$bin" ] || continue
  base="${bin%.exe}"
  IFS='-' read -r _ _ goos goarch <<< "$base"
  case "$goos" in
    linux) os=linux ;;
    darwin) os=macos ;;
    windows) os=windows ;;
    *) echo "unknown goos: $goos" >&2; exit 1 ;;
  esac
  case "$goarch" in
    amd64) arch=x64 ;;
    arm64) arch=arm64 ;;
    *) echo "unknown goarch: $goarch" >&2; exit 1 ;;
  esac

  archive="$PRODUCT-v$VERSION-$os-$arch"
  archive_path="$(pwd)/$archive"
  staging="$(mktemp -d)"
  cp "$bin" "$staging/$bin"
  cp ../LICENSE "$staging/LICENSE"
  if [ "$os" = "linux" ]; then
    tar -czf "$archive_path.tar.gz" -C "$staging" .
  else
    (cd "$staging" && zip -q -r "$archive_path.zip" .)
  fi
  rm -rf "$staging" "$bin"
  echo "packaged $archive"
done

archives=(*.zip *.tar.gz)
((${#archives[@]} > 0)) || { echo "no release archives produced" >&2; exit 1; }
sha256sum -- "${archives[@]}" > SHA256SUMS.txt
echo "SHA256SUMS.txt:"
cat SHA256SUMS.txt
