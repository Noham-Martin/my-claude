package claude

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
)

type Instance struct {
	ID        string    `json:"id"`
	Project   string    `json:"project"`
	Branch    string    `json:"branch"`
	Worktree  string    `json:"worktree,omitempty"`
	StartedAt time.Time `json:"started_at"`
}

func instancesPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude-manager", "instances.json")
}

func PidDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude-manager", "pids")
}

func LoadInstances() ([]Instance, error) {
	data, err := os.ReadFile(instancesPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var instances []Instance
	if err := json.Unmarshal(data, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}

func SaveInstances(instances []Instance) error {
	path := instancesPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	if instances == nil {
		instances = []Instance{}
	}

	data, err := json.MarshalIndent(instances, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// PruneStale removes instances whose shell process is no longer running.
// It reads the PID from the sidecar file written by the .command script.
func PruneStale(instances []Instance) []Instance {
	var alive []Instance
	for _, inst := range instances {
		pid := readPID(inst.ID)
		if pid == 0 {
			// PID file not yet written — give 5s for the terminal to start
			if time.Since(inst.StartedAt) < 5*time.Second {
				alive = append(alive, inst)
			}
			// Otherwise no PID file and old enough = dead, drop it
			continue
		}
		if isProcessAlive(pid) {
			alive = append(alive, inst)
		} else {
			// Dead — clean up PID file
			os.Remove(filepath.Join(PidDir(), inst.ID))
		}
	}
	return alive
}

func NewInstanceID() string {
	return uuid.New().String()
}

func AddInstance(id, project, branch, worktree string) (Instance, error) {
	instances, err := LoadInstances()
	if err != nil {
		instances = nil
	}

	instances = PruneStale(instances)

	inst := Instance{
		ID:        id,
		Project:   project,
		Branch:    branch,
		Worktree:  worktree,
		StartedAt: time.Now(),
	}

	instances = append(instances, inst)
	if err := SaveInstances(instances); err != nil {
		return inst, err
	}

	return inst, nil
}

func RemoveInstance(id string) error {
	instances, err := LoadInstances()
	if err != nil {
		return err
	}

	var filtered []Instance
	for _, inst := range instances {
		if inst.ID != id {
			filtered = append(filtered, inst)
		}
	}

	os.Remove(filepath.Join(PidDir(), id))
	return SaveInstances(filtered)
}

// KillInstance closes the Terminal window (while the .command name is still
// in the title), kills the process, and removes the instance record.
func KillInstance(id string) {
	// Close the Terminal window FIRST — the window title still contains
	// the .command filename at this point. After the process dies, the
	// title changes and we can no longer find the window.
	CloseTerminalWindow(id)

	// Now kill the process
	pid := readPID(id)
	if pid > 0 {
		if p, err := os.FindProcess(pid); err == nil {
			_ = p.Signal(syscall.SIGKILL)
		}
	}

	// Remove instance record and PID file
	_ = RemoveInstance(id)
}

func readPID(instanceID string) int {
	data, err := os.ReadFile(filepath.Join(PidDir(), instanceID))
	if err != nil {
		return 0
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0
	}
	return pid
}

func isProcessAlive(pid int) bool {
	if pid <= 0 {
		return false
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (inst Instance) DisplayName() string {
	return filepath.Base(inst.Project)
}

func (inst Instance) IsAlive() bool {
	pid := readPID(inst.ID)
	if pid == 0 {
		return time.Since(inst.StartedAt) < 5*time.Second
	}
	return isProcessAlive(pid)
}

func (inst Instance) Status() string {
	if inst.IsAlive() {
		return "running"
	}
	return "dead"
}

func (inst Instance) Age() string {
	elapsed := time.Since(inst.StartedAt).Truncate(time.Minute)
	if elapsed < time.Minute {
		return "just now"
	}
	if elapsed < time.Hour {
		return fmt.Sprintf("%dm ago", int(elapsed.Minutes()))
	}
	return fmt.Sprintf("%dh%dm ago", int(elapsed.Hours()), int(elapsed.Minutes())%60)
}
