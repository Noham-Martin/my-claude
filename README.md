# nohamm-workflow v2.0

Complete development workflow ecosystem for Claude Code: agents, orchestration, session management, context engineering, and quality enforcement.

## Philosophy

- **Consistency over preference** — new code must feel like it always belonged.
- **TDD as non-negotiable** — RED → GREEN → IMPROVE, always.
- **Structured knowledge capture** — every insight is persistent and reusable.
- **Context hygiene** — fresh agent windows, strategic state management.

## Architecture

```
User Commands (/plan, /verify, /debug, ...)
        │
   Commands Layer (16 commands — thin wrappers)
        │
    Agents Layer (6 agents — specialized, model-routed)
        │
    Hooks Layer (safety guards, quality checks, reminders)
        │
  Persistence Layer (.planning/, dd-pantry/, skills/)
```

## Commands (16)

### Development

| Command | Description |
|---------|-------------|
| `/plan --<task>` | Detailed implementation plan with REQ-IDs and wave ordering |
| `/build-fix` | Incremental build error fixing (auto-detects build system) |
| `/tests` | Backfill missing tests from branch changes (TDD discipline) |
| `/debug --<issue>` | Systematic debugging with persistent state in `.planning/debug/` |
| `/verify [--quick\|--full\|--pre-pr]` | Quality gate: build, typecheck, tests, lint, security, git |

### Review

| Command | Description |
|---------|-------------|
| `/code-review --pr <n>` | Completeness-first review with security section |
| `/orchestrate --<preset>` | Chain agents: `feature`, `bugfix`, `refactor`, `review` |

### Session Management

| Command | Description |
|---------|-------------|
| `/progress` | Show status and suggest next action |
| `/session-save [--<label>]` | Save session state for later resumption |
| `/session-resume [--<label>]` | Restore context from a previous session |
| `/checkpoint --<name>` | Create, list, or restore named git checkpoints |

### Knowledge

| Command | Description |
|---------|-------------|
| `/learn` | Extract reusable patterns with confidence scoring and domain tags |
| `/import --<folder>` | Import from dd-pantry or `--project` for `.planning/` |
| `/export --<folder>` | Export to dd-pantry or `--project` with YAML frontmatter |
| `/skill-create --<folder>` | Create reusable skill from pantry context |

### Meta

| Command | Description |
|---------|-------------|
| `/help` | List all commands |

## Agents (6)

| Agent | Model | Tools | Auto-invoked when |
|-------|-------|-------|-------------------|
| **planner** | opus | Read, Glob, Grep, Bash | Complex feature requests, multi-file changes |
| **reviewer** | sonnet | Read, Glob, Grep, Bash | Code review needed after changes |
| **tdd-guide** | sonnet | Read, Glob, Grep, Bash, Edit, Write | Bug fixes, test creation |
| **debugger** | sonnet | Read, Glob, Grep, Bash, Edit, Write | Investigating bugs or failures |
| **security-reviewer** | sonnet | Read, Glob, Grep | Changes touching auth, input, APIs, sensitive data |
| **build-resolver** | sonnet | Read, Glob, Grep, Bash, Edit | Build failures |

## Rules (6)

| Rule | Purpose |
|------|---------|
| **coding-style** | Consistency over preference. Match existing patterns. |
| **git-workflow** | `nohamm/<branch>`, imperative commits, no co-authors. |
| **testing** | Strict TDD. Coverage derived from codebase. |
| **performance** | Model routing: opus for planning, sonnet for execution. |
| **agents** | Auto-invocation rules, handoff protocol, context hygiene. |
| **sessions** | STATE.md conventions, session lifecycle, debug persistence. |

## Skills (4)

| Skill | Purpose |
|-------|---------|
| **orchestrate** | Agent chaining templates and handoff format |
| **debug-session** | Debug file templates and methodology |
| **session-manager** | STATE.md and session file templates |
| **verification-loop** | Quality gate methodology and pre-PR checklist |

## Hooks

| Event | Hook | Behavior |
|-------|------|----------|
| PreToolUse (Bash) | `safety-guard.sh` | Blocks `rm -rf /`, force-push to main, DROP TABLE |
| PreToolUse (Edit/Write) | `branch-guard.sh` | Blocks file edits on main/master branch |
| PostToolUse (Edit/Write) | `lint-check.sh` | Async lint on edited files |
| Stop | `stop-reminder.sh` | Warns about uncommitted changes |

## Session Management

### Save and resume workflow

```
# End of session
/session-save --feature-auth

# New session
/session-resume --feature-auth
```

State files live in `.planning/` per project:
- `STATE.md` — short-term memory (max 100 lines)
- `sessions/` — full session snapshots
- `debug/` — persistent debug sessions
- `checkpoints.md` — named git save points

### Quick status check

```
/progress
```

Reads all state and suggests the single most important next action.

## Orchestration

Chain agents for end-to-end workflows:

```
/orchestrate --feature    # planner → tdd-guide → reviewer → security-reviewer
/orchestrate --bugfix     # debugger → tdd-guide → reviewer
/orchestrate --refactor   # planner → reviewer → tdd-guide
/orchestrate --review     # reviewer → security-reviewer
```

Each agent gets a fresh context window. Results pass via standardized handoff documents. Final verdict: SHIP / NEEDS WORK / BLOCKED.

## Pantry System

Cross-project knowledge persistence via `~/dd/dd-pantry/`:
- `/import --<folder>` loads existing notes
- `/export --<folder>` saves new notes (auto-numbered with YAML frontmatter)
- `/skill-create --<folder>` converts notes into reusable skills
- `/import --project` and `/export --project` work with `.planning/` instead

## Directory Structure

```
nohamm-workflow/
├── .claude-plugin/plugin.json     # Plugin manifest (v2.0.0)
├── .claude/settings.local.json    # Permissions
├── agents/                        # 6 specialized sub-agents
├── commands/                      # 16 slash commands
├── hooks/hooks.json               # Hook configuration
├── rules/                         # 6 always-active rules
├── scripts/                       # Hook helper scripts
└── skills/                        # 4 skills with templates
```
