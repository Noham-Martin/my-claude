---
name: orchestrate
description: Chain multiple agents in sequence for end-to-end workflows. Use for feature development, bugfix, refactor, or review pipelines.
argument-hint: --<preset>
allowed-tools: [Read, Glob, Grep, Bash, Edit, Write]
user-invocable: false
---

# Orchestration Skill

This skill powers the `/orchestrate` command by providing the orchestration logic and handoff templates.

## Presets

| Preset | Chain |
|--------|-------|
| `feature` | planner → tdd-guide → reviewer → security-reviewer |
| `bugfix` | debugger → tdd-guide → reviewer |
| `refactor` | planner → reviewer → tdd-guide |
| `review` | reviewer → security-reviewer |

## Orchestration Pattern

1. Parse preset → determine agent chain.
2. For each agent: spawn via `Task()` with fresh context, pass handoff from previous agent.
3. Collect results, format as handoff document.
4. After all agents: aggregate into final report with SHIP / NEEDS WORK / BLOCKED verdict.

## Handoff Template

Use the template in `handoff-template.md` for passing context between agents.

## Error Handling

- If an agent fails: log it, skip to next, note the gap in the final report.
- If a critical agent fails: stop the chain and report immediately.
