---
description: Save current session state for later resumption
argument-hint: [--<label>]
allowed-tools: [Read, Glob, Grep, Bash, Write]
---

# /session-save — Save Session State

Create a handoff file so work can be resumed in a new session without losing context.

## Usage
- `/session-save` — save with auto-generated label (date-based)
- `/session-save --<label>` — save with a custom label

## Process

### 1) Analyze current session

Review what happened in this session:
- What tasks were worked on
- What files were modified
- What decisions were made
- What is still in progress
- What is blocked or needs attention

### 2) Capture git state

Run these commands and record the output:
- `git branch --show-current` — current branch
- `git status --short` — uncommitted changes
- `git log --oneline -5` — recent commits
- `git diff --stat` — unstaged change summary

### 3) Write session file

Create `.planning/sessions/YYYY-MM-DD-<label>.md` (create directories if needed).

If no label is provided, use the current branch name or a summary slug.

File structure:

```markdown
---
date: YYYY-MM-DD
branch: <current branch>
label: <label>
---

# Session: <label>

## Accomplished
- <what was completed this session>

## In Progress
- <what is partially done, with file paths and line numbers>
- <current state of each item>

## Decisions Made
- <decision>: <rationale>

## Blockers & Open Questions
- <anything unresolved>

## Mental Context
- <the "why" behind the current approach>
- <things that aren't obvious from the code alone>
- <gotchas discovered that the next session needs to know>

## Git State
- Branch: <branch>
- Uncommitted changes: <list>
- Recent commits: <list>

## Next Steps
1. <first thing to do when resuming>
2. <second thing>
3. <third thing>
```

### 4) Update STATE.md

Write or update `.planning/STATE.md` as short-term memory (capped at 100 lines).

STATE.md is a compressed version of the session file — just the essentials:

```markdown
# State

Last updated: YYYY-MM-DD
Branch: <branch>
Session: <path to session file>

## Current Focus
<1-2 lines: what you are working on>

## Key Decisions
- <most important recent decisions>

## Next Steps
1. <first priority>
2. <second priority>
3. <third priority>

## Blockers
- <anything blocking progress>
```

If STATE.md already exists and exceeds 100 lines after update, truncate older entries (keep the most recent state).

### 5) Pantry export

The pantry (`~/dd/dd-pantry/`) is the persistent cross-project knowledge library. It MUST stay up to date.

After writing the session file, check if this session produced knowledge worth persisting:
- **Decisions made** that affect the project beyond this session
- **Architecture or design choices** with rationale
- **Learnings** (debugging insights, workarounds, integration patterns)
- **Requirements discovered** or refined during implementation

If any of the above exist:
1. Identify the relevant pantry folder. Check if a folder already exists for this project or feature:
   - List folders in `~/dd/dd-pantry/` and look for a match.
   - If no match: ask the user which folder to use or whether to create one.
2. Run the `/export` logic: create a new numbered file in that pantry folder with the decisions, learnings, and context from this session.
3. Focus the pantry export on **durable knowledge** — not ephemeral session state. The pantry is a library, not a scratchpad.

If nothing worth persisting: skip and note "No pantry-worthy context this session."

### 6) Confirm

Output:
```
Session saved.
  File: .planning/sessions/YYYY-MM-DD-<label>.md
  State: .planning/STATE.md (updated)
  Pantry: <path to exported file> (or "skipped — no durable context")

To resume later: /session-resume
```

## Rules

- Never invent facts about the session. Only record what actually happened.
- If unsure about something, mark it as "uncertain" rather than omitting it.
- Always include file paths and line numbers for in-progress work.
- Keep STATE.md under 100 lines — it is read at every session start.
- Always check if the pantry needs updating. The pantry is the persistent library — if decisions were made, it must be updated.
