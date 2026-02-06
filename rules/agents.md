# Agent Routing Rules

## When to Auto-Invoke Agents

The following agents should be invoked automatically based on the task context:

- **planner**: Complex feature requests, multi-file changes, tasks requiring exploration before implementation.
- **reviewer**: After code changes when review is needed, or when asked to review a PR.
- **tdd-guide**: Bug fixes, test backfilling, any task involving test creation.
- **debugger**: When investigating bugs, unexpected behavior, or failures.
- **security-reviewer**: Changes touching authentication, authorization, user input handling, API endpoints, or sensitive data.
- **build-resolver**: When the build is failing.

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

Each agent receives:
1. The user's original request.
2. All previous handoff documents in the chain.
3. Clear instructions on what THIS agent should focus on.

## Context Hygiene

- Each agent spawned via `Task()` gets a fresh context window. This is intentional.
- Inline all necessary context into the agent prompt — do not assume agents can read files from a previous agent's context.
- Keep orchestrator context lean; delegate heavy work to agents.
