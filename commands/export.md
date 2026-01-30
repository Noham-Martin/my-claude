# /export — Write new pantry context file into dd-pantry folder

You are running the /export command.

## Usage
/export --<folder>

## Target path
/Users/noham.martin/dd/dd-pantry/<folder>

Replace `<folder>` with the value passed after `--`.

## Goals
- Ensure the target folder exists (create it if missing).
- Create a NEW Markdown file in that folder to store additional context.
- Always follow the file numbering scheme:
  - First file: `00-<name-of-file>.md`
  - Next: `01-...`, `02-...`, `03-...`, etc.
- The content should be **very detailed** and useful later.
- Output MUST be Markdown, and the stored file MUST be Markdown.

## Process (must follow in order)

### 1) Resolve folder argument
- Parse the command invocation to extract `<folder>` from `--<folder>`.
- Construct the absolute path:
  `/Users/noham.martin/dd/dd-pantry/<folder>`

### 2) Ensure folder exists
- If the folder does not exist: create it (including parents).
- If it is newly created, treat this as “first time”.

### 3) Determine next file number
- List files in the folder (non-recursive is enough for numbering).
- Identify files that match the pattern:
  `NN-*.md` where NN is a 2-digit number (00–99).
- If none exist: next number = `00`
- Else: next number = max(NN) + 1
- Always format as 2 digits (e.g., 00, 01, 02, ... 10, 11)

### 4) Determine `<name-of-file>`
- Choose a short, descriptive slug based on the context you are exporting.
- Rules for the name:
  - lowercase
  - words separated by hyphens
  - no spaces
  - avoid overly generic names like “notes”
Examples:
- `implementation-plan-feature-x`
- `architecture-decisions-metrics-v2`
- `debugging-session-2026-01-30`
- `integration-checklist-foo-service`

### 5) Create the new file
- Filename:
  `NN-<name-of-file>.md`
- Write it to:
  `/Users/noham.martin/dd/dd-pantry/<folder>/NN-<name-of-file>.md`

### 6) What to write (Markdown only, very detailed)
Write an “export note” that is maximally useful later. Prefer too much detail over too little.

Include these sections (in this order):

# Title (human-friendly)
- Include the feature/topic and the date (YYYY-MM-DD)

## Purpose
- Why this file exists
- What problem it helps solve later

## Current context
- What is happening right now in the project
- What led to the current state

## Decisions
- Bullet list of decisions made
- For each: rationale + alternatives considered (if known)

## Requirements & constraints
- Functional requirements
- Non-functional constraints (performance, compatibility, security, operational constraints)

## Implementation notes
- Key components touched
- Data flow
- Edge cases
- Any tricky parts and how to avoid mistakes

## Testing & verification
- What was tested / what must be tested
- How to run the relevant tests
- Any gotchas or flakiness notes

## Rollout / operations (if relevant)
- Flags, configs, migrations, monitoring
- Rollback plan

## TODO / next steps
- A concrete checklist

## Links / references (optional)
- Paths, file names, PR links (if provided), docs pointers

### 7) Confirmation output (Markdown)
After writing the file, output:
- the absolute path to the created file
- a short summary of what was captured
- the next suggested action

## Style rules
- Output MUST be Markdown.
- The stored file MUST be Markdown.
- Prefer too much detail over too little.
- Do not invent facts; if unsure, label as assumption.