package planning

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type Session struct {
	Filename  string
	Path      string
	Date      string `yaml:"date"`
	Branch    string `yaml:"branch"`
	Label     string `yaml:"label"`
	NextSteps []string
	Body      string
}

func LoadSessions(projectDir string) ([]Session, error) {
	dir := filepath.Join(projectDir, ".planning", "sessions")
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var sessions []Session
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		s, err := parseSession(path)
		if err != nil {
			continue
		}
		s.Filename = e.Name()
		s.Path = path
		sessions = append(sessions, s)
	}

	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].Filename > sessions[j].Filename
	})

	return sessions, nil
}

func DeleteSession(path string) error {
	return os.Remove(path)
}

func parseSession(path string) (Session, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Session{}, err
	}

	content := string(data)
	var s Session

	if strings.HasPrefix(content, "---\n") {
		end := strings.Index(content[4:], "\n---")
		if end != -1 {
			frontmatter := content[4 : 4+end]
			if err := yaml.Unmarshal([]byte(frontmatter), &s); err != nil {
				return Session{}, err
			}
			s.Body = strings.TrimSpace(content[4+end+4:])
		}
	}

	s.NextSteps = extractNextSteps(s.Body)
	return s, nil
}

func extractNextSteps(body string) []string {
	var steps []string
	scanner := bufio.NewScanner(strings.NewReader(body))
	inNextSteps := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if strings.Contains(strings.ToLower(trimmed), "next steps") && strings.HasPrefix(trimmed, "#") {
			inNextSteps = true
			continue
		}

		if inNextSteps {
			if strings.HasPrefix(trimmed, "#") {
				break
			}
			if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
				steps = append(steps, strings.TrimSpace(trimmed[2:]))
			} else if item, ok := parseNumberedItem(trimmed); ok {
				steps = append(steps, item)
			}
		}
	}

	return steps
}

// parseNumberedItem checks if a line matches "N. text" (e.g. "1. foo", "12. bar")
// and returns the text after the dot.
func parseNumberedItem(line string) (string, bool) {
	for i, c := range line {
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '.' && i > 0 {
			return strings.TrimSpace(line[i+1:]), true
		}
		break
	}
	return "", false
}

func (s Session) Summary() string {
	n := len(s.NextSteps)
	if n > 0 {
		return fmt.Sprintf("%d next step(s)", n)
	}
	return ""
}
