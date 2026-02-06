# Session Management Rules

## State Files

All session state lives in `.planning/` relative to the project root:
- `.planning/STATE.md` — short-term memory, read at every session start.
- `.planning/sessions/` — full session snapshots for resumption.
- `.planning/debug/` — persistent debug sessions.
- `.planning/checkpoints.md` — named git checkpoints.

## STATE.md

- STATE.md is the quick-reference file. Keep it under **100 lines**.
- Update STATE.md at logical boundaries:
  - After completing a task
  - Before switching context
  - Before ending a session
- When STATE.md exceeds 100 lines, truncate older entries and keep only the most recent state.
- STATE.md should always contain: current focus, key decisions, next steps, and blockers.

## Session Lifecycle

1. **Start**: Check for `.planning/STATE.md`. If it exists, read it for context.
2. **During**: Update STATE.md at logical boundaries.
3. **End**: Run `/session-save` to create a full session snapshot.
4. **Resume**: Run `/session-resume` to restore context from the last session.

## Debug Sessions

- Debug state persists in `.planning/debug/`.
- Active sessions are files with `Status: gathering|investigating|fixing|verifying`.
- Resolved sessions are moved to `.planning/debug/resolved/`.
- Debug files are the single source of truth for debugging — always update the file before taking actions.

## Checkpoints

- Checkpoints combine git state + session context.
- Use checkpoints before risky operations (refactors, migrations, experimental changes).
- Checkpoint names must be unique and descriptive.
