/skill-create — Create a reusable skill from pantry context

Usage:
/skill-create --<folder>

<folder> refers to:
/Users/noham.martin/dd/dd-pantry/<folder>

This command turns accumulated knowledge about a feature / task / investigation
into a reusable Claude skill.

Core idea:
- Pantry folder = what you learned while doing the work
- Skill = what you want to reuse in the future

Do NOT save anything without user confirmation.
Output is plain text (CLI-style), not Markdown rendering.

Steps (must follow in order):

1) Resolve pantry folder
- Extract <folder> from the command flag (--<folder>)
- Compute path: /Users/noham.martin/dd/dd-pantry/<folder>
- If the folder does not exist:
  - Create it
  - Mark this as "first time" (likely little or no context yet)

2) Import pantry context
- Read all files in the folder (recursively)
- Treat Markdown files as primary sources
- Build an internal summary of:
  - what problem was worked on
  - decisions made
  - constraints discovered
  - mistakes / pitfalls
  - TODOs and follow-ups
- If the folder is empty, explicitly state that context is minimal

3) Identify reusable patterns
Based on the imported context, look for ONE strong reusable pattern such as:
- an API integration flow you might redo in another project
- a debugging sequence that saved time
- a workaround for a known limitation
- a repeatable testing or rollout approach
- a cross-repo coordination pattern

Rules:
- Prefer depth over breadth
- One run = one skill
- Skip trivial or one-off fixes

4) Draft the skill (verbose by default)

Create a single skill draft with this structure (plain text):

---
name: <descriptive-slug>
description: Reusable pattern extracted from pantry "<folder>"
version: 1.0.0
source: pantry-based-skill-create
---

Title: <Human-friendly title>

Context:
When and why this pattern applies.

Problem:
What recurring problem this solves.

Solution:
Step-by-step description of what to do.
Include commands, code snippets, config notes when relevant.

Key decisions:
Important choices and why they were made.

Testing / verification:
How to validate that this pattern is implemented correctly.

Pitfalls:
What usually goes wrong.
What to double-check.

When to use:
Clear trigger conditions for applying this skill again.

Assumptions:
Explicit list of assumptions or unknowns.

5) Show draft and ask for confirmation
- Print the full draft to the CLI
- Propose a filename:
  ~/.claude/skills/learned/<descriptive-slug>.md
- Ask:
  "Save this skill? (yes / no / edit)"

6) Save (only if confirmed)
- If yes: write the file and confirm path
- If edit: ask what section to modify
- If no: do nothing

Safety rules:
- Never auto-save
- Never create multiple skills in one run
- Never invent facts not present in pantry context (mark assumptions clearly)