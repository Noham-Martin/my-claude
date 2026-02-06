---
name: debug-session
description: Systematic debugging with persistent state. Provides debug file templates and methodology for the /debug command and debugger agent.
user-invocable: false
---

# Debug Session Skill

This skill provides the debugging methodology and templates for persistent debug sessions. Debug sessions work like regular sessions — they are named, listable, and resumable by name.

## Debug File Location

All debug state lives in `.planning/debug/` relative to the project root.
Resolved sessions move to `.planning/debug/resolved/`.

## Naming Convention

Files are named `debug-YYYY-MM-DD-<slug>.md`. The `<slug>` is the canonical session name used for resume:
- `/debug --resume --<slug>` resumes a specific session.
- `/debug --list` shows all sessions with their slugs.

## Session Lifecycle

1. **Start** — `/debug --<issue>` creates a new named debug file.
2. **During** — the debug file is updated before every action (the file IS the brain).
3. **List** — `/debug --list` shows all active and resolved sessions.
4. **Resume** — `/debug --resume --<name>` picks up a specific session. `/debug --resume` picks up the most recent.
5. **Resolve** — file moves to `.planning/debug/resolved/`.

## Methodology

1. **Gather symptoms** — record what is failing (immutable once written).
2. **Form hypotheses** — list 3-5 possible causes ranked by likelihood.
3. **Gather evidence** — test each hypothesis, record all findings (append-only).
4. **Eliminate** — mark hypotheses as ELIMINATED or CONFIRMED with evidence.
5. **Resolve** — apply minimal fix, verify, document.

## Key Rules

- Update the debug file BEFORE taking actions, not after.
- The Symptoms section is immutable after initial write.
- The Evidence section is append-only.
- The debug file IS the debugging brain — it must always reflect current state.

## Debug File Template

Use the template in `debug-template.md` for new debug sessions.
