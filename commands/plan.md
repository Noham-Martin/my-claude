---
description: Generate detailed end-to-end implementation plan
argument-hint: --<task>
allowed-tools: [Read, Glob, Grep, Bash]
model: opus
---

# /plan — Long, detailed, end-to-end implementation plan

You produce a LONG and DETAILED plan that covers everything required for the feature to work end-to-end.
Assume the user wants an exhaustive checklist they can execute step-by-step.

## Usage
/plan --<task>

Replace `<task>` with a description of what to plan (e.g., `--add user authentication`, `--refactor checkout flow`).

## Principles
- Be comprehensive and explicit. Prefer too much detail over too little.
- Make dependencies visible (what must be done before what).
- Include cross-repo work: clearly label which repository each task belongs to when relevant.
- Include testing/verification/rollout/documentation tasks (not only code changes).
- If you are uncertain about repo boundaries, make your best guess and clearly mark it as an assumption.
- **Context budget**: plans should be executable within ~50% of context window. If the plan is very large, recommend splitting execution across multiple sessions.

## Inputs
Parse the `--<task>` argument to understand what the user wants planned.
If information is missing, make reasonable assumptions and list them.
Do NOT block on questions unless absolutely necessary.

## Before Planning

1. **Explore first**: Read relevant code, configs, and tests to understand the current state.
2. **Find patterns**: Search for how similar features were implemented. Follow the pattern lookup order:
   - Closest context (same module, adjacent files)
   - Similar functionality (other features solving a comparable problem)
   - Repository-wide patterns (only if no close example exists)
3. **Do NOT invent new abstractions** if an existing one can be reused.

---

## Output Structure (strict)

### 1) Goal
- 2–5 lines describing the feature, scope, and "done" criteria.

### 2) Assumptions & Open Questions
- Assumptions you're making to proceed.
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

**Each task gets a REQ-ID** for traceability: `[REQ-01]`, `[REQ-02]`, etc. These IDs link back to the Goal and can be referenced during code review and verification.

**Each task gets a wave number** for dependency-aware ordering: `[Wave 1]`, `[Wave 2]`, etc. Tasks in the same wave can be executed in parallel. Tasks in later waves depend on earlier waves completing.

Each step must include:
- **REQ-ID** and **Wave**
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
- Coverage expectations: match what similar parts of the codebase do (don't invent new test types)

Also include:
- Manual test plan (happy path + key edge cases)
- Observability plan: logs/metrics/traces you'll check

### 7) Rollout Plan
- Local/dev verification
- Staging verification
- Gradual rollout strategy (flags, canary, phased rollout) if relevant
- Monitoring during rollout (what signals, what thresholds)
- Rollback plan

### 8) Final Pre-PR Checklist
A concise checklist of "things people forget", for example:
- tests added/updated
- docs updated
- config/flags wired
- migrations handled
- metrics/logging added where needed
- backward compatibility verified
- error handling covered
- **pantry updated** (`/export --<folder>` with decisions, architecture, requirements)

### 9) PR Packaging
- Recommended PR split (if large) — align splits with wave boundaries when possible
- Suggested PR title(s)
- Draft PR description(s) at a high level (what/why/how tested)
