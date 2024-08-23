#!/bin/bash

# Define the root directory and output file
ROOT_DIR="."
OUTPUT_FILE=".github/dependabot.yml"

# Create the output directory if it doesn't exist
mkdir -p "$(dirname "$OUTPUT_FILE")"

# Start the dependabot.yml file
cat <<EOL > "$OUTPUT_FILE"
# Please see the documentation for all configuration options:
# https://docs.github.com/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file

version: 2
updates:
EOL

# Find all directories containing go.work files and append to dependabot.yml
find "$ROOT_DIR" -name "go.mod" -exec dirname {} \; | sort | while read -r dir; do
  cat <<EOL >> "$OUTPUT_FILE"
  - package-ecosystem: "gomod"
    directory: "$dir"
    schedule:
      interval: "weekly"
    commit-message:
      prefix: "deps"
    open-pull-requests-limit: 1
EOL
done

echo "Generated $OUTPUT_FILE for Go workspaces."