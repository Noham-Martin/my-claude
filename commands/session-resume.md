---
description: Restore context from a previous session
argument-hint: [--<label>]
allowed-tools: [Read, Glob, Grep, Bash]
---

# /session-resume — Resume Previous Session

Restore full context from a previous session so you can pick up where you left off.

## Usage
- `/session-resume` — resume the most recent session
- `/session-resume --<label>` — resume a specific session by label

## Process

### 1) Load STATE.md

Check for `.planning/STATE.md` in the current directory.
- If it exists: read it for quick context overview.
- If not: note that no state file exists.

### 2) Find session file

- If `--<label>` is provided: look for `.planning/sessions/*-<label>.md`
- If no label: find the most recent file in `.planning/sessions/` by date prefix.
- If no session files exist: report "No saved sessions found. Nothing to resume."

Read the full session file.

### 3) Verify git state

Compare the session file's recorded git state with current reality:
- `git branch --show-current` — is it the same branch?
- `git status --short` — any unexpected changes?
- `git log --oneline -5` — do recent commits match?

If the git state has diverged, warn the user:
```
WARNING: Git state has changed since the session was saved.
  Expected branch: <saved>
  Current branch: <actual>
  New commits since save: <count>
```

Do NOT modify git state. Just report discrepancies.

### 4) Output context restore

Present a structured summary (Markdown):

```markdown
## Session Restored: <label>

**Saved:** <date>
**Branch:** <branch>

### What was accomplished
- <from session file>

### What is in progress
- <from session file, with file paths>

### Key decisions
- <from session file>

### Blockers
- <from session file>

### Mental context
- <from session file>

### Git state
- <current state + any discrepancies>
```

### 5) Suggest next actions

Based on the "Next Steps" and "In Progress" sections, suggest what to do:

```
Suggested next actions:
1. <most important next step>
2. <second step>
3. <third step>
```

If there are active debug sessions in `.planning/debug/`, mention them:
```
Note: Active debug session found — run /debug --resume to continue.
```

## Rules

- This command is READ-ONLY. It does not modify any files or git state.
- If context seems stale (>7 days old), mention that things may have changed.
- If multiple session files match a label, use the most recent one.
