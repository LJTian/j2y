#!/bin/bash

# 定义仓库信息
OWNER="LJTian"
REPO="j2y"
BINARY="j2y"

# 获取最新版本号 (通过 GitHub API)
get_latest_release() {
  curl --silent "https://api.github.com/repos/$OWNER/$REPO/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

# 检测操作系统和架构
OS="$(uname -s)"
ARCH="$(uname -m)"

# 映射架构名称以匹配 GoReleaser 的命名规范
case $OS in
  "Linux")
    OS="linux"
    ;;
  "Darwin")
    OS="darwin"
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

case $ARCH in
  "x86_64")
    ARCH="amd64"
    ;;
  "arm64"|"aarch64")
    ARCH="arm64"
    ;;
  *)
    echo "Unsupported Architecture: $ARCH"
    exit 1
    ;;
esac

VERSION=$(get_latest_release)
# 之前的 GoReleaser 配置生成的文件名格式是: j2y_v1.0.2_darwin_arm64.tar.gz
# 注意：这里假设你的 tag 是 v 开头的，比如 v1.0.2
FILE="${BINARY}_${VERSION}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$OWNER/$REPO/releases/download/$VERSION/$FILE"

echo "Detected: $OS $ARCH"
echo "Downloading $BINARY $VERSION..."

# 创建临时目录
TMP_DIR=$(mktemp -d)
curl -sL "$DOWNLOAD_URL" -o "$TMP_DIR/$FILE"

if [ $? -ne 0 ]; then
    echo "Download failed! Please check your network or if the release asset exists."
    exit 1
fi

# 解压
echo "Installing..."
tar -xzf "$TMP_DIR/$FILE" -C "$TMP_DIR"

# 移动到 /usr/local/bin (需要 sudo 权限)
if [ -w "/usr/local/bin" ]; then
    mv "$TMP_DIR/$BINARY" "/usr/local/bin/$BINARY"
else
    echo "Sudo permission required to move binary to /usr/local/bin"
    sudo mv "$TMP_DIR/$BINARY" "/usr/local/bin/$BINARY"
fi

# 清理
rm -rf "$TMP_DIR"

echo "Success! $BINARY installed."
echo "Try running: $BINARY --help"
