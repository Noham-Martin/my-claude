---
description: Extract reusable patterns from current session as skills
allowed-tools: [Read, Glob, Grep, Bash, Write]
---

# /learn - Extract Reusable Patterns

Analyze the current session and extract any patterns worth saving as skills.

## Trigger

Run `/learn` at any point during a session when you've solved a non-trivial problem.

## What to Extract

Look for:

1. **Error Resolution Patterns**
   - What error occurred?
   - What was the root cause?
   - What fixed it?
   - Is this reusable for similar errors?

2. **Debugging Techniques**
   - Non-obvious debugging steps
   - Tool combinations that worked
   - Diagnostic patterns

3. **Workarounds**
   - Library quirks
   - API limitations
   - Version-specific fixes

4. **Project-Specific Patterns**
   - Codebase conventions discovered
   - Architecture decisions made
   - Integration patterns

## Dedup Check

Before drafting a new skill, search existing skills at `~/.claude/skills/learned/` for similar patterns:
- Read the filenames and scan for overlap.
- If a similar skill exists: propose updating it instead of creating a new one.
- If no overlap: proceed with a new skill.

## Output Format

Create a skill file at `~/.claude/skills/learned/[pattern-name].md`:

```markdown
# [Descriptive Pattern Name]

**Extracted:** [Date]
**Context:** [Brief description of when this applies]
**Domain:** [debugging | performance | architecture | testing | integration | workaround | tooling]
**Confidence:** [0.0-1.0 — how broadly applicable is this pattern?]

## Problem
[What problem this solves - be specific]

## Solution
[The pattern/technique/workaround]

## Example
[Code example if applicable]

## Evidence
[What session context led to this extraction — the problem encountered, what was tried, what worked]

## When to Use
[Trigger conditions - what should activate this skill]

## When NOT to Use
[Conditions where this pattern does not apply or could be harmful]
```

## Confidence Scoring

- **0.1–0.3**: One-off fix that might apply elsewhere. Low certainty.
- **0.4–0.6**: Seen this pattern a couple times. Moderate reusability.
- **0.7–0.9**: Well-established pattern. High confidence it applies broadly.
- **1.0**: Reserved. Never assign 1.0 — there are always edge cases.

## Domain Tags

Assign exactly one domain:
- `debugging` — error diagnosis and resolution techniques
- `performance` — optimization patterns
- `architecture` — structural decisions and design patterns
- `testing` — test strategies and testing-related patterns
- `integration` — cross-system or API integration patterns
- `workaround` — library/tool/API limitation workarounds
- `tooling` — build tools, CI/CD, developer experience

## Process

1. Review the session for extractable patterns
2. Identify the most valuable/reusable insight
3. Run the dedup check against existing skills
4. Draft the skill file with domain and confidence
5. Ask user to confirm before saving
6. Save to `~/.claude/skills/learned/`

## Notes

- Don't extract trivial fixes (typos, simple syntax errors)
- Don't extract one-time issues (specific API outages, etc.)
- Focus on patterns that will save time in future sessions
- Keep skills focused - one pattern per skill
