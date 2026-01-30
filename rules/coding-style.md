# Coding Style Rules

## Consistency Over Preference

The primary goal is to make new code feel like it has **always belonged** in the codebase.

Before writing or modifying code, the agent MUST search the codebase for existing patterns and reuse them.

---

## Pattern Lookup Order

When looking for how to implement something, follow this order strictly:

1. **Closest context**
   - Same reducer, module, feature, or product area
   - Adjacent files in the same directory

2. **Similar functionality**
   - Other reducers or features solving a comparable problem

3. **Repository-wide patterns**
   - Only if no close example exists

Do NOT invent new abstractions, styles, or patterns if an existing one can be reused.

---

## Code Coherence Rules

- Match existing:
  - Naming conventions
  - File structure
  - Function signatures
  - Error handling patterns
  - Data flow and state management style
- Prefer duplication over introducing a new abstraction unless the codebase already abstracts similar logic.
- Avoid stylistic refactors unless explicitly requested.

When in doubt, **mirror what already exists** rather than improving it.