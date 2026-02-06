---
description: Create a Jira epic with tasks from pantry context
argument-hint: --<pantry-folder>
allowed-tools: [Read, Glob, Grep, Bash, WebFetch]
---

# /jira-create — Create Jira Epic from Pantry

Read a pantry folder and create a full Jira epic with tasks.

## Usage
- `/jira-create --auth-service` — create epic from pantry folder

## Arguments

1. **Pantry folder** (required) — the folder in `~/dd/dd-pantry/<folder>/` to use as context.

## Process

### 1) Read pantry context

Read all files in `~/dd/dd-pantry/<folder>/` recursively.

Extract:
- What the project/feature is about (purpose, scope)
- Requirements and constraints
- Implementation plan (if any)
- Tasks, milestones, deliverables
- Decisions made and their rationale

### 2) Design the epic and tasks

Based on the pantry context, draft:

- **Epic**: title and description summarizing the full scope
- **Tasks**: one ticket per meaningful unit of work

For each task, determine the appropriate status:
- Work already completed → "Done"
- Work actively in progress → "In Progress"
- Work not yet started → "Selected for Development"

Present the plan to the user:

```
Epic: "Auth Service — JWT Authentication"
  Backend authentication system with JWT tokens, refresh rotation, and rate limiting.

  Tasks:
    1. "Setup auth middleware" (Done)
    2. "Implement JWT token generation" (Done)
    3. "Add refresh token rotation" (In Progress)
    4. "Add rate limiting to auth endpoints" (Selected for Development)
    5. "Write integration tests for auth flow" (Selected for Development)
    6. "Deploy auth service to staging" (Selected for Development)

  Total: 6 tasks (2 done, 1 in progress, 3 to do)

Proceed? (y/n)
```

### 3) Create the epic

Use the Jira MCP tools to create the epic:
- Type: Epic
- Summary: concise title
- Description: full context from pantry (scope, requirements, decisions)
- Assignee: Noham Martin
- Component: backend-ingestion

### 4) Create the tasks

For each task, create a ticket:
- Type: Story (or Task, match the project's conventions)
- Summary: imperative form, concise
- Description: detailed context from pantry
- Assignee: Noham Martin
- Component: backend-ingestion
- Status: "Selected for Development", "In Progress", or "Done"
- Link to the epic as parent

### 5) Report

```
Created epic PROJ-200: "Auth Service — JWT Authentication"
  Created: 6 tasks (PROJ-201 through PROJ-206)
    2 Done, 1 In Progress, 3 Selected for Development

  Epic: https://yourcompany.atlassian.net/browse/PROJ-200
```

## Ticket Rules

- **Assignee**: always Noham Martin
- **Component**: always `backend-ingestion`
- **Status**: only use "Selected for Development", "In Progress", or "Done"
- **Summary**: imperative form, concise (e.g., "Add JWT refresh token rotation")
- **Description**: include relevant context from the pantry notes
- **Task granularity**: each task should be a meaningful unit of work (not too small, not too large). Aim for tasks that take 1-3 days.

## Rules

- Always show the full plan and ask for confirmation before creating anything.
- Never create duplicate epics. If the user wants to update an existing epic, suggest `/jira-update` instead.
- Order tasks logically (dependencies first, then independent work).
- If the pantry context is insufficient to create meaningful tasks, say so and ask for more context.
