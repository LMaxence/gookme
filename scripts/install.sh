#!/bin/bash

# Define variables
REPO_OWNER="LMaxence"
REPO_NAME="gookme"
BINARY_NAME="gookme"

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
# OS is darwin on MacOS and default to linux for all other OSes
if [ "$OS" == "darwin" ]; then
    echo "Downloading Gookme for MacOS..."
    OS="macos"
else
    echo "Downloading Gookme for Linux..."
    OS="linux"
fi

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

# Construct the download URL for the latest release
URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/latest/download/$BINARY_NAME-$OS-$ARCH"

# Download the binary
echo "Downloading $BINARY_NAME from $URL..."
curl -L -o "$BINARY_NAME" "$URL"

# Make the binary executable
chmod +x "$BINARY_NAME"

# Move the binary to a directory in the PATH
mv "$BINARY_NAME" "/usr/local/bin/$BINARY_NAME"

gookme --version
echo "Successfully installed Gookme."