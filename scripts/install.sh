#!/bin/bash

# Define variables
REPO_OWNER="LMaxence"
REPO_NAME="gookme"
BINARY_NAME="gookme"

# Determine the OS and architecture
OS=$(uname | tr '[:upper:]' '[:lower:]')
# OS is darwin on MacOS and defaults to linux for all other OSes
if [ "$OS" == "darwin" ]; then
    echo "Downloading Gookme for MacOS..."
    OS="darwin"
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

# Create the .local/bin directory if it doesn't exist
LOCAL_BIN="$HOME/.local/bin"
if [ ! -d "$LOCAL_BIN" ]; then
    echo "Creating $LOCAL_BIN directory..."
    mkdir -p "$LOCAL_BIN"
fi

# Move the binary to the .local/bin directory
mv "$BINARY_NAME" "$LOCAL_BIN/$BINARY_NAME"

# Verify the installation
"$LOCAL_BIN/$BINARY_NAME" --version
echo "Successfully installed Gookme."