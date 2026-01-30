---
description: Code review focused on completeness, quality, and improvements
allowed-tools: [Read, Glob, Grep, Bash]
---

# /code-review — Completeness-first, then quality, then improvements

You are reviewing to ensure:
1) NOTHING is missing (all required modifications done)
2) Code is coherent, readable, and consistent with the project
3) Optional improvements: performance, security, maintainability

## Hard requirements
- Follow rules/testing.md, rules/coding-style.md, rules/git-workflow.md.
- First prioritize completeness over style or optimization.
- Be actionable: each finding must say where + what + how to fix.

---

## Step 0 — Inputs
If a diff/PR is available, review it.
If not, ask for:
- list of changed files OR a patch/diff
But still provide the completeness checklist immediately.

---

## Output format (strict)

### 1) Summary
- 3–6 bullets: what the change does and what parts of the system it touches.

### 2) Completeness Checklist (most important)
Goal: catch “I forgot to change X” issues.

Check for missing updates in these categories (only apply what’s relevant):
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

### 3) Test Review
- Are tests present when applicable?
- Do they follow the repo’s usual approach (no unusual test types)?
- Are there important cases missing?
- Are tests readable and maintainable?

Output:
- What’s good
- What’s missing
- Minimal changes to improve tests (aligned with repo style)

### 4) Code Quality & Coherence
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

### 5) Optional Improvements (only after the above)
- Performance: obvious inefficiencies, unnecessary work
- Security & safety: unsafe logging, injection risks, missing validation, data exposure
- Maintainability: simpler structure, better factoring, clearer docs/comments

### 6) Final Verdict
Choose one:
- **Request changes**
- **Approve with follow-ups**
- **Approve**

Include top 1–3 reasons.
