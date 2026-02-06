# Handoff Template

Use this structure when passing context between agents in an orchestration chain.

```markdown
## Handoff: <from-agent> → <to-agent>

### What was done
- <bullet summary of this agent's work>

### Findings
- <key findings, issues discovered, decisions made>

### Modified Files
- <files created, edited, or that need attention>

### Unresolved Questions
- <anything this agent could not resolve>

### Recommendations for Next Agent
- <specific guidance for what the next agent should focus on>
```

## Rules

- Always fill every section. Use "None" if a section has nothing to report.
- Keep findings actionable — say what needs to happen, not just what was observed.
- List modified files with their full paths.
- Recommendations should be specific to the next agent's role.
