---
description: Create, list, or restore named git checkpoints with session state
argument-hint: --<name> | --list | --restore <name>
allowed-tools: [Read, Glob, Grep, Bash, Write]
---

# /checkpoint — Named Git Checkpoints

Create named save points that combine git state + session context. Useful for risky operations or before trying different approaches.

## Usage
- `/checkpoint --<name>` — create a named checkpoint
- `/checkpoint --list` — list all checkpoints
- `/checkpoint --restore <name>` — restore a checkpoint

## Process

### Creating a checkpoint (`--<name>`)

1. **Capture git state**:
   - `git stash create` — create a stash object without modifying working tree
   - `git rev-parse HEAD` — current commit hash
   - `git branch --show-current` — current branch
   - `git diff --stat` — summary of uncommitted changes

2. **If there are uncommitted changes**: also run `git stash push -m "checkpoint: <name>"` to save them.

3. **Write checkpoint entry** to `.planning/checkpoints.md` (create if needed):

   Append:
   ```markdown
   ## <name>
   - Date: YYYY-MM-DD HH:MM
   - Branch: <branch>
   - Commit: <hash> (<commit message>)
   - Stash: <stash ref or "none">
   - Changes: <summary of uncommitted changes, or "clean">
   - Status: saved
   ```

4. **Output**:
   ```
   Checkpoint created: <name>
     Branch: <branch>
     Commit: <short hash>
     Stash: <yes/no>
   ```

### Listing checkpoints (`--list`)

1. Read `.planning/checkpoints.md`.
2. If it does not exist: report "No checkpoints found."
3. Output a plain text table:

   ```
   Name            Date        Branch       Commit   Stash
   ──────────────  ──────────  ───────────  ───────  ─────
   before-refactor 2026-02-06  nohamm/feat  a1b2c3d  yes
   pre-migration   2026-02-05  nohamm/feat  d4e5f6a  no
   ```

### Restoring a checkpoint (`--restore <name>`)

1. Read `.planning/checkpoints.md` and find the named checkpoint.
2. If not found: report "Checkpoint '<name>' not found." and list available ones.
3. **Warn the user** before restoring:
   ```
   WARNING: Restoring checkpoint '<name>' will:
   - Switch to branch: <branch>
   - Reset to commit: <hash>
   - Apply stashed changes (if any)

   Current uncommitted changes will be lost unless you create a checkpoint first.
   Proceed? (Waiting for confirmation)
   ```
4. Wait for user confirmation.
5. If confirmed:
   - `git checkout <branch>` (if different from current)
   - `git reset --hard <commit>` (if commit differs from HEAD)
   - `git stash pop` (if stash exists for this checkpoint)
   - Update checkpoint status to "restored" in checkpoints.md
6. **Output**:
   ```
   Checkpoint restored: <name>
     Branch: <branch>
     Commit: <hash>
     Stash applied: <yes/no>
   ```

## Rules

- Always warn before destructive operations (restore).
- Never auto-restore without user confirmation.
- Checkpoint names must be unique. If a name already exists, ask for a different one.
- Keep `.planning/checkpoints.md` as the single source of truth.
- Output is plain text, not Markdown.
