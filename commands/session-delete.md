---
description: Delete a saved session
argument-hint: --<label>
allowed-tools: [Bash]
---

Find the session file matching `*-<label>.md` in `.planning/sessions/`. If not found, say "No session found matching '<label>'." If found, show the filename and ask for confirmation, then `rm` it.
