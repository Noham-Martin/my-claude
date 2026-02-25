package tui

import (
	"fmt"
	"strings"

	"github.com/nohamm/claude-manager/planning"
)

type sessionsView struct {
	sessions []planning.Session
	cursor   int
}

func newSessionsView(sessions []planning.Session) sessionsView {
	return sessionsView{sessions: sessions}
}

func reloadSessions(v sessionsView, projectDir string) sessionsView {
	sessions, err := planning.LoadSessions(projectDir)
	if err != nil {
		return v
	}
	v.sessions = sessions
	if v.cursor >= len(v.sessions) {
		v.cursor = max(0, len(v.sessions)-1)
	}
	return v
}

func (v *sessionsView) up() {
	if v.cursor > 0 {
		v.cursor--
	}
}

func (v *sessionsView) down() {
	if v.cursor < len(v.sessions)-1 {
		v.cursor++
	}
}

func (v *sessionsView) selected() *planning.Session {
	if len(v.sessions) == 0 {
		return nil
	}
	return &v.sessions[v.cursor]
}

func (v sessionsView) render(width int) string {
	var b strings.Builder

	b.WriteString(normalStyle.Bold(true).Render("Planning Sessions"))
	b.WriteString(" ")
	b.WriteString(detailStyle.Render("(.planning/sessions/)"))
	b.WriteString("\n")
	b.WriteString(normalStyle.Render(strings.Repeat("─", max(1, min(width-4, 50)))))
	b.WriteString("\n")

	if len(v.sessions) == 0 {
		b.WriteString(detailStyle.Render("  No sessions found"))
		b.WriteString("\n")
		return b.String()
	}

	for i, s := range v.sessions {
		prefix := "  "
		style := normalStyle
		if i == v.cursor {
			prefix = "> "
			style = selectedStyle
		}

		b.WriteString(style.Render(prefix + s.Filename))
		b.WriteString("\n")

		var details []string
		if s.Branch != "" {
			details = append(details, "branch: "+s.Branch)
		}
		if summary := s.Summary(); summary != "" {
			details = append(details, summary)
		}
		if len(details) > 0 {
			b.WriteString(detailStyle.Render(fmt.Sprintf("  %s", strings.Join(details, " | "))))
			b.WriteString("\n")
		}
	}

	return b.String()
}
