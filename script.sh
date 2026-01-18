#!/bin/sh
set -e # Exit immediately if a command exits with a non-zero status

REPO="Bharath-code/git-scope"
INSTALL_DIR="${HOME}/.local/bin"
BINARY="git-scope"

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

log() { printf "${GREEN}> %s${NC}\n" "$1"; }
err() { printf "${RED}ERROR: %s${NC}\n" "$1" >&2; exit 1; }

# 1. Validation
command -v curl >/dev/null || err "curl is required."
command -v tar >/dev/null  || err "tar is required."

# 2. Detect System
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64|amd64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) err "Unsupported architecture: $ARCH" ;;
esac

case "$OS" in
    linux|darwin) ;;
    mingw*|msys*|cygwin*) OS="windows" ;;
    *) err "Unsupported OS: $OS" ;;
esac

# 3. Fetch Latest Version
log "Checking latest version..."
VERSION_JSON=$(curl -sL "https://api.github.com/repos/${REPO}/releases/latest")
VERSION=$(echo "$VERSION_JSON" | grep -o '"tag_name": *"[^"]*"' | sed 's/.*"tag_name": *"//;s/"//')

if [ -z "$VERSION" ]; then
    err "Could not find latest release. Check internet or GitHub API limits."
fi

# 4. Prepare Download
# Assumes asset format: git-scope_1.0.0_linux_amd64.tar.gz
CLEAN_VER="${VERSION#v}"
ASSET_NAME="${BINARY}_${CLEAN_VER}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${ASSET_NAME}"

TMP_DIR=$(mktemp -d)
trap 'rm -rf "$TMP_DIR"' EXIT INT TERM

# 5. Download & Install
log "Downloading ${VERSION}..."
if ! curl -sSL -f -o "$TMP_DIR/release.tar.gz" "$URL"; then
    err "Download failed. Asset $ASSET_NAME may not exist."
fi

tar -xzf "$TMP_DIR/release.tar.gz" -C "$TMP_DIR"

# Find binary (handles potential subfolders or extensions like .exe)
FOUND_BIN=$(find "$TMP_DIR" -type f -name "${BINARY}*" | head -n 1)
[ -f "$FOUND_BIN" ] || err "Binary not found in archive."

mkdir -p "$INSTALL_DIR"
mv "$FOUND_BIN" "$INSTALL_DIR/$BINARY"
chmod +x "$INSTALL_DIR/$BINARY"

# 6. Final Check
if ! command -v "$BINARY" >/dev/null && ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    log "Installed successfully to $INSTALL_DIR"
    printf "  ! Warning: $INSTALL_DIR is not in your PATH.\n"
    printf "  ! Add this to your shell config: export PATH=\"\$HOME/.local/bin:\$PATH\"\n"
else
    log "Success! Run '$BINARY' to start."
fi
