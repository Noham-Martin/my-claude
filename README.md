# nohamm-workflow

Personal Claude Code plugin for consistent development workflow.

## Rules

| Rule | Purpose |
|------|---------|
| coding-style | Consistency over preference. Match existing codebase patterns. |
| git-workflow | Branch naming (`nohamm/<name>`), commit style, PR descriptions. |
| testing | TDD workflow. Derive coverage from codebase, not ideals. |
| performance | Model selection and Plan Mode guidance. |

## Commands

| Command | Purpose |
|---------|---------|
| `/plan` | Detailed end-to-end implementation plan |
| `/code-review` | Completeness-first review |
| `/verify` | Quick pre-PR check (build, tests, lint, git status) |
| `/tests` | Backfill missing tests from branch history |
| `/build-fix` | Iteratively fix build errors |
| `/learn` | Extract reusable patterns from session |
| `/import --<folder>` | Load context from pantry |
| `/export --<folder>` | Save context to pantry |
| `/skill-create --<folder>` | Convert pantry context to skill |

## Pantry System

Context persistence via `~/dd/dd-pantry/`:
- `/import` loads existing notes
- `/export` saves new notes (auto-numbered)
- `/skill-create` converts notes into reusable skills
