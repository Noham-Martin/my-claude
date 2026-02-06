---
description: Incrementally fix build errors one at a time
allowed-tools: [Read, Glob, Grep, Bash, Edit]
---

# /build-fix — Incrementally fix build errors

Fix build errors one at a time, verifying each fix before moving to the next.

## Build System Detection

Before running anything, detect the build system:
1. Check for `BUILD` or `BUILD.bazel` files → bazel
2. Check for `package.json` → npm/yarn/pnpm (check lockfiles to determine which)
3. Check for `Cargo.toml` → cargo
4. Check for `go.mod` → go
5. Check for `Makefile` → make
6. Check for `pom.xml` or `build.gradle` → maven/gradle

Use the detected build system for all build commands.

## Process (strict)

### 1) Run build
Execute the appropriate build command and capture the full output.

### 2) Parse error output
- Group errors by file.
- Sort by severity: syntax errors > type errors > import errors > warnings.
- Count total errors.

### 3) For each error (one at a time)
1. **Show context**: Read the file around the error location.
2. **Explain**: What is wrong and why.
3. **Propose fix**: The minimal change to resolve this specific error.
4. **Apply fix**: Edit the file.
5. **Re-run build**: Verify the error is resolved.
6. **Check regressions**: Ensure the fix did not introduce new errors.

### 4) Stop conditions
Stop immediately if:
- A fix introduces MORE errors than it resolves.
- The same error persists after 3 fix attempts.
- 10 errors have been fixed in one session (pause and report).
- The user requests a pause.

### 5) Report

After fixing or stopping:
```
Errors fixed: N
Errors remaining: N
New errors introduced: N (if any)

Fixed:
- [file:line] description of fix

Remaining:
- [file:line] error message

Next steps:
- ...
```

## Rules

- Fix ONE error at a time. Never batch fixes.
- Always re-run the build after each fix to verify.
- Prefer the smallest possible change. Do not refactor while fixing builds.
- If an error is ambiguous, read surrounding code for context before guessing.
- Never suppress warnings or errors with comments/annotations unless that is the correct fix.
