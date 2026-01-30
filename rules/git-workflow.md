# Git Workflow Rules

## Branching

When creating a branch, the name MUST follow this format:

nohamm/<branch-name>

- Always use lowercase and hyphens.
- Branch names should be concise and descriptive.

---

## Commits

- Each commit MUST contain:
  - A **single-line commit message**
  - Starting with a **verb** in the imperative form (e.g. "add", "fix", "refactor", "remove")
- No additional description lines unless explicitly requested.
- The commit MUST NOT include Claude as a co-author.
- The user must always be the **sole commit author**.

---

## Pull Requests

When it is time to push and open a PR:

- Always provide a **clear PR description**, including:
  - What was changed
  - Why it was changed
  - Any relevant context or trade-offs
- Assume enough context is available to write a meaningful description.
- Do NOT leave the PR description empty or minimal.
