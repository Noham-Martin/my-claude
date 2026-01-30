# /import — Load pantry context from dd-pantry folder

You are running the /import command.

## Usage
/import --<folder>

## Target path
/Users/noham.martin/dd/dd-pantry/<folder>

Replace `<folder>` with the value passed after `--`.

## Goals
- Ensure the target folder exists (create it if missing).
- Read **every file** in that folder (recursively).
- Build a thorough, structured understanding of the context.
- Output a **detailed Markdown digest** that the user can use as “loaded context”.

## Process (must follow in order)

### 1) Resolve folder argument
- Parse the command invocation to extract `<folder>` from `--<folder>`.
- Construct the absolute path:
  `/Users/noham.martin/dd/dd-pantry/<folder>`

### 2) Ensure folder exists
- If the folder does not exist: create it (including parents).
- If it is newly created: note that this is the first import and there may be no context yet.

### 3) Inventory files
- List all files under the folder **recursively**.
- Prefer Markdown files, but still read any text-like files (txt, md, json, yaml, etc.) if present.
- Ignore obvious binary files (images, archives) unless the user explicitly wants them.

### 4) Read all files (recursively)
- Read the content of each file.
- Keep track of:
  - filename
  - relative path
  - short summary
  - key facts (decisions, constraints, conventions, TODOs)
  - any open questions or contradictions between files

### 5) Produce the “Loaded Context Digest” (Markdown output only)
Always output in Markdown and include:

#### A) Folder status
- Path
- Exists? (created or already existed)
- Number of files read

#### B) File index
A table:
| # | File | Purpose (1 line) | Key topics |
|---|------|-------------------|-----------|

#### C) Consolidated context
- What this folder is about (purpose)
- Current state
- Requirements & constraints
- Known decisions
- Known risks
- Dependencies / external systems
- Important conventions (naming, testing, patterns)
- Things that must not change

#### D) Actionable next steps
- Clear checklist of next actions inferred from the notes

#### E) Open questions
- List questions that, if answered, would remove ambiguity

## Style rules
- Output MUST be Markdown.
- Prefer too much detail over too little.
- Do not hallucinate file contents; only use what you read.
- If no files exist, output a “No existing context yet” section plus a suggested starter structure.