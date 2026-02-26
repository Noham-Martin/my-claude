package tui

import (
	"fmt"
	"strings"

	"github.com/nohamm/claude-manager/claude"
)

type instancesView struct {
	instances []claude.Instance
	cursor    int
}

func newInstancesView(instances []claude.Instance) instancesView {
	return instancesView{instances: instances}
}

func reloadInstances(v instancesView) instancesView {
	instances, err := claude.LoadInstances()
	if err != nil {
		return v
	}
	v.instances = claude.PruneStale(instances)
	_ = claude.SaveInstances(v.instances)
	if v.cursor >= len(v.instances) {
		v.cursor = max(0, len(v.instances)-1)
	}
	return v
}

func (v *instancesView) up() {
	if v.cursor > 0 {
		v.cursor--
	}
}

func (v *instancesView) down() {
	if v.cursor < len(v.instances)-1 {
		v.cursor++
	}
}

func (v *instancesView) selected() *claude.Instance {
	if len(v.instances) == 0 {
		return nil
	}
	return &v.instances[v.cursor]
}

func (v instancesView) render(width int) string {
	var b strings.Builder

	b.WriteString(normalStyle.Bold(true).Render("Claude Instances"))
	b.WriteString("\n")
	b.WriteString(normalStyle.Render(strings.Repeat("─", max(1, min(width-4, 50)))))
	b.WriteString("\n")

	if len(v.instances) == 0 {
		b.WriteString(detailStyle.Render("  No instances"))
		b.WriteString("\n")
		return b.String()
	}

	for i, inst := range v.instances {
		prefix := "  "
		style := normalStyle
		if i == v.cursor {
			prefix = "> "
			style = selectedStyle
		}

		status := inst.Status()
		statusStr := statusStyle.Render(status)
		if status == "dead" {
			statusStr = errorStyle.Render(status)
		}

		b.WriteString(fmt.Sprintf("%s%s  %s", prefix, style.Render(inst.DisplayName()), statusStr))
		b.WriteString("\n")

		var details []string
		if inst.Branch != "" {
			details = append(details, "branch: "+inst.Branch)
		}
		details = append(details, inst.Age())
		if inst.Worktree != "" {
			details = append(details, "worktree")
		}

		b.WriteString(detailStyle.Render(fmt.Sprintf("  %s", strings.Join(details, " | "))))
		b.WriteString("\n")
	}

	return b.String()
}
