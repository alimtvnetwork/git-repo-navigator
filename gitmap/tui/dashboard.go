package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/model"
)

type dashboardModel struct {
	repos  []model.ScanRecord
	cursor int
}

func newDashboardModel(repos []model.ScanRecord) dashboardModel {
	return dashboardModel{repos: repos}
}

func (m dashboardModel) Update(msg tea.Msg) (dashboardModel, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch {
	case keys.down(keyMsg):
		if m.cursor < len(m.repos)-1 {
			m.cursor++
		}
	case keys.up(keyMsg):
		if m.cursor > 0 {
			m.cursor--
		}
	case keys.refresh(keyMsg):
		// Placeholder — would re-run git status checks
	}

	return m, nil
}

func (m dashboardModel) View() string {
	if len(m.repos) == 0 {
		return styleHint.Render(constants.TUINoRepos)
	}

	var b strings.Builder

	header := fmt.Sprintf("  %-4s %-20s %-12s %-8s",
		"", constants.TUIColSlug, constants.TUIColBranch, constants.TUIColStatus)
	b.WriteString(styleHeader.Render(header))
	b.WriteString("\n")

	for i, r := range m.repos {
		status := styleClean.Render("clean")
		line := fmt.Sprintf("%-20s %-12s %s", r.Slug, r.Branch, status)

		if i == m.cursor {
			b.WriteString(styleCursorRow.Render("> " + line))
		} else {
			b.WriteString(styleNormalRow.Render("  " + line))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styleHint.Render(
		fmt.Sprintf("  %d repositories  •  press r to refresh", len(m.repos))))

	return b.String()
}
