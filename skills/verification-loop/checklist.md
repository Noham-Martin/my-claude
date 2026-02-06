# Pre-PR Checklist

Verify each item before opening a PR. This checklist is used by `/verify --pre-pr`.

## Code

- [ ] Build passes
- [ ] Type checks pass (if applicable)
- [ ] Lint passes
- [ ] No hardcoded secrets in diff
- [ ] No debug statements left (console.log, print, debugger)
- [ ] No large commented-out code blocks
- [ ] No TODO/FIXME left in new code (unless intentional and tracked)

## Tests

- [ ] Related tests pass
- [ ] New behavior has test coverage
- [ ] Test style matches existing codebase conventions
- [ ] No test regressions

## Completeness

- [ ] All call sites updated for changed interfaces
- [ ] Config/feature flags wired
- [ ] DB migrations included (if schema changed)
- [ ] Docs updated (if user-facing behavior changed)
- [ ] Error handling covers new failure modes

## Git

- [ ] On a feature branch (not main/master)
- [ ] All changes committed
- [ ] Commit messages follow conventions (imperative verb, single line)
- [ ] Branch pushed to remote

## PR

- [ ] PR title is concise (<70 chars)
- [ ] PR description explains what, why, and how tested
