#!/bin/bash
# stop-reminder.sh — After Claude stops responding, remind about uncommitted work.
# Informational only — does not block.

set -euo pipefail

# Check if we are in a git repo
if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    exit 0
fi

# Count uncommitted changes
CHANGES=$(git status --porcelain 2>/dev/null | wc -l | tr -d ' ')

if [ "$CHANGES" -gt 0 ]; then
    echo "Reminder: You have $CHANGES uncommitted change(s). Run /verify before committing." >&2
fi

exit 0
