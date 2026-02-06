---
name: debug-session
description: Systematic debugging with persistent state. Provides debug file templates and methodology for the /debug command and debugger agent.
user-invocable: false
---

# Debug Session Skill

This skill provides the debugging methodology and templates for persistent debug sessions.

## Debug File Location

All debug state lives in `.planning/debug/` relative to the project root.
Resolved sessions move to `.planning/debug/resolved/`.

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
