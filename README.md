# nohamm-workflow v2.0

Complete development workflow ecosystem for Claude Code.

---

## Sessions

Save and restore working context. State lives in `.planning/sessions/`.

| Command | What it does |
|---------|-------------|
| `/session-save --<label>` | Save current session (also exports to pantry) |
| `/session-list` | List all saved sessions |
| `/session-resume --<label>` | Resume a specific session |
| `/session-resume` | Resume most recent session |
| `/session-delete --<label>` | Delete a saved session |
| `/progress` | Quick status check — suggests the ONE next action |

## Debug Sessions

Named, persistent debug investigations. State lives in `.planning/debug/`.

| Command | What it does |
|---------|-------------|
| `/debug --<issue>` | Start a new debug session |
| `/debug --list` | List all debug sessions (active + resolved) |
| `/debug --resume` | Resume most recent active session |
| `/debug --resume --<name>` | Resume a specific session by slug |
| `/debug --delete --<name>` | Delete a debug session |

Each session is a file (`debug-YYYY-MM-DD-<slug>.md`) with immutable symptoms, ranked hypotheses, and append-only evidence. Resolved sessions move to `.planning/debug/resolved/`.

## Review

| Command | What it does |
|---------|-------------|
| `/code-review` | Review current branch changes |
| `/code-review --pr <n>` | Review a PR by number |

## Orchestration

Chain agents end-to-end. **Bold steps** are implementation pauses — Claude presents the plan and waits for your go-ahead before coding.

| Command | Chain |
|---------|-------|
| `/orchestrate --feature --<prompt>` | planner → **implement** → tdd-guide → reviewer → security-reviewer |
| `/orchestrate --bugfix --<prompt>` | debugger → **implement fix** → tdd-guide → reviewer |
| `/orchestrate --refactor --<prompt>` | planner → **implement** → tdd-guide → reviewer |
| `/orchestrate --review` | reviewer → security-reviewer |

## Pantry

Persistent cross-project knowledge library at `~/dd/dd-pantry/`. Auto-updated by `/session-save`, `/orchestrate`, and `/plan`.

| Command | What it does |
|---------|-------------|
| `/import --<folder>` | Load context from pantry |
| `/export --<folder>` | Save context to pantry |
| `/learn` | Extract reusable patterns from current session |
| `/skill-create --<folder>` | Turn pantry notes into a reusable skill |

## Code

| Command | What it does |
|---------|-------------|
| `/plan --<task>` | Generate implementation plan (opus, REQ-IDs, wave ordering) |
| `/tests` | Backfill missing tests using strict TDD |
| `/build-fix` | Auto-detect build system, fix errors one at a time |
| `/verify` | Quick check: build + tests |
| `/verify --full` | Full gate: build + typecheck + tests + lint + security |
| `/verify --pre-pr` | Full gate + diff review + draft PR description |