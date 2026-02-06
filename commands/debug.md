---
description: Systematic debugging with persistent state across sessions
argument-hint: --<issue> | --list | --resume [--<name>] | --delete --<name>
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
---

# /debug — Persistent Debug Sessions

Systematic debugging that persists across sessions. Debug state lives in `.planning/debug/` so you never lose progress. Named sessions let you run multiple investigations in parallel and resume any of them by name.

## Usage
- `/debug --<issue description>` — start a new debug session
- `/debug --list` — list all debug sessions (active and resolved)
- `/debug --resume` — resume the most recent active session
- `/debug --resume --<name>` — resume a specific session by name/slug
- `/debug --delete --<name>` — delete a debug session by name

## Process

### 0) Route by flag

- If `--list`: skip to the **List sessions** section below.
- If `--resume`: skip to the **Resume session** section below.
- If `--delete`: skip to the **Delete session** section below.
- Otherwise: this is a new debug session — continue to step 1.

### 1) Check for existing sessions

Look in `.planning/debug/` for active sessions (files NOT in `resolved/` subfolder).

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

---

## List sessions

When invoked with `--list`.

### 1) Scan debug files

Scan `.planning/debug/` for all `*.md` files (both active and in `resolved/` subfolder).

If `.planning/debug/` does not exist or is empty: report "No debug sessions found." and stop.

### 2) Parse each file

For each file, extract:
- **Name/slug** (from filename: `debug-YYYY-MM-DD-<slug>.md` → the slug part)
- **Date** (from filename prefix or `Started:` line)
- **Status** (from `Status:` line — gathering, investigating, fixing, verifying, resolved)
- **Description** (from the `# Debug: <description>` heading)

### 3) Output (plain text)

Sort by date, most recent first. Group active sessions above resolved.

```
Debug sessions:

  Active:
  Name               Date        Status          Description
  ─────              ────        ──────          ───────────
  login-500          2026-02-06  investigating   Login returns 500 on empty password
  cache-miss         2026-02-05  gathering       Redis cache misses on user lookup

  Resolved:
  auth-loop          2026-02-03  resolved        Auth redirect loop after token refresh

  Total: 3 sessions (2 active, 1 resolved)

Resume a session:  /debug --resume --login-500
```

Keep it plain text, scannable, terminal-friendly. No Markdown.

---

## Resume session

When invoked with `--resume` (with or without `--<name>`).

### 1) Find the session to resume

- If `--<name>` is provided: look for `.planning/debug/debug-*-<name>.md` (match the slug portion). If not found, also try partial matches and report what was found.
- If no name: find the most recent active (non-resolved) debug file by date.
- If no active sessions exist: report "No active debug sessions. Run `/debug --list` to see all sessions." and stop.

### 2) Load and display context

Read the debug file and present:

```
Resuming debug session: <description>
  File: .planning/debug/<filename>
  Status: <status>
  Started: <date>

  Symptoms: <count> recorded
  Hypotheses: <active>/<total> still active
  Evidence: <count> entries

  Last evidence: <summary of most recent evidence entry>
```

### 3) Continue investigation

Pick up from where the session left off:
- If `Status: gathering` → continue recording symptoms and forming hypotheses.
- If `Status: investigating` → review active hypotheses and continue evidence gathering.
- If `Status: fixing` → continue applying and verifying the fix.
- If `Status: verifying` → verify the fix and close.

Skip to step 4 (Form hypotheses) or step 5 (Gather evidence) depending on status.

---

## Delete session

When invoked with `--delete --<name>`.

### 1) Find the session

Look for `.planning/debug/debug-*-<name>.md` (active) or `.planning/debug/resolved/debug-*-<name>.md` (resolved).

If not found: report "No debug session found matching '<name>'. Run `/debug --list` to see all sessions." and stop.

### 2) Confirm deletion

Show the session details (name, date, status, description) and ask for confirmation:

```
Delete debug session?
  Name: <name>
  Date: <date>
  Status: <status>
  Description: <description>

This will permanently delete the debug file. Proceed? (y/n)
```

### 3) Delete

Delete the file. Report:

```
Deleted debug session: <name>
```

---

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
- Debug file names use the slug as the canonical name for `--resume --<name>`.
