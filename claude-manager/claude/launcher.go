package claude

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// LaunchInTerminal opens a new Terminal.app window running claude.
// The .command script writes its shell PID to a sidecar file.
// cm uses the PID to detect liveness and to close the window.
func LaunchInTerminal(instanceID, dir string, args ...string) error {
	claudeCmd := "claude"
	if len(args) > 0 {
		claudeCmd = "claude " + strings.Join(args, " ")
	}

	pidDir := PidDir()

	script := fmt.Sprintf(`#!/bin/bash
unset CLAUDECODE
mkdir -p %q
echo $$ > %q

cd %s
%s
`, pidDir, filepath.Join(pidDir, instanceID), shellEscape(dir), claudeCmd)

	tmpDir := filepath.Join(os.TempDir(), "claude-manager")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}

	displayName := filepath.Base(dir)
	scriptPath := filepath.Join(tmpDir, fmt.Sprintf("%s—%s.command", displayName, instanceID))
	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		return fmt.Errorf("failed to write launch script: %w", err)
	}

	cmd := exec.Command("open", scriptPath)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to open terminal: %w\n%s", err, string(out))
	}

	return nil
}

// CloseTerminalWindow finds the Terminal.app window associated with the given
// instance ID (by matching the .command filename in the window title) and
// closes it via System Events.
func CloseTerminalWindow(instanceID string) {
	needle := instanceID
	debugPath := filepath.Join(os.TempDir(), "claude-manager", "close-debug.log")

	// Log that we were called
	_ = os.WriteFile(debugPath, []byte(fmt.Sprintf("called with: %s\n", needle)), 0644)

	script := fmt.Sprintf(`
tell application "System Events"
	tell process "Terminal"
		repeat with w in every window
			if name of w contains %q then
				perform action "AXPress" of button 1 of w
				set maxWait to 20
				repeat maxWait times
					try
						if (count of sheets of w) > 0 then exit repeat
					end try
					delay 0.1
				end repeat
				try
					click button "Terminate" of sheet 1 of w
				end try
				return "closed"
			end if
		end repeat
	end tell
end tell
return "not found"`, needle)

	cmd := exec.Command("osascript", "-e", script)
	out, err := cmd.CombinedOutput()

	// Append result
	f, _ := os.OpenFile(debugPath, os.O_APPEND|os.O_WRONLY, 0644)
	if f != nil {
		fmt.Fprintf(f, "result: %s\nerr: %v\n", string(out), err)
		f.Close()
	}
}

// FocusTerminalWindow finds the Terminal.app window associated with the given
// instance ID (by matching the .command filename in the window name via System
// Events) and brings it to the front.
func FocusTerminalWindow(instanceID string) bool {
	needle := instanceID

	script := fmt.Sprintf(`
tell application "Terminal" to activate
tell application "System Events"
	tell process "Terminal"
		repeat with w in every window
			if name of w contains %q then
				perform action "AXRaise" of w
				return "focused"
			end if
		end repeat
	end tell
end tell
return "not found"`, needle)

	cmd := exec.Command("osascript", "-e", script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.TrimSpace(string(out)) == "focused"
}

func shellEscape(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}
