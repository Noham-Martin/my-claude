---
description: List all saved sessions
allowed-tools: [Read, Glob, Grep, Bash]
---

# /session-list — List Saved Sessions

Show all saved sessions so you can pick which one to resume.

## Usage
`/session-list`

No arguments. This command is read-only.

## Process

### 1) Find session files

Look for all files in `.planning/sessions/` matching `*.md`.

If `.planning/sessions/` does not exist or is empty: report "No saved sessions found." and stop.

### 2) Parse each session file

For each file, extract from the YAML frontmatter and content:
- **Date** (from filename prefix `YYYY-MM-DD` or frontmatter)
- **Label** (from frontmatter or filename)
- **Branch** (from frontmatter)
- **Status summary** (first bullet from "In Progress" if it exists, otherwise first bullet from "Accomplished")

### 3) Check for active debug sessions

Scan `.planning/debug/` for active (non-resolved) debug files. Note their count.

### 4) Output (plain text)

Sort sessions by date, most recent first.

```
Saved sessions:

  Label              Date        Branch              Status
  ─────              ────        ──────              ──────
  feature-auth       2026-02-06  nohamm/auth         In progress: JWT middleware
  login-bug          2026-02-05  nohamm/fix-login    Resolved: empty password crash
  metrics-v2         2026-02-03  nohamm/metrics      In progress: dashboard queries

  Total: 3 sessions
  Active debug sessions: 1

Resume a session:  /session-resume --feature-auth
```

Keep it plain text, scannable, terminal-friendly. No Markdown.

## Rules

- This command is READ-ONLY. It does not modify any files.
- If a session file cannot be parsed (malformed frontmatter), still list it with "?" for missing fields.
- Always show the resume command hint at the bottom.
