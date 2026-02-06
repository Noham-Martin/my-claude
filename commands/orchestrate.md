---
description: Chain multiple agents in sequence for end-to-end workflows
argument-hint: --<preset> --<prompt>
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
---

# /orchestrate — Chained Agent Workflows

Run a sequence of specialized agents, each passing a standardized handoff document to the next.

## Usage
`/orchestrate --<preset> --<prompt>`

Examples:
- `/orchestrate --feature --add JWT authentication to the API`
- `/orchestrate --bugfix --login returns 500 on empty password`
- `/orchestrate --refactor --split monolithic UserService into separate concerns`
- `/orchestrate --review`

The prompt is passed to the first agent in the chain as its task. If no prompt is provided, ask the user what they want to work on before starting.

## Presets

| Preset | Chain | Use When |
|--------|-------|----------|
| `feature` | planner → **implement** → tdd-guide → reviewer → security-reviewer | Building a new feature end-to-end |
| `bugfix` | debugger → **implement fix** → tdd-guide → reviewer | Fixing a bug with proper test coverage |
| `refactor` | planner → **implement** → tdd-guide → reviewer | Restructuring code safely |
| `review` | reviewer → security-reviewer | Reviewing before merge |

**Bold steps** are implementation pauses — the orchestrator presents the plan/fix from the previous agent, waits for user confirmation, then implements the code before resuming the chain.

Parse the preset from `--<preset>`. If an unrecognized preset is given, list the available ones and ask the user to choose.

---

## Process

### 1) Initialize

- Parse the preset and the prompt.
- If no prompt is provided: ask the user what they want to work on. Do not start the chain without a prompt.
- Announce the chain to the user:
  ```
  Orchestrating: feature
  Prompt: add JWT authentication to the API
  Chain: planner → implement → tdd-guide → reviewer → security-reviewer
  ```

### 2) Execute the chain

For each step in the chain:

#### Agent steps

- For the **first agent**: use the user's original request as input.
- For **subsequent agents**: include the previous handoff document(s) as input context.
- Spawn the agent using `Task()` with the appropriate subagent type.
- Capture the agent's output and format it as a handoff document (see format below).

#### Implementation steps

When the chain reaches an **implement** step:

1. Present the previous agent's output (the plan or the identified fix).
2. Ask the user for confirmation before implementing.
3. Implement the code following the plan/fix from the previous agent.
4. After implementation is done, produce a handoff document summarizing what was implemented (files changed, decisions made).
5. Resume the chain with the next agent.

#### Progress reporting

After each step completes, show status:
```
[1/5] planner      ... DONE
[2/5] implement    ... DONE (12 files changed)
[3/5] tdd-guide    ... DONE
[4/5] reviewer     ... IN PROGRESS
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
