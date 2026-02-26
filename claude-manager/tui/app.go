package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nohamm/claude-manager/claude"
	"github.com/nohamm/claude-manager/planning"
)

type tab int

const (
	tabInstances tab = iota
	tabSessions
)

type inputMode int

const (
	inputNone inputMode = iota
	inputBranch
	inputConfirmDelete
	inputProjectPicker
)

type pickerIntent int

const (
	intentNewClaude pickerIntent = iota
	intentWorktree
)

type model struct {
	projectDir string
	width      int
	height     int

	activeTab tab
	sessions  sessionsView
	instances instancesView

	input     inputMode
	inputText string
	status    string
	err       string

	pickerIntent pickerIntent
	pickerDir    string

	projects      []claude.Project
	projectCursor int
}

type statusMsg string
type errMsg string
type refreshMsg struct{}

func NewModel(projectDir string) model {
	sessions, _ := planning.LoadSessions(projectDir)
	instances, _ := claude.LoadInstances()
	instances = claude.PruneStale(instances)
	_ = claude.SaveInstances(instances)
	projects, _ := claude.LoadProjects()

	return model{
		projectDir: projectDir,
		sessions:   newSessionsView(sessions),
		instances:  newInstancesView(instances),
		projects:   projects,
		width:      80,
		height:     24,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("claude-manager")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case statusMsg:
		m.status = string(msg)
		m.err = ""
		return m, nil

	case errMsg:
		m.err = string(msg)
		m.status = ""
		return m, nil

	case refreshMsg:
		m.sessions = reloadSessions(m.sessions, m.projectDir)
		m.instances = reloadInstances(m.instances)
		return m, nil

	case tea.KeyMsg:
		if m.input != inputNone {
			return m.handleInput(msg)
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab":
			if m.activeTab == tabSessions {
				m.activeTab = tabInstances
			} else {
				m.activeTab = tabSessions
			}
			// Refresh data when switching tabs
			m.instances = reloadInstances(m.instances)
			m.sessions = reloadSessions(m.sessions, m.projectDir)
			m.status = ""
			m.err = ""
			return m, nil

		case "up", "k":
			m.moveUp()
			return m, nil

		case "down", "j":
			m.moveDown()
			return m, nil

		case "n":
			if m.activeTab == tabInstances {
				m.pickerIntent = intentNewClaude
				m.input = inputProjectPicker
				m.projectCursor = 0
				m.projects, _ = claude.LoadProjects()
				m.status = ""
				m.err = ""
				return m, nil
			}
			return m, nil

		case "w":
			if m.activeTab == tabInstances {
				m.pickerIntent = intentWorktree
				m.input = inputProjectPicker
				m.projectCursor = 0
				m.projects, _ = claude.LoadProjects()
				m.status = ""
				m.err = ""
				return m, nil
			}
			return m, nil

		case "d":
			if m.activeTab == tabSessions {
				if s := m.sessions.selected(); s != nil {
					m.input = inputConfirmDelete
					m.inputText = ""
					return m, nil
				}
			} else if m.activeTab == tabInstances {
				return m.dismissInstance()
			}
			return m, nil

		case "r":
			if m.activeTab == tabSessions {
				if s := m.sessions.selected(); s != nil && s.Label != "" {
					return m.resumeSession(s.Label)
				}
			}
			return m, nil

		case "enter":
			if m.activeTab == tabSessions {
				if s := m.sessions.selected(); s != nil {
					return m.viewSession(s)
				}
			} else if m.activeTab == tabInstances {
				return m.focusInstance()
			}
			return m, nil
		}
	}

	return m, nil
}

func (m model) handleInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.input == inputProjectPicker {
		return m.handleProjectPicker(msg)
	}

	switch msg.String() {
	case "esc":
		m.input = inputNone
		m.inputText = ""
		return m, nil

	case "enter":
		return m.submitInput()

	case "backspace":
		if len(m.inputText) > 0 {
			m.inputText = m.inputText[:len(m.inputText)-1]
		}
		return m, nil

	default:
		if len(msg.String()) == 1 {
			m.inputText += msg.String()
		}
		return m, nil
	}
}

func (m model) submitInput() (tea.Model, tea.Cmd) {
	switch m.input {
	case inputBranch:
		m.input = inputNone
		branch := strings.TrimSpace(m.inputText)
		m.inputText = ""
		if branch == "" {
			return m, nil
		}
		return m.launchWorktree(m.pickerDir, branch)

	case inputConfirmDelete:
		m.input = inputNone
		if strings.ToLower(strings.TrimSpace(m.inputText)) == "y" {
			return m.deleteSelected()
		}
		m.inputText = ""
		return m, nil

	default:
		return m, nil
	}
}

func (m model) handleProjectPicker(msg tea.KeyMsg) (model, tea.Cmd) {
	// Total items: "current directory" + project aliases
	total := 1 + len(m.projects)

	switch msg.String() {
	case "esc":
		m.input = inputNone
		return m, nil

	case "up", "k":
		if m.projectCursor > 0 {
			m.projectCursor--
		}
		return m, nil

	case "down", "j":
		if m.projectCursor < total-1 {
			m.projectCursor++
		}
		return m, nil

	case "enter":
		dir := m.projectDir
		if m.projectCursor > 0 {
			dir = m.projects[m.projectCursor-1].Path
		}

		if m.pickerIntent == intentWorktree {
			m.input = inputBranch
			m.inputText = ""
			m.pickerDir = dir
			return m, nil
		}

		m.input = inputNone
		return m.launchClaude(dir)
	}

	return m, nil
}

func (m *model) moveUp() {
	switch m.activeTab {
	case tabSessions:
		m.sessions.up()
	case tabInstances:
		m.instances.up()
	}
}

func (m *model) moveDown() {
	switch m.activeTab {
	case tabSessions:
		m.sessions.down()
	case tabInstances:
		m.instances.down()
	}
}

func (m model) launchClaude(dir string) (model, tea.Cmd) {
	gitRoot := claude.GetGitRoot(dir)
	branch := claude.GetCurrentBranch(dir)
	id := claude.NewInstanceID()

	if err := claude.LaunchInTerminal(id, dir); err != nil {
		m.err = err.Error()
		return m, nil
	}

	_, _ = claude.AddInstance(id, gitRoot, branch, "")
	m.status = fmt.Sprintf("Launched Claude in %s", filepath.Base(gitRoot))
	m.instances = reloadInstances(m.instances)
	return m, nil
}

func (m model) launchWorktree(dir, branch string) (model, tea.Cmd) {
	gitRoot := claude.GetGitRoot(dir)
	id := claude.NewInstanceID()

	worktreePath, err := claude.CreateWorktree(gitRoot, branch)
	if err != nil {
		m.err = fmt.Sprintf("Worktree error: %s", err)
		return m, nil
	}

	if err := claude.LaunchInTerminal(id, worktreePath); err != nil {
		m.err = err.Error()
		return m, nil
	}

	_, _ = claude.AddInstance(id, gitRoot, branch, worktreePath)
	m.status = fmt.Sprintf("Launched Claude in worktree: %s", branch)
	m.instances = reloadInstances(m.instances)
	return m, nil
}

func (m model) resumeSession(label string) (model, tea.Cmd) {
	dir := m.projectDir
	gitRoot := claude.GetGitRoot(dir)
	branch := claude.GetCurrentBranch(dir)
	id := claude.NewInstanceID()
	prompt := fmt.Sprintf("/session-resume --%s", label)

	if err := claude.LaunchInTerminal(id, dir, "--prompt", fmt.Sprintf("%q", prompt)); err != nil {
		m.err = err.Error()
		return m, nil
	}

	_, _ = claude.AddInstance(id, gitRoot, branch, "")
	m.status = fmt.Sprintf("Resumed session: %s", label)
	m.instances = reloadInstances(m.instances)
	return m, nil
}

func (m model) focusInstance() (model, tea.Cmd) {
	inst := m.instances.selected()
	if inst == nil {
		return m, nil
	}

	if claude.FocusTerminalWindow(inst.ID) {
		m.status = fmt.Sprintf("Focused %s", inst.DisplayName())
	} else {
		m.err = fmt.Sprintf("Could not find terminal for %s", inst.DisplayName())
	}
	return m, nil
}

func (m model) dismissInstance() (model, tea.Cmd) {
	inst := m.instances.selected()
	if inst == nil {
		return m, nil
	}

	name := inst.DisplayName()
	claude.KillInstance(inst.ID)
	m.status = fmt.Sprintf("Killed %s", name)
	m.instances = reloadInstances(m.instances)
	return m, nil
}

func (m model) deleteSelected() (model, tea.Cmd) {
	s := m.sessions.selected()
	if s == nil {
		return m, nil
	}

	if err := planning.DeleteSession(s.Path); err != nil {
		m.err = fmt.Sprintf("Delete failed: %s", err)
		return m, nil
	}

	m.status = fmt.Sprintf("Deleted %s", s.Filename)
	m.sessions = reloadSessions(m.sessions, m.projectDir)
	return m, nil
}

func (m model) viewSession(s *planning.Session) (model, tea.Cmd) {
	lines := strings.Count(s.Body, "\n") + 1
	m.status = fmt.Sprintf("Viewing %s (%d lines) — branch: %s", s.Filename, lines, s.Branch)
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render(" claude-manager "))
	b.WriteString("\n\n")

	instancesTab := inactiveTabStyle.Render("Instances")
	sessionsTab := inactiveTabStyle.Render("Sessions")
	if m.activeTab == tabInstances {
		instancesTab = activeTabStyle.Render("Instances")
	} else {
		sessionsTab = activeTabStyle.Render("Sessions")
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, instancesTab, sessionsTab))
	b.WriteString("\n\n")

	switch m.activeTab {
	case tabInstances:
		b.WriteString(m.instances.render(m.width))
	case tabSessions:
		b.WriteString(m.sessions.render(m.width))
	}

	if m.input != inputNone {
		b.WriteString("\n")
		switch m.input {
		case inputBranch:
			b.WriteString(inputStyle.Render("Branch name: "))
			b.WriteString(m.inputText)
			b.WriteString("█")
		case inputConfirmDelete:
			s := m.sessions.selected()
			name := ""
			if s != nil {
				name = s.Filename
			}
			b.WriteString(inputStyle.Render(fmt.Sprintf("Delete %s? (y/n): ", name)))
			b.WriteString(m.inputText)
			b.WriteString("█")
		case inputProjectPicker:
			b.WriteString(m.renderProjectPicker())
		}
		b.WriteString("\n")
	}

	if m.status != "" {
		b.WriteString("\n")
		b.WriteString(statusStyle.Render(m.status))
		b.WriteString("\n")
	}
	if m.err != "" {
		b.WriteString("\n")
		b.WriteString(errorStyle.Render(m.err))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	if m.activeTab == tabSessions {
		b.WriteString(helpStyle.Render("[r] resume  [d] delete  [enter] view  [tab] switch  [q] quit"))
	} else {
		b.WriteString(helpStyle.Render("[n] new claude  [w] worktree  [enter] focus  [d] kill  [tab] switch  [q] quit"))
	}

	return b.String()
}

func (m model) renderProjectPicker() string {
	var b strings.Builder

	b.WriteString(inputStyle.Render("Select project:"))
	b.WriteString("\n")

	// "Current directory" is always the first option
	prefix := "  "
	style := normalStyle
	if m.projectCursor == 0 {
		prefix = "> "
		style = selectedStyle
	}
	b.WriteString(fmt.Sprintf("  %s%s", prefix, style.Render("current directory")))
	b.WriteString(detailStyle.Render(fmt.Sprintf("  %s", m.projectDir)))
	b.WriteString("\n")

	for i, p := range m.projects {
		prefix = "  "
		style = normalStyle
		if m.projectCursor == i+1 {
			prefix = "> "
			style = selectedStyle
		}
		b.WriteString(fmt.Sprintf("  %s%s", prefix, style.Render(p.Name)))
		b.WriteString(detailStyle.Render(fmt.Sprintf("  %s", p.Path)))
		b.WriteString("\n")
	}

	b.WriteString(helpStyle.Render("  [enter] select  [esc] cancel"))

	return b.String()
}
