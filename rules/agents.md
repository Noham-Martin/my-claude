# Agent Routing Rules

## Agents Are Opt-In

Agents run only when explicitly invoked via commands (`/orchestrate`, `/debug`, `/code-review`, `/plan`, `/tests`, `/build-fix`). Do NOT auto-spawn agents for regular tasks.

## Model Routing

- Use **opus** for planning and architecture decisions (planner agent).
- Use **sonnet** for everything else (reviewer, tdd-guide, debugger, security-reviewer, build-resolver).

## Parallel Execution

When agents are independent, spawn them concurrently:
- reviewer and security-reviewer can run in parallel (both are read-only).
- planner must complete before tdd-guide or reviewer (sequential dependency).

## Handoff Protocol

When agents pass work to each other (e.g., via `/orchestrate`), use the standardized handoff format:

```
## Handoff: <from-agent> → <to-agent>

### What was done
### Findings
### Modified Files
### Unresolved Questions
### Recommendations for Next Agent
```

## Context Hygiene

- Each agent spawned via `Task()` gets a fresh context window. This is intentional.
- Inline all necessary context into the agent prompt — do not assume agents can read files from a previous agent's context.
- Keep orchestrator context lean; delegate heavy work to agents.
