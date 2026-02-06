---
description: Chain multiple agents in sequence for end-to-end workflows
argument-hint: --<preset>
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
---

# /orchestrate — Chained Agent Workflows

Run a sequence of specialized agents, each passing a standardized handoff document to the next.

## Usage
`/orchestrate --<preset>`

## Presets

| Preset | Agent Chain | Use When |
|--------|------------|----------|
| `feature` | planner → tdd-guide → reviewer → security-reviewer | Building a new feature end-to-end |
| `bugfix` | debugger → tdd-guide → reviewer | Fixing a bug with proper test coverage |
| `refactor` | planner → reviewer → tdd-guide | Restructuring code safely |
| `review` | reviewer → security-reviewer | Reviewing before merge |

Parse the preset from `--<preset>`. If an unrecognized preset is given, list the available ones and ask the user to choose.

---

## Process

### 1) Initialize

- Parse the preset and determine the agent chain.
- Announce the chain to the user:
  ```
  Orchestrating: feature
  Chain: planner → tdd-guide → reviewer → security-reviewer
  ```

### 2) Execute each agent in sequence

For each agent in the chain:

#### a) Prepare context
- For the **first agent**: use the user's original request as input.
- For **subsequent agents**: include the previous agent's handoff document as input context.

#### b) Invoke the agent
- Spawn the agent using `Task()` with the appropriate subagent type.
- Pass it:
  - The user's original request
  - The accumulated handoff chain (all previous handoffs)
  - Clear instructions on what THIS agent should focus on

#### c) Collect results
- Capture the agent's output.
- Format it as a handoff document (see format below).

#### d) Report progress
- After each agent completes, show a brief status:
  ```
  [1/4] planner    ... DONE
  [2/4] tdd-guide  ... DONE
  [3/4] reviewer   ... IN PROGRESS
  ```

### 3) Aggregate final report

After all agents complete, produce a combined report:

```
## Orchestration Report: <preset>

### Chain Executed
1. planner — DONE
2. tdd-guide — DONE
3. reviewer — DONE
4. security-reviewer — DONE

### Key Findings
- <aggregated top findings from all agents>

### Modified Files
- <union of all files modified across agents>

### Unresolved Items
- <anything flagged but not resolved>

### Final Verdict
SHIP | NEEDS WORK | BLOCKED

Reason: <1-3 lines>
```

### 4) Update pantry

The pantry (`~/dd/dd-pantry/`) is the persistent knowledge library and MUST be kept current.

After producing the final report:
1. Identify the relevant pantry folder for this project/feature.
2. Export the orchestration results to the pantry using the `/export` logic:
   - Decisions made during the chain
   - Architecture or design choices from the planner
   - Security findings from the security-reviewer
   - Key findings and trade-offs
3. Focus on **durable knowledge** — not build logs or transient test results.
4. If the verdict is SHIP: include a summary of what shipped and why.
5. If NEEDS WORK or BLOCKED: include what is unresolved so the next session can pick it up.

---

## Handoff Document Format

Each agent produces a handoff document for the next agent:

```
## Handoff: <from-agent> → <to-agent>

### What was done
- <bullet summary of this agent's work>

### Findings
- <key findings, issues found, decisions made>

### Modified Files
- <files created, edited, or that need attention>

### Unresolved Questions
- <anything this agent could not resolve>

### Recommendations for Next Agent
- <specific guidance for what the next agent should focus on>
```

---

## Verdict Rules

- **SHIP**: All agents passed. No blockers, no unresolved security issues, tests pass.
- **NEEDS WORK**: Minor issues found. List what needs to be addressed.
- **BLOCKED**: Critical issues found (security vulnerabilities, failing tests, missing implementation). List blockers.

## Error Handling

- If an agent fails or produces no useful output: log the failure, skip to the next agent, and note the gap in the final report.
- If a critical agent fails (e.g., tdd-guide in `bugfix` chain): stop the chain and report immediately.
- Never silently swallow agent failures.

## Rules

- Each agent gets a fresh context via `Task()` — this is intentional for context hygiene.
- Always pass the full handoff chain so later agents have context from earlier ones.
- Do not skip agents in the chain unless one fails critically.
- The orchestrator (this command) does NOT do the work — it only coordinates.
