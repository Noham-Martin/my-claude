---
description: Quick project verification - checks build, tests, lint, and git status
allowed-tools: [Read, Glob, Grep, Bash]
---

# /verify — Quick project verification (CLI output only)

Goal:
Quickly check whether the current branch is in a good state to be pushed or opened as a PR.

Do NOT format the output as Markdown.
Output should be plain text, readable in Claude CLI.

Steps to perform (in order):

1) Build check
- Identify the usual build command for this project.
- Run the build.
- If the build fails:
  - Report that the build failed.
  - Show the main error.
  - Stop here.

2) Tests related to current branch
- Look at the git diff of the current branch.
- Identify which parts of the code were added or modified.
- Determine which existing tests are related to those changes.
- Run only those tests (do not invent new test types).
- Follow rules/testing.md when evaluating results.
- Report:
  - Which tests were run
  - Whether they passed or failed

3) Lint check
- Run the standard lint command for the project.
- Report whether lint passed.
- If it failed, summarize the main issues.

4) Git status check
Inspect git state and clearly report:
- Unstaged changes (yes / no)
- Staged but uncommitted changes (yes / no)
- Commits not pushed to the remote branch (yes / no)

Do not modify git state.
Do not commit.
Do not push.

Output format (plain text, no Markdown):

Start with a short summary:

Build: PASS / FAIL
Tests: PASS / FAIL
Lint: PASS / FAIL
Git:
- unstaged changes: yes / no
- staged but not committed: yes / no
- commits not pushed: yes / no

If something failed, explain briefly:
- what failed
- why it likely failed
- the exact command the user should run next

End with a clear verdict:

READY TO PUSH
or
NOT READY TO PUSH (fix issues above)

Keep the output concise, factual, and CLI-friendly.