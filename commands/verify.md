---
description: Quick project verification - checks build, tests, lint, and git status
argument-hint: [--quick|--full|--pre-pr]
allowed-tools: [Read, Glob, Grep, Bash]
---

# /verify — Quick project verification

Run a sequential quality gate to check if the project is ready.

Output is **plain text** (NOT Markdown). Keep it concise and CLI-friendly.

## Modes

- `--quick` (default): Build + Tests only. Fast check during development.
- `--full`: Build + Type Check + Tests + Lint + Security Scan + Git Status.
- `--pre-pr`: Full + Diff Review + PR description draft.

Parse the mode from the arguments. If no flag is provided, default to `--quick`.

---

## Steps (run in order, stop on first failure unless noted)

### 1) Build check
- Detect the build system: look for BUILD/BUILD.bazel (bazel), package.json (npm/yarn/pnpm), Cargo.toml (cargo), go.mod (go), Makefile (make).
- Run the appropriate build command.
- Report: PASS or FAIL.
- If FAIL: show the first error and STOP. Suggest running `/build-fix`.

### 2) Type check (--full and --pre-pr only)
- Detect type checker: `tsc --noEmit` for TypeScript, `mypy` for Python, `go vet` for Go, etc.
- Run the type checker.
- Report: PASS, FAIL (with error count), or SKIP (no type checker detected).

### 3) Tests
- Check `git diff --name-only` to identify changed files.
- Find related test files (same directory, matching naming convention).
- Run only those tests (not the full suite).
- Report: PASS (with count) or FAIL (with failing test names).
- If no related tests found: report SKIP with a note.

### 4) Lint (--full and --pre-pr only)
- Detect linter: eslint, ruff/flake8, golangci-lint, etc.
- Run the linter on changed files only.
- Report: PASS, FAIL (with issue count), or SKIP (no linter detected).

### 5) Security scan (--full and --pre-pr only)
- Quick scan of `git diff` for obvious security issues:
  - Hardcoded secrets (API keys, passwords, tokens in string literals)
  - `eval()`, `exec()`, `dangerouslySetInnerHTML` usage
  - Disabled security features (e.g., `verify=False`, `--no-verify`)
- Report: PASS, WARN (with findings), or SKIP.

### 6) Git status
- Report:
  - Current branch name
  - Unstaged changes (count)
  - Staged but uncommitted changes (count)
  - Commits ahead of remote (count)
  - Whether the branch has a remote tracking branch

### 7) Diff review (--pre-pr only)
- Run `git diff $(git merge-base HEAD main)..HEAD --stat` for a summary.
- Quick scan of the full diff for obvious issues: TODO/FIXME left in, debug statements (console.log, print, debugger), large commented-out blocks.
- Report findings if any.

### 8) PR description draft (--pre-pr only)
- Based on the diff and commit history, draft a PR title and description.
- Follow git-workflow.md rules for PR descriptions.

---

## Final Output Format (plain text)

```
Mode: quick | full | pre-pr

Build     ... PASS | FAIL
TypeCheck ... PASS | FAIL | SKIP       (full/pre-pr only)
Tests     ... PASS (N) | FAIL | SKIP
Lint      ... PASS | FAIL | SKIP       (full/pre-pr only)
Security  ... PASS | WARN | SKIP       (full/pre-pr only)
Git       ... branch: X, unstaged: N, staged: N, ahead: N

[Diff review findings, if any]         (pre-pr only)
[Draft PR title + description]         (pre-pr only)

Verdict: READY TO PUSH | NOT READY TO PUSH
[Reason if not ready]
```

Do NOT use Markdown headers, bold, or formatting in the output. Keep it plain text, scannable, terminal-friendly.
