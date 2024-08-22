#!/bin/sh

# Check Go formatting
unformatted=$(gofmt -d -e -l .)
if [ -n "$unformatted" ]; then
  echo "$unformatted"
  exit 1
else
  echo "All files are properly formatted."
fi