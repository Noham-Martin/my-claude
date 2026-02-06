---
description: Show current status and suggest the next action
allowed-tools: [Read, Glob, Grep, Bash]
---

# /progress — Where am I? What should I do next?

Quick status check that reads all available state and intelligently suggests what to do next.

## Usage
`/progress`

No arguments. This command is read-only.

## Process

### 1) Read available state

Check for and read (if they exist):
- `.planning/STATE.md` — short-term memory
- `.planning/sessions/` — latest session file (most recent by date)
- `.planning/debug/` — any active (non-resolved) debug files

### 2) Check git state

- `git branch --show-current`
- `git status --short`
- `git log --oneline -3`
- `git diff --stat` (uncommitted changes)

### 3) Analyze and route

Based on what is found, determine the current situation and suggest the best next action.

**Decision tree:**

1. **Active debug session exists** (file in `.planning/debug/` with status != resolved):
   → "Active debug session found: <description>. Run `/debug --resume --<name>` to continue." (include the slug from the filename)

2. **Uncommitted changes + tests not run**:
   → "You have uncommitted changes. Run `/verify --quick` before committing."

3. **Uncommitted changes + on main/master branch**:
   → "You have uncommitted changes on main. Create a feature branch first: `git checkout -b nohamm/<branch-name>`"

4. **Commits ahead of remote**:
   → "You have N unpushed commits. Run `/verify --pre-pr` then push."

5. **Session file has unfinished "In Progress" items**:
   → "Previous session had work in progress: <summary>. Pick up where you left off or run `/session-resume` for full context."

6. **STATE.md has "Next Steps"**:
   → "Suggested next steps from last session: <list>"

7. **Clean state, nothing pending**:
   → "All clear. Ready for new work."

### 4) Output (plain text)

```
Branch: <branch>
Status: <clean | N uncommitted changes | N staged>
Recent: <last commit message>

<situation-specific message and suggested action>
```

Keep it short — this is a quick "where am I" check, not a deep analysis.

## Rules

- This command is READ-ONLY. It does not modify files or git state.
- Always check git state, even if no .planning/ files exist.
- Prefer the most actionable suggestion. Don't list everything — pick the ONE most important next step.
- Output is plain text, not Markdown.
