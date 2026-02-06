---
description: List all nohamm-workflow plugin commands
allowed-tools: []
---

# Nohamm Workflow Commands

## Development

| Command | Description |
|---------|-------------|
| `/plan --<task>` | Generate detailed end-to-end implementation plan |
| `/build-fix` | Incrementally fix build errors one at a time |
| `/tests` | Backfill missing tests based on current branch changes |
| `/debug --<issue>` | Systematic debugging with persistent state |
| `/debug --list` | List all debug sessions (active and resolved) |
| `/debug --resume [--<name>]` | Resume a debug session (by name or most recent) |
| `/verify [--quick\|--full\|--pre-pr]` | Quality gate: build, tests, lint, security, git status |

## Review

| Command | Description |
|---------|-------------|
| `/code-review --pr <n>` | Code review focused on completeness, security, and quality |
| `/orchestrate --<preset>` | Chain agents: feature, bugfix, refactor, review |

## Session Management

| Command | Description |
|---------|-------------|
| `/progress` | Show current status and suggest next action |
| `/session-save [--<label>]` | Save session state for later resumption |
| `/session-list` | List all saved sessions |
| `/session-resume [--<label>]` | Restore context from a previous session |
| `/checkpoint --<name>` | Create, list, or restore named git checkpoints |

## Knowledge

| Command | Description |
|---------|-------------|
| `/learn` | Extract reusable patterns from current session |
| `/import --<folder>` | Import context from dd-pantry or `--project` |
| `/export --<folder>` | Export context to dd-pantry or `--project` |
| `/skill-create --<folder>` | Create a reusable skill from pantry context |

## Meta

| Command | Description |
|---------|-------------|
| `/help` | Show this list |
