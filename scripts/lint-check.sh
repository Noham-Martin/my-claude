#!/bin/bash
# lint-check.sh — Run linter on edited files (async, post-edit).
# This is informational only — does not block.

set -euo pipefail

# Read tool input from stdin
INPUT=$(cat)

# Extract the file path that was edited
FILE_PATH=$(echo "$INPUT" | python3 -c "
import sys, json
try:
    data = json.load(sys.stdin)
    print(data.get('tool_input', {}).get('file_path', ''))
except:
    print('')
" 2>/dev/null || echo "")

if [ -z "$FILE_PATH" ] || [ ! -f "$FILE_PATH" ]; then
    exit 0
fi

# Determine file extension
EXT="${FILE_PATH##*.}"

case "$EXT" in
    ts|tsx|js|jsx)
        # Check if eslint is available
        if command -v npx >/dev/null 2>&1 && [ -f "$(git rev-parse --show-toplevel 2>/dev/null)/.eslintrc"* ] 2>/dev/null; then
            npx eslint --quiet "$FILE_PATH" 2>/dev/null || true
        fi
        ;;
    py)
        # Check if ruff is available, fallback to flake8
        if command -v ruff >/dev/null 2>&1; then
            ruff check "$FILE_PATH" 2>/dev/null || true
        elif command -v flake8 >/dev/null 2>&1; then
            flake8 "$FILE_PATH" 2>/dev/null || true
        fi
        ;;
    go)
        if command -v golangci-lint >/dev/null 2>&1; then
            golangci-lint run "$FILE_PATH" 2>/dev/null || true
        fi
        ;;
esac

exit 0
