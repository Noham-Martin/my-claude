---
description: Generate detailed end-to-end implementation plan
allowed-tools: [Read, Glob, Grep, Bash]
---

# /plan — Long, detailed, end-to-end implementation plan

You produce a LONG and DETAILED plan that covers everything required for the feature to work end-to-end.
Assume the user wants an exhaustive checklist they can execute step-by-step.

## Principles
- Be comprehensive and explicit. Prefer too much detail over too little.
- Make dependencies visible (what must be done before what).
- Include cross-repo work: clearly label which repository each task belongs to when relevant.
- Include testing/verification/rollout/documentation tasks (not only code changes).
- If you are uncertain about repo boundaries, make your best guess and clearly mark it as an assumption.

## Inputs
Use what the user provided. If information is missing, make reasonable assumptions and list them.
Do NOT block on questions unless absolutely necessary.

---

## Output Structure (strict)

### 1) Goal
- 2–5 lines describing the feature, scope, and “done” criteria.

### 2) Assumptions & Open Questions
- Assumptions you’re making to proceed.
- Open questions that could change implementation.

### 3) Architecture / Data Flow Overview
Explain the end-to-end flow at a high level:
- Entry points (API/UI/CLI/background job)
- Core processing steps
- Data stores / queues / caches involved
- Outputs (API response, DB writes, emitted events, metrics, UI updates)

### 4) Repo Map (where work will happen)
Provide a table mapping tasks to repos, best-effort:

| Repo | What changes are needed | Why / Notes |
|------|--------------------------|------------|
| <repo A> | ... | ... |
| <repo B> | ... | ... |

If only one repo is involved, still include this section.

### 5) Detailed Task Breakdown (ordered checklist)
Give a *numbered* checklist with clear sub-steps.
Each step must include:
- **Objective**
- **Files/areas to touch (best guess)**
- **Exact change to make**
- **Validation for that step** (how to verify locally)

Include (when applicable):
- Config changes
- Feature flags
- Schema migrations
- Permissions / authz changes
- Backfills / data migrations
- CI pipeline changes
- SDK/client changes
- Docs updates

### 6) Testing & Verification Plan (explicit)
If tests are applicable, follow the TDD discipline:
- What tests to write first (RED)
- How to run them and confirm failure
- Implementation steps (GREEN)
- Test improvements/refactors (IMPROVE)
- What full suite to run before shipping
- Coverage expectations: match what similar parts of the codebase do (don’t invent new test types)

Also include:
- Manual test plan (happy path + key edge cases)
- Observability plan: logs/metrics/traces you’ll check

### 7) Rollout Plan
- Local/dev verification
- Staging verification
- Gradual rollout strategy (flags, canary, phased rollout) if relevant
- Monitoring during rollout (what signals, what thresholds)
- Rollback plan

### 8) Final Pre-PR Checklist
A concise checklist of “things people forget”, for example:
- tests added/updated
- docs updated
- config/flags wired
- migrations handled
- metrics/logging added where needed
- backward compatibility verified
- error handling covered

### 9) PR Packaging
- Recommended PR split (if large)
- Suggested PR title(s)
- Draft PR description(s) at a high level (what/why/how tested)