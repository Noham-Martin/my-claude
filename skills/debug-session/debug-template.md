# Debug File Template

Use this structure when creating a new debug session file at `.planning/debug/debug-YYYY-MM-DD-<slug>.md`.

```markdown
# Debug: <short description>
Status: gathering
Started: YYYY-MM-DD

## Symptoms (immutable — do not edit after initial write)
- <what is failing>
- <error messages, stack traces>
- <expected vs observed behavior>
- <reproduction steps>

## Hypotheses
1. [ACTIVE] <most likely cause>
   - Evidence needed: <what would confirm or eliminate>
2. [ACTIVE] <second possibility>
   - Evidence needed: <what would confirm or eliminate>
3. [ACTIVE] <less likely>
   - Evidence needed: <what would confirm or eliminate>

## Evidence (append-only)
### YYYY-MM-DD HH:MM — <what was checked>
- Finding: <what was observed>
- Implication: <what this means for the hypotheses>

## Resolution
- Root cause: <to be filled when resolved>
- Fix applied: <to be filled>
- Verified: <yes/no>
- Related tests: <tests that cover this>
```

## Status Flow

`gathering` → `investigating` → `fixing` → `verifying` → `resolved`
