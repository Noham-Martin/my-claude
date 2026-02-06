---
name: verification-loop
description: Multi-phase quality gate for verifying project readiness. Provides the checklist and methodology used by the /verify command.
user-invocable: false
---

# Verification Loop Skill

This skill provides the quality gate methodology and checklists for the `/verify` command.

## Modes

| Mode | Steps |
|------|-------|
| `--quick` | Build + Tests |
| `--full` | Build + Type Check + Tests + Lint + Security Scan + Git Status |
| `--pre-pr` | Full + Diff Review + PR Description Draft |

## Step Order

Steps run sequentially. Stop on first failure (except git status, which always runs).

1. **Build** — detect build system, run build.
2. **Type Check** — detect type checker, run it.
3. **Tests** — find related tests from diff, run only those.
4. **Lint** — detect linter, run on changed files.
5. **Security Scan** — scan diff for hardcoded secrets, eval/exec, disabled security features.
6. **Git Status** — branch, uncommitted changes, unpushed commits.
7. **Diff Review** — scan for TODOs, debug statements, commented-out code.
8. **PR Description** — draft title + description from diff and commit history.

## Pre-PR Checklist

See `checklist.md` for the full pre-PR verification checklist.
