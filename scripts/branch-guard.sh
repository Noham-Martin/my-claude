#!/bin/bash
# branch-guard.sh — Prevent file edits when on main/master branch.
# Exit code 2 = block. Exit code 0 = allow.

set -euo pipefail

# Check if we are in a git repository
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    # Not a git repo — allow edits
    exit 0
fi

BRANCH=$(git branch --show-current 2>/dev/null || echo "")

if [ "$BRANCH" = "main" ] || [ "$BRANCH" = "master" ]; then
    echo "BLOCKED: You are on '$BRANCH'. Switch to a feature branch before editing files." >&2
    echo "  Create one with: git checkout -b nohamm/<branch-name>" >&2
    exit 2
fi

exit 0
