---
description: Update a Jira epic with current progress from pantry or PRs
argument-hint: --<epic-key> --<pantry-folder | PR list>
allowed-tools: [Read, Glob, Grep, Bash, WebFetch]
---

# /jira-update — Update Jira Epic

Sync a Jira epic with current progress. Creates missing tickets, updates statuses, and ensures the epic reflects reality.

## Usage
- `/jira-update --PROJ-123 --auth-service` — update epic from pantry folder
- `/jira-update --PROJ-123 --#456 #789 #790` — update epic from a list of PRs

## Arguments

1. **Epic key** (required) — the Jira epic to update (e.g., `PROJ-123`)
2. **Context source** (required) — either:
   - A pantry folder name (loads from `~/dd/dd-pantry/<folder>/`)
   - A list of PR numbers prefixed with `#` (fetches via `gh pr view`)

## Jira Access

Use the Atlassian MCP server (configured in dd-source as `atlassian` via `https://mcp.atlassian.com/v1/sse`). All Jira operations go through the MCP tools (search, create, update, transition).

## Process

### 1) Gather context

**If pantry folder:**
- Read all files in `~/dd/dd-pantry/<folder>/` recursively.
- Extract: tasks completed, tasks in progress, decisions made, remaining work.

**If PR list:**
- For each PR number, run `gh pr view <number>` to get title, description, status, and changed files.
- Extract: what was implemented, what is merged vs open vs draft.

### 2) Read current epic state

Use the Jira MCP tools to:
- Fetch the epic and all its child tickets (stories/tasks).
- For each ticket: key, summary, status, assignee, component.

### 3) Diff and plan changes

Compare the gathered context against the current epic state:

- **New work not in Jira** → plan ticket creation
- **Completed work with open tickets** → plan status transition to "Done"
- **Work with open PRs awaiting review** → plan status transition to "In Review"
- **In-progress work** → plan status transition to "In Progress"
- **Planned but not started** → ensure status is "Selected for Development"
- **Tickets that match existing work** → update summary/description if stale

Present the plan to the user before making changes:

```
Epic: PROJ-123 — Auth Service

  Create:
    + "Implement JWT refresh token rotation" (Selected for Development)
    + "Add rate limiting to auth endpoints" (Selected for Development)

  Update status:
    PROJ-124 "Setup auth middleware" → Done (was: In Progress)
    PROJ-125 "Write integration tests" → In Progress (was: Selected for Development)

  No changes:
    PROJ-126 "Deploy to staging" (Selected for Development)

Proceed? (y/n)
```

### 4) Apply changes

After user confirmation:

**Creating tickets:**
- Type: Story (or Task, match existing ticket types in the epic)
- Summary: concise, imperative form
- Description: detailed context from pantry/PRs
- Assignee: Noham Martin
- Component: backend-ingestion
- Status: one of "Selected for Development", "In Progress", "In Review", or "Done"
- Link to the epic as parent

**Updating tickets:**
- Transition status using the Jira workflow transitions
- Update description if new context is available

### 5) Report

```
Updated epic PROJ-123:
  Created: 2 tickets (PROJ-130, PROJ-131)
  Updated: 2 tickets (PROJ-124 → Done, PROJ-125 → In Progress)
  Unchanged: 1 ticket
```

## Ticket Rules

- **Assignee**: always Noham Martin
- **Component**: always `backend-ingestion`
- **Status**: only use "Selected for Development", "In Progress", "In Review", or "Done"
- **Summary**: imperative form, concise (e.g., "Add JWT refresh token rotation")
- **Description**: include relevant context, file paths, PR links where applicable

## Rules

- Always show the plan and ask for confirmation before creating or modifying tickets.
- Never delete tickets. If a ticket seems obsolete, flag it but do not remove it.
- Never change the epic title or description unless explicitly asked.
- If a PR is merged, the corresponding ticket should be "Done".
- If a PR is open and ready for review, the corresponding ticket should be "In Review".
- If a PR is draft or work is ongoing, the corresponding ticket should be "In Progress".
