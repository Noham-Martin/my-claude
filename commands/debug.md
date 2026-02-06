---
description: Systematic debugging with persistent state across sessions
argument-hint: --<issue> | --resume
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
---

# /debug — Persistent Debug Sessions

Systematic debugging that persists across sessions. Debug state lives in `.planning/debug/` so you never lose progress.

## Usage
- `/debug --<issue description>` — start a new debug session
- `/debug --resume` — resume the most recent active session

## Process

### 1) Check for existing sessions

Look in `.planning/debug/` for active sessions (files NOT in `resolved/` subfolder).

- If `--resume` flag: load the most recent active debug file and skip to step 4.
- If `--<issue>` flag: check if an active session already matches this issue. If so, ask whether to resume it or start fresh.
- If `.planning/debug/` does not exist: create it.

### 2) Create debug file

Create `.planning/debug/debug-YYYY-MM-DD-<slug>.md` where `<slug>` is a short hyphenated description of the issue.

Write the initial structure BEFORE starting any investigation:

```markdown
# Debug: <short description>
Status: gathering
Started: YYYY-MM-DD

## Symptoms (immutable — do not edit after initial write)
- <what is failing>
- <error messages, stack traces>
- <expected vs observed behavior>
- <reproduction steps>

## Hypotheses
(to be filled)

## Evidence (append-only)
(to be filled)

## Resolution
(to be filled when resolved)
```

### 3) Record symptoms

Gather all available information about the issue:
- Read error messages, logs, stack traces.
- Ask the user for reproduction steps if not provided.
- Write symptoms to the debug file. Once written, the Symptoms section is **immutable**.

### 4) Form hypotheses

List 3-5 possible root causes, ranked by likelihood:

```markdown
## Hypotheses
1. [ACTIVE] Most likely cause — <description>
   - Evidence needed: <what would confirm or eliminate>
2. [ACTIVE] Second possibility — <description>
   - Evidence needed: <what would confirm or eliminate>
3. [ACTIVE] Less likely — <description>
   - Evidence needed: <what would confirm or eliminate>
```

### 5) Gather evidence and eliminate

For each hypothesis, starting with the most likely:

1. Describe what you are about to check.
2. **Update the debug file BEFORE taking the action** (the file IS the debugging brain).
3. Run the investigation (search code, read files, run commands).
4. Record findings in the Evidence section (append-only):

```markdown
### YYYY-MM-DD HH:MM — <what was checked>
- Finding: <what was observed>
- Implication: <what this means for the hypotheses>
```

5. Update hypothesis status: `[ACTIVE]` → `[ELIMINATED]` (with reason) or `[CONFIRMED]`.

Repeat until one hypothesis is confirmed or all are eliminated.

### 6) Resolve

Once root cause is confirmed:
1. Update Status to `fixing`.
2. Propose a minimal fix.
3. Apply the fix.
4. Verify the fix resolves the original symptoms.
5. Verify no regressions (run related tests if they exist).
6. Update Status to `resolved`.
7. Fill in the Resolution section:

```markdown
## Resolution
- Root cause: <what was actually wrong>
- Fix applied: <what was changed, with file paths>
- Verified: yes
- Related tests: <tests that cover this, if any>
```

8. Move the file to `.planning/debug/resolved/` (create the folder if needed).

### 7) Suggest learning

If the pattern seems reusable, suggest:
"This debugging pattern might be worth saving. Run `/learn` to extract it as a skill."

## Stop Conditions

- If stuck after 3 rounds of evidence gathering with no progress: pause and ask the user for more context.
- If all hypotheses are eliminated: form new hypotheses based on what was learned, or escalate.
- Never make speculative fixes without evidence.

## Rules

- Update the debug file BEFORE taking actions, not after.
- Never modify the Symptoms section after initial write.
- Only append to the Evidence section.
- Prefer reading code and logs over making speculative changes.
- Keep the debug file as the single source of truth.
