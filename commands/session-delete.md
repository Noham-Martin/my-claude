---
description: Delete a saved session
argument-hint: --<label>
allowed-tools: [Read, Glob, Grep, Bash, Write]
---

# /session-delete — Delete a Saved Session

Remove a saved session file from `.planning/sessions/`.

## Usage
- `/session-delete --<label>` — delete a specific session by label

## Process

### 1) Find the session file

Look for `.planning/sessions/*-<label>.md`.

If not found: report "No session found matching '<label>'. Run `/session-list` to see all sessions." and stop.

If multiple files match (e.g. same label on different dates), list them and ask which one to delete.

### 2) Confirm deletion

Show the session details and ask for confirmation:

```
Delete session?
  Label: <label>
  Date: <date>
  Branch: <branch>
  Status: <first bullet from In Progress or Accomplished>

This will permanently delete the session file. Proceed? (y/n)
```

### 3) Delete

Delete the file. Report:

```
Deleted session: <label>
  File: .planning/sessions/<filename>
```

### 4) Update STATE.md if needed

If `.planning/STATE.md` references the deleted session, remove that reference.

## Rules

- Always require a label. Do not allow deleting without specifying which session.
- Always confirm before deleting.
- Do not delete STATE.md or any other files — only the session file itself.
