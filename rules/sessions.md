# Session Management Rules

## State Files

All session state lives in `.planning/` relative to the project root:
- `.planning/STATE.md` — short-term memory, read at every session start.
- `.planning/sessions/` — full session snapshots for resumption.
- `.planning/debug/` — persistent debug sessions.

## Session Lifecycle

Session management is opt-in. Use `/session-save` and `/session-resume` when you want persistence. No automatic session tracking or pantry export during normal work.

## Debug Sessions

- Debug state persists in `.planning/debug/`.
- Active sessions are files with `Status: gathering|investigating|fixing|verifying`.
- Resolved sessions are moved to `.planning/debug/resolved/`.

## Pantry Integration

The pantry (`~/dd/dd-pantry/`) is updated when you explicitly run `/export`, `/session-save`, or `/learn`. No automatic pantry sync during regular work.
