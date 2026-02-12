# Testing Rules

## Default Behavior

Follow the project's existing testing conventions. Match the same level of coverage, test types, assertion style, and mocking strategy used in similar parts of the codebase.

When adding tests, look at adjacent test files first to understand the expected patterns.

## TDD Discipline

Strict TDD (red-green-refactor) is used when explicitly requested via `/tests` or `/orchestrate`. Outside of those commands, use your judgment — write tests when they add value, skip ceremony when the change is trivial.

## Coverage Rules

- Coverage expectations must be derived from the existing codebase, not invented.
- Do NOT introduce new kinds of tests if they are not already used in similar areas.
