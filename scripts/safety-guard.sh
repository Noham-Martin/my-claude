#!/bin/bash
# safety-guard.sh — Block dangerous bash commands before execution.
# Exit code 2 = block the tool use. Exit code 0 = allow.
#
# Receives JSON on stdin with the tool input.

set -euo pipefail

# Read the tool input from stdin
INPUT=$(cat)

# Extract the command being run
COMMAND=$(echo "$INPUT" | python3 -c "
import sys, json
try:
    data = json.load(sys.stdin)
    # The command is in tool_input.command for Bash tool
    print(data.get('tool_input', {}).get('command', ''))
except:
    print('')
" 2>/dev/null || echo "")

if [ -z "$COMMAND" ]; then
    exit 0
fi

# --- BLOCK: Destructive commands ---

# rm -rf with broad targets
if echo "$COMMAND" | grep -qE 'rm\s+(-[a-zA-Z]*r[a-zA-Z]*f|--recursive\s+--force|-[a-zA-Z]*f[a-zA-Z]*r)\s+(/|~|\.\.|\./)'; then
    echo "BLOCKED: 'rm -rf' targeting root, home, or parent directories is not allowed." >&2
    exit 2
fi

# git push --force to main/master
if echo "$COMMAND" | grep -qE 'git\s+push\s+.*--force.*\s+(main|master)'; then
    echo "BLOCKED: Force-pushing to main/master is not allowed." >&2
    exit 2
fi
if echo "$COMMAND" | grep -qE 'git\s+push\s+.*-f\s+.*\s+(main|master)'; then
    echo "BLOCKED: Force-pushing to main/master is not allowed." >&2
    exit 2
fi

# git reset --hard
if echo "$COMMAND" | grep -qE 'git\s+reset\s+--hard'; then
    echo "WARNING: 'git reset --hard' discards uncommitted changes. Ensure this is intentional." >&2
    # Allow but warn — user can still deny via permission prompt
    exit 0
fi

# git checkout . / git restore .
if echo "$COMMAND" | grep -qE 'git\s+(checkout|restore)\s+\.$'; then
    echo "WARNING: This discards all uncommitted changes." >&2
    exit 0
fi

# git clean -f
if echo "$COMMAND" | grep -qE 'git\s+clean\s+-[a-zA-Z]*f'; then
    echo "WARNING: 'git clean -f' removes untracked files permanently." >&2
    exit 0
fi

# DROP TABLE / TRUNCATE
if echo "$COMMAND" | grep -qiE '(drop\s+table|truncate\s+table)'; then
    echo "BLOCKED: Destructive database operations are not allowed." >&2
    exit 2
fi

# All clear
exit 0
