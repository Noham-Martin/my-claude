package claude

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateWorktree(projectDir, branchName string) (string, error) {
	repoName := filepath.Base(projectDir)
	sanitized := sanitizeBranch(branchName)

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	worktreePath := filepath.Join(home, ".forktree", repoName, sanitized)

	if err := os.MkdirAll(filepath.Dir(worktreePath), 0755); err != nil {
		return "", fmt.Errorf("failed to create worktree parent dir: %w", err)
	}

	// Create branch
	cmd := exec.Command("git", "branch", branchName)
	cmd.Dir = projectDir
	if out, err := cmd.CombinedOutput(); err != nil {
		// Branch may already exist, that's fine
		if !strings.Contains(string(out), "already exists") {
			return "", fmt.Errorf("failed to create branch: %s", string(out))
		}
	}

	// Reuse existing worktree if the directory already exists
	if _, err := os.Stat(worktreePath); err == nil {
		return worktreePath, nil
	}

	// Create worktree
	cmd = exec.Command("git", "worktree", "add", worktreePath, branchName)
	cmd.Dir = projectDir
	if out, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("failed to create worktree: %s", string(out))
	}

	return worktreePath, nil
}

func RemoveWorktree(projectDir, worktreePath string) error {
	cmd := exec.Command("git", "worktree", "remove", "--force", worktreePath)
	cmd.Dir = projectDir
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to remove worktree: %s", string(out))
	}
	return nil
}

func GetCurrentBranch(dir string) string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func GetGitRoot(dir string) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return dir
	}
	return strings.TrimSpace(string(out))
}

func sanitizeBranch(name string) string {
	s := strings.ReplaceAll(name, "/", "-")
	s = strings.ReplaceAll(s, "\\", "-")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, ":", "-")

	var result strings.Builder
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' {
			result.WriteRune(c)
		}
	}
	return result.String()
}
