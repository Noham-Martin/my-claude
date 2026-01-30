---
description: Backfill missing tests based on current branch changes
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
---

# /tests — Backfill missing tests from current branch history

You are running the /tests command.

Your goal is to ensure that **everything introduced or modified on the current branch is properly tested**, following the existing testing culture of the codebase.

This command is especially useful when tests were postponed or partially written.

---

## Core Principles

- Completeness first: ensure no new or modified behavior is left untested.
- Follow the existing test strategy of the codebase (see rules/testing.md).
- Do NOT invent new test types or coverage standards.
- Prefer alignment with similar features over idealized testing.
- Always follow TDD discipline, even when tests are written after implementation.

---

## Step 1 — Inspect git history (current branch)

Analyze the git history of the **current branch only**, starting from the divergence point with the base branch.

Identify:
- Files added
- Files modified
- Files deleted (ignore for testing unless behavior moved elsewhere)

Group changes by:
- Feature
- Module
- Responsibility

Output an initial summary:
- Commit range analyzed
- List of impacted files
- High-level description of what behavior was introduced or changed

---

## Step 2 — Identify test gaps

For each relevant change:
- Determine whether it introduces:
  - new behavior
  - modified behavior
  - new edge cases
- Check whether corresponding tests already exist or were added.

Explicitly call out:
- Files/features with **no tests**
- Files/features with **partial coverage**
- Files/features where tests exist but don’t reflect the new behavior

---

## Step 3 — Locate reference test patterns

Before writing any test:
- Look for **similar features** in the codebase.
- Identify:
  - test file structure
  - naming conventions
  - assertion style
  - mocking strategy
  - coverage depth

Use these as the template.
Do NOT introduce new testing styles.

---

## Step 4 — TDD backfill process (strict)

For each missing or incomplete test, follow this exact flow:

### RED
- Write tests that express the expected behavior introduced on this branch.
- Ensure tests would fail against the base branch or without the new code.
- Explicitly state *why* the test should fail.

### GREEN
- Run the tests against the current branch.
- Verify they pass without modifying implementation unless strictly necessary.

### IMPROVE
- Improve test readability and structure.
- Remove duplication.
- Align naming and organization with existing tests.
- Do NOT over-test beyond what similar features do.

Re-run the full relevant test suite after improvements.

---

## Step 5 — Coverage validation

Coverage expectations MUST be derived from the codebase:
- Compare with tests for similar features.
- Match the same:
  - depth
  - breadth
  - granularity

Rules:
- Do NOT add tests the codebase typically does not create.
- Do NOT chase coverage numbers blindly.
- Consistency > maximal coverage.

---

## Step 6 — Output (strict)

Produce the following sections in Markdown:

### Summary of Changes Reviewed
- Commits inspected
- Features/modules impacted

### Missing or Weak Test Areas
- File / feature
- What behavior is currently untested
- Why it matters

### Tests Added
For each test:
- File path
- What behavior it covers
- Which commit/change it validates

### Tests Improved (if any)
- What was changed
- Why it improves clarity or correctness

### Remaining Gaps (if any)
- Explicitly list anything left untested
- Justify why (out of scope, covered indirectly, etc.)

### Final Recommendation
Choose one:
- **Test coverage sufficient**
- **Additional tests required before merge**

Provide 1–3 concrete next steps if work remains.

---

## Safety & Discipline Rules

- Never modify production code unless required to make tests pass.
- Never add speculative tests for behavior not introduced on this branch.
- Never rewrite large existing test suites unless explicitly requested.