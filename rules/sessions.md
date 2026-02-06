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

## Pantry Integration

The pantry (`~/dd/dd-pantry/`) is the persistent cross-project knowledge library. It MUST stay current.

**When to update the pantry:**
- After creating a PR — export what was built, why, and key decisions.
- After modifying a PR — export what changed and why.
- After a code review — export decisions, trade-offs, or patterns discovered.
- After an orchestration chain completes — export key findings and outcomes.
- After a session with significant decisions — `/session-save` handles this automatically.
- After resolving a non-trivial bug — export the debugging pattern via `/learn` or `/export`.

**What belongs in the pantry:**
- Decisions with rationale and alternatives considered
- Architecture and design choices
- Requirements discovered or refined
- Integration patterns and API contracts
- Debugging insights and workarounds
- Trade-offs and their justifications

**What does NOT belong in the pantry:**
- Ephemeral session state (that goes in `.planning/`)
- Build logs or test output
- Trivial fixes or one-off issues

**Rule:** If you are unsure whether something belongs in the pantry, export it. It is better to have too much in the library than to lose knowledge.
