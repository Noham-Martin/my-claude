---
description: Code review focused on completeness, quality, and improvements
argument-hint: --pr <number-or-url>
allowed-tools: [Read, Glob, Grep, Bash]
---

# /code-review — Completeness-first, then quality, then improvements

You are reviewing to ensure:
1) NOTHING is missing (all required modifications done)
2) Code is coherent, readable, and consistent with the project
3) Security is not compromised
4) Optional improvements: performance, maintainability

## Hard requirements
- Follow rules/testing.md, rules/coding-style.md, rules/git-workflow.md.
- First prioritize completeness over style or optimization.
- Be actionable: each finding must say where + what + how to fix.

---

## Step 0 — Inputs

### Option A: `--pr <number-or-url>` provided
If the user passed `--pr`:
1. Extract the PR identifier (number or full URL).
2. Fetch PR metadata and diff using the `gh` CLI:
   - `gh pr view <pr> --json title,body,url,headRefName,baseRefName`
   - `gh pr diff <pr>`
3. Review the PR diff (not the current branch).

### Option B: No `--pr` flag
Review the current branch changes:
1. Determine the base branch (usually `main` or `master`).
2. Get the diff: `git diff $(git merge-base HEAD <base>)..HEAD`
3. List changed files: `git diff --name-only $(git merge-base HEAD <base>)..HEAD`

### Option C: No diff available
If neither option works, ask for:
- list of changed files OR a patch/diff
But still provide the completeness checklist immediately.

## Diff Size Awareness

Before starting the review, count the diff size:
- **Small** (<300 lines): review the full diff inline.
- **Medium** (300-1000 lines): review file-by-file, group findings by file.
- **Large** (>1000 lines): review file-by-file, summarize each file first, then provide grouped findings. Recommend splitting the PR if it spans unrelated concerns.

---

## Output format (strict)

### 1) Summary
- 3–6 bullets: what the change does and what parts of the system it touches.
- Diff size: small / medium / large (with line count).

### 2) Completeness Checklist (most important)
Goal: catch "I forgot to change X" issues.

Check for missing updates in these categories (only apply what's relevant):
- Public interfaces / API contracts updated (and versioning/back-compat handled)
- Call sites updated (search for old function names/fields)
- Config / feature flags wired through everywhere needed
- DB/schema changes + migrations + backfills handled
- Validation & error handling added where needed
- Permissions/authz/authn updated (if relevant)
- Serialization/deserialization changes accounted for
- Frontend/client/SDK updates if the API/data shape changed
- Docs/README/changelog updates if user-facing behavior changed
- Metrics/logging/tracing added/updated (and not too noisy)
- Tests added/updated (and no flaky gaps)
- Build/CI updated if tooling or paths changed

For each item:
- if done (with pointer to file/area)
- if missing (with exact fix)

### 3) Security Review
Check the diff for:
- Injection risks (SQL, command, XSS, path traversal)
- Missing auth/authz checks on new endpoints
- Sensitive data in logs or error messages
- Hardcoded secrets or credentials
- Missing input validation at system boundaries

If changes touch auth, user input, APIs, or sensitive data: be thorough.
If changes are purely internal logic: a brief "no security concerns" is fine.

### 4) Test Review
- Are tests present when applicable?
- Do they follow the repo's usual approach (no unusual test types)?
- Are there important cases missing?
- Are tests readable and maintainable?

Output:
- What's good
- What's missing
- Minimal changes to improve tests (aligned with repo style)

### 5) Code Quality & Coherence
Focus on:
- readability and structure
- naming
- duplication vs abstraction
- error paths and edge cases
- consistency with existing module conventions

Provide:
- Blockers
- Suggestions
- Nits

### 6) Optional Improvements (only after the above)
- Performance: obvious inefficiencies, unnecessary work
- Maintainability: simpler structure, better factoring, clearer docs/comments

### 7) Final Verdict
Choose one:
- **Request changes**
- **Approve with follow-ups**
- **Approve**

Include top 1–3 reasons.

### 8) Pantry Reminder

After the verdict, remind the user to update the pantry if the review surfaced durable knowledge:
- Architecture decisions or trade-offs discussed
- Patterns or anti-patterns identified
- Requirements clarified or refined

Output:
```
Pantry: If this review surfaced decisions or learnings, run /export --<folder> to update your library.
```
