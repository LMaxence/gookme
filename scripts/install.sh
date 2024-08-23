#!/bin/bash

# Define variables
REPO_OWNER="LMaxence"
REPO_NAME="gookme"
BINARY_NAME="gookme"

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names to the expected format
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Determine the file extension for Windows
EXT=""
if [ "$OS" = "mingw64_nt-10.0" ] || [ "$OS" = "msys_nt-10.0" ]; then
    OS="windows"
    EXT=".exe"
fi

# Construct the download URL for the latest release
URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/latest/download/$BINARY_NAME-$OS-$ARCH$EXT"

# Download the binary
echo "Downloading $BINARY_NAME from $URL..."
curl -L -o "$BINARY_NAME$EXT" "$URL"

# Make the binary executable
chmod +x "$BINARY_NAME$EXT"

# Move the binary to a directory in the PATH
mv "$BINARY_NAME$EXT" "/usr/local/bin/$BINARY_NAME$EXT"

gookme --version
echo "Successfully installed Gookme."