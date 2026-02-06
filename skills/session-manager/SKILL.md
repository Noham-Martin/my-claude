---
name: session-manager
description: Session persistence and resumption. Provides templates for STATE.md and session files used by /session-save, /session-resume, and /progress.
user-invocable: false
---

# Session Manager Skill

This skill provides the templates and conventions for session management across the workflow.

## File Locations

- `.planning/STATE.md` — short-term memory (max 100 lines), read at session start.
- `.planning/sessions/YYYY-MM-DD-<label>.md` — full session snapshots.

## STATE.md

Quick-reference file. Must stay under 100 lines. Updated at logical boundaries.

See `state-template.md` for the structure.

## Session Files

Full snapshots for resumption. Created by `/session-save`, read by `/session-resume`.

See `session-template.md` for the structure.

## Session Lifecycle

1. **Start** → Read STATE.md for context.
2. **During** → Update STATE.md at task boundaries.
3. **End** → `/session-save` creates full snapshot + updates STATE.md.
4. **Resume** → `/session-resume` loads STATE.md + latest session file.
