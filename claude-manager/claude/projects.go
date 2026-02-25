package claude

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func projectsPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude-manager", "projects.json")
}

func LoadProjects() ([]Project, error) {
	data, err := os.ReadFile(projectsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var projects []Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
