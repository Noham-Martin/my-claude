# nohamm-workflow v2.0

Complete development workflow ecosystem for Claude Code.

## Philosophy

- **Consistency over preference** — new code must feel like it always belonged.
- **TDD as non-negotiable** — RED, GREEN, IMPROVE. Always.
- **The pantry is the library** — every decision, learning, and pattern gets persisted. Knowledge never dies with the session.
- **Context hygiene** — fresh agent windows, short-term state in `.planning/`, long-term knowledge in the pantry.

---

## Day-to-Day Workflows

### Starting a new feature

```
/plan --add user authentication
```

The planner (opus) explores the codebase, finds existing patterns, and produces a detailed plan with REQ-IDs, wave-based task ordering, and a full pre-PR checklist. Review the plan, then start implementing.

For the full pipeline with automated review:

```
/orchestrate --feature
```

This chains: **planner** (plan it) → **tdd-guide** (test it) → **reviewer** (review it) → **security-reviewer** (secure it). Each agent runs in a fresh context. You get a final SHIP / NEEDS WORK / BLOCKED verdict. The orchestrator also exports the results to your pantry.

### Fixing a bug

Quick start — jump straight into systematic debugging:

```
/debug --login returns 500 on empty password
```

The debugger creates a persistent file in `.planning/debug/` with immutable symptoms, ranked hypotheses, and append-only evidence. If you need to stop and come back later, the debug file survives across sessions — just resume it:

```
/debug --resume
```

For the full bug pipeline (debug + test + review):

```
/orchestrate --bugfix
```

This chains: **debugger** (find it) → **tdd-guide** (test it) → **reviewer** (review the fix).

### Before pushing

Quick check during development (build + tests only):

```
/verify
```

Full quality gate before opening a PR (build + typecheck + tests + lint + security scan + git status):

```
/verify --full
```

Full gate plus a diff review and draft PR description:

```
/verify --pre-pr
```

### Reviewing code

Review a PR by number:

```
/code-review --pr 1234
```

Review current branch changes:

```
/code-review
```

The review covers completeness (missing call sites, configs, migrations), security (injection, auth, data exposure), tests, and code quality. After the verdict, it reminds you to update the pantry if the review surfaced decisions.

For a combined code + security review:

```
/orchestrate --review
```

### Refactoring safely

```
/orchestrate --refactor
```

Chains: **planner** (design the refactor) → **reviewer** (catch regressions) → **tdd-guide** (ensure test coverage).

### Build is broken

```
/build-fix
```

Auto-detects the build system (bazel, npm, cargo, go, make, maven, gradle). Fixes errors one at a time, re-running the build after each fix. Stops if a fix makes things worse.

### Tests are missing

```
/tests
```

Inspects the branch diff, finds untested behavior, locates reference test patterns from similar features, and backfills using strict TDD (write failing test → make it pass → improve).

---

## Session Management

Sessions let you save and restore your working context. Useful when you work on multiple things in parallel or need to pick up tomorrow where you left off today.

### Saving a session

```
/session-save --feature-auth
```

This saves to `.planning/sessions/2026-02-06-feature-auth.md`:
- What you accomplished
- What is still in progress (with file paths)
- Decisions made and their rationale
- Blockers and open questions
- Mental context (the "why" that isn't obvious from code)
- Git state snapshot

It also checks if anything from this session belongs in the pantry (decisions, architecture choices, learnings) and exports it automatically.

### Listing sessions

When running multiple sessions in parallel, see what's available:

```
/session-list
```

```
Saved sessions:

  Label              Date        Branch              Status
  ─────              ────        ──────              ──────
  feature-auth       2026-02-06  nohamm/auth         In progress: JWT middleware
  login-bug          2026-02-05  nohamm/fix-login    In progress: narrowing hypotheses
  metrics-v2         2026-02-03  nohamm/metrics      In progress: dashboard queries

  Total: 3 sessions
  Active debug sessions: 1
```

### Resuming a session

Resume a specific session by label:

```
/session-resume --feature-auth
```

Or resume the most recent one:

```
/session-resume
```

Loads the session file, verifies git state hasn't diverged, and suggests next actions.

### Quick status check

```
/progress
```

Reads all state and tells you the ONE most important thing to do next:
- "You have 3 unpushed commits — run `/verify --pre-pr` then push"
- "Previous session had work in progress on auth middleware — run `/session-resume --feature-auth`"
- "All clear — ready for new work"

### Checkpoints

Before risky operations (big refactors, migrations, trying a different approach):

```
/checkpoint --before-refactor
```

This creates a named save point (git stash + checkpoint log). If things go wrong:

```
/checkpoint --restore before-refactor
```

List all checkpoints:

```
/checkpoint --list
```

### Where state lives

```
.planning/
├── STATE.md           # Short-term memory (max 100 lines), read at session start
├── sessions/          # Full session snapshots (one per /session-save)
├── checkpoints.md     # Named git save points
└── debug/             # Persistent debug sessions (separate from session management)
    └── resolved/      # Completed debug sessions
```

---

## Debugging

Debugging is its own workflow, separate from session management. Debug files persist in `.planning/debug/` and survive across sessions automatically.

### Starting a debug session

```
/debug --login returns 500 on empty password
```

Creates `.planning/debug/debug-2026-02-06-login-500.md` with:
- **Symptoms** (immutable once written)
- **Hypotheses** ranked by likelihood
- **Evidence** (append-only, gathered step by step)
- **Resolution** (filled when the root cause is found)

### Resuming a debug session

```
/debug --resume
```

Picks up the most recent active debug session. All hypotheses, evidence, and eliminated causes are still there — no context lost.

### After resolving

The debug file moves to `.planning/debug/resolved/`. If the pattern is reusable, the debugger suggests running `/learn` to extract it as a skill.

---

## The Pantry

The pantry (`~/dd/dd-pantry/`) is your persistent knowledge library across projects and sessions. It is the most important persistence layer — sessions come and go, but the pantry is forever.

### When it gets updated

The pantry is wired into the workflow at every milestone:

| Event | What happens |
|-------|-------------|
| `/session-save` | Auto-exports decisions and learnings to the pantry |
| `/orchestrate` completes | Exports findings, architecture decisions, and outcomes |
| `/code-review` verdict | Reminds you to export if the review surfaced durable knowledge |
| `/plan` pre-PR checklist | "Pantry updated" is a mandatory checklist item |
| `/learn` | Extracts patterns as skills (which can also feed the pantry) |

### Commands

```
/import --auth-service       # Load context from ~/dd/dd-pantry/auth-service/
/export --auth-service       # Save new context (auto-numbered: 03-api-design.md)
/skill-create --auth-service # Turn pantry knowledge into a reusable skill
```

For project-local context (`.planning/` instead of pantry):

```
/import --project
/export --project
```

### What belongs in the pantry

- Decisions with rationale and alternatives considered
- Architecture and design choices
- Requirements discovered or refined
- Integration patterns and API contracts
- Debugging insights and workarounds
- PR context: what was changed, why, trade-offs

### What does NOT belong

- Ephemeral session state (that goes in `.planning/`)
- Build logs or test output
- Trivial fixes

**Rule of thumb:** if you'd want to remember it next month, it goes in the pantry.

---

## Knowledge Extraction

### After solving a hard problem

```
/learn
```

Reviews the current session for reusable patterns: error resolution techniques, debugging sequences, workarounds, project-specific conventions. Assigns a confidence score (0.0-1.0) and domain tag. Checks for duplicates against existing skills before saving to `~/.claude/skills/learned/`.

### From accumulated pantry notes

```
/skill-create --auth-service
```

Reads all pantry files for a topic, identifies ONE strong reusable pattern, drafts a skill with trigger conditions and pitfalls, and asks for confirmation before saving.

---

## Agents

Six specialized agents, each with restricted tools and a specific model:

| Agent | Model | What it does | Invoked by |
|-------|-------|-------------|------------|
| **planner** | opus | Explores codebase, produces detailed plans with REQ-IDs and waves | `/plan`, `/orchestrate --feature`, `/orchestrate --refactor` |
| **reviewer** | sonnet | Completeness-first code review with security checks | `/code-review`, `/orchestrate --*` |
| **tdd-guide** | sonnet | Strict RED-GREEN-IMPROVE test writing | `/tests`, `/orchestrate --feature`, `/orchestrate --bugfix` |
| **debugger** | sonnet | Systematic debugging with persistent state files | `/debug`, `/orchestrate --bugfix` |
| **security-reviewer** | sonnet | OWASP Top 10, injection, auth, data exposure | `/orchestrate --feature`, `/orchestrate --review` |
| **build-resolver** | sonnet | Auto-detect build system, fix errors one at a time | `/build-fix` |

Agents auto-invoke based on context too: complex feature request triggers planner, bug report triggers debugger, code changes trigger reviewer.

---

## Hooks

Always-on safety and quality enforcement:

| Hook | When | What it does |
|------|------|-------------|
| **safety-guard** | Before any Bash command | Blocks `rm -rf /`, force-push to main, `DROP TABLE`. Warns on `git reset --hard`, `git clean`. |
| **branch-guard** | Before any file edit | Blocks edits on main/master. Tells you to create a feature branch. |
| **lint-check** | After any file edit | Runs the appropriate linter (eslint/ruff/golangci-lint) in the background. |
| **stop-reminder** | After Claude stops | Warns if you have uncommitted changes. |

---

## Rules

Always-active guidelines that shape every interaction:

| Rule | Core idea |
|------|-----------|
| **coding-style** | Search for existing patterns before writing anything. Mirror what exists. |
| **git-workflow** | Branch: `nohamm/<name>`. Commits: imperative verb, single line, no co-authors. |
| **testing** | TDD always. Coverage derived from codebase, never invented. |
| **performance** | Opus for planning. Sonnet for execution. Plan mode for multi-file changes. |
| **agents** | Auto-invoke rules, handoff protocol, parallel execution when independent. |
| **sessions** | STATE.md under 100 lines. Pantry updated at every milestone. Debug files are the source of truth. |

---

## Quick Reference

| I want to... | Run |
|--------------|-----|
| Plan a feature | `/plan --<task>` |
| Build + test + review + security in one go | `/orchestrate --feature` |
| Fix a bug end-to-end | `/orchestrate --bugfix` |
| Debug something | `/debug --<issue>` |
| Resume a debug session | `/debug --resume` |
| Run tests for my changes | `/tests` |
| Fix build errors | `/build-fix` |
| Quick check before pushing | `/verify` |
| Full check before PR | `/verify --pre-pr` |
| Review a PR | `/code-review --pr <n>` |
| Save my work | `/session-save --<label>` |
| See all my saved sessions | `/session-list` |
| Resume a specific session | `/session-resume --<label>` |
| What should I do next? | `/progress` |
| Save point before risky change | `/checkpoint --<name>` |
| Load context from pantry | `/import --<folder>` |
| Save knowledge to pantry | `/export --<folder>` |
| Extract a reusable pattern | `/learn` |
| Turn pantry notes into a skill | `/skill-create --<folder>` |
| See all commands | `/help` |

---

## Directory Structure

```
nohamm-workflow/
├── .claude-plugin/plugin.json     # Plugin manifest (v2.0.0)
├── .claude/settings.local.json    # Permissions
├── agents/                        # 6 specialized sub-agents
│   ├── planner.md
│   ├── reviewer.md
│   ├── tdd-guide.md
│   ├── debugger.md
│   ├── security-reviewer.md
│   └── build-resolver.md
├── commands/                      # 17 slash commands
├── hooks/hooks.json               # Hook configuration
├── rules/                         # 6 always-active rules
├── scripts/                       # Hook helper scripts
│   ├── safety-guard.sh
│   ├── branch-guard.sh
│   ├── lint-check.sh
│   └── stop-reminder.sh
└── skills/                        # 4 skills with templates
    ├── orchestrate/
    ├── debug-session/
    ├── session-manager/
    └── verification-loop/
```
