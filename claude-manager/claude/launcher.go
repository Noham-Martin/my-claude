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
	home, _ := os.UserHomeDir()
	pluginDir := filepath.Join(home, ".claude", "plugins", "nohamm-workflow")

	claudeCmd := fmt.Sprintf("claude --plugin-dir %s", shellEscape(pluginDir))
	if len(args) > 0 {
		claudeCmd += " " + strings.Join(args, " ")
	}

	pidDir := PidDir()

	script := fmt.Sprintf(`#!/bin/zsh -l
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
return "not found"`, instanceID)

	cmd := exec.Command("osascript", "-e", script)
	_, _ = cmd.CombinedOutput()
}

// FocusTerminalWindow finds the Terminal.app window associated with the given
// instance ID (by matching the .command filename in the window name via System
// Events) and brings it to the front.
func FocusTerminalWindow(instanceID string) bool {
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
return "not found"`, instanceID)

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
