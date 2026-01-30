# Testing Rules

## Test-Driven Development (TDD)

When tests are applicable, the agent MUST follow this order strictly:

1. **Write tests first**
   - Define the expected behavior before writing any implementation.
   - Tests must reflect real usage, not hypothetical behavior.

2. **Run tests and verify failure**
   - Confirm that tests fail for the correct reason.
   - Do NOT proceed if tests pass unexpectedly.

3. **Implement the minimal solution**
   - Write only the code required to make the tests pass.
   - Avoid premature optimization or refactoring at this stage.

4. **Run tests and verify success**
   - All newly added tests must pass.
   - Existing tests must not regress.

5. **Improve tests and refactor**
   - Clarify test names, structure, and assertions.
   - Remove duplication and improve readability.
   - Refactor implementation only if it preserves behavior.

6. **Re-run tests after refactor**
   - All tests must still pass.
   - Re-run the full relevant test suite.

---

## Coverage Rules

- Coverage expectations must be **derived from the existing codebase**, not invented.
- Before adding tests, inspect **similar features or modules** and follow:
  - The same level of coverage
  - The same types of tests (unit, integration, e2e)
- **Do NOT introduce new kinds of tests** if they are not already used in similar areas.
- Match existing conventions for:
  - Test structure
  - Assertion style
  - Mocking strategy

The goal is **consistency with the codebase**, not maximal coverage.