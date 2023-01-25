package bubbletea

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	left, right := m.viewportLeft.View(), m.viewportRight.View()

	return lipgloss.JoinVertical(lipgloss.Top,
		lipgloss.JoinHorizontal(lipgloss.Top, left, right),
		m.status.View(),
	)
}
