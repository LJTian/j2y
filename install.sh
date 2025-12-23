#!/bin/bash

OWNER="LJTian"
REPO="j2y"
BINARY="j2y"

get_latest_release() {
  curl --silent "https://api.github.com/repos/$OWNER/$REPO/releases/latest" |
    grep '"tag_name":' |
    sed -E 's/.*"([^"]+)".*/\1/'
}

OS="$(uname -s)"
ARCH="$(uname -m)"

case $OS in
  "Linux") OS="linux" ;;
  "Darwin") OS="darwin" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case $ARCH in
  "x86_64") ARCH="amd64" ;;
  "arm64"|"aarch64") ARCH="arm64" ;;
  *) echo "Unsupported Architecture: $ARCH"; exit 1 ;;
esac

# 获取 Tag (例如 v1.0.1-rc1)
TAG=$(get_latest_release)

# 👇 关键修改：去掉开头的 'v' 用于拼接文件名 (变成 1.0.1-rc1)
VERSION=${TAG#v}

# 拼接文件名 (GoReleaser 生成的文件不带 v)
FILE="${BINARY}_${VERSION}_${OS}_${ARCH}.tar.gz"

# 拼接下载链接 (URL 路径里需要带 v 的 Tag)
DOWNLOAD_URL="https://github.com/$OWNER/$REPO/releases/download/$TAG/$FILE"

echo "Detected: $OS $ARCH"
echo "Downloading $BINARY $VERSION from $TAG..."
echo "URL: $DOWNLOAD_URL" # 打印出来方便调试

TMP_DIR=$(mktemp -d)
# 增加 -f 参数，如果 404 直接报错而不是下载错误页面
curl -sfL "$DOWNLOAD_URL" -o "$TMP_DIR/$FILE"

if [ $? -ne 0 ]; then
    echo "❌ Download failed! File not found on GitHub."
    echo "Expected file: $FILE"
    exit 1
fi

echo "Installing..."
tar -xzf "$TMP_DIR/$FILE" -C "$TMP_DIR"

if [ -w "/usr/local/bin" ]; then
    mv "$TMP_DIR/$BINARY" "/usr/local/bin/$BINARY"
else
    echo "sudo permission required..."
    sudo mv "$TMP_DIR/$BINARY" "/usr/local/bin/$BINARY"
fi

rm -rf "$TMP_DIR"

echo "✅ Success! $BINARY installed."
$BINARY --help
