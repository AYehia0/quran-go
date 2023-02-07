// move between pages vim like

package bubbletea

import (
	"fmt"

	"github.com/AYehia0/quran-go/pkg/page"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// the main UI update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.listRight, cmd = m.listRight.Update(msg)
	cmds = append(cmds, cmd)

	m.listLeft, cmd = m.listLeft.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		halfSize := msg.Width / 2
		bubbleHeight := msg.Height - page.Height

		m.listRight.SetSize(halfSize, bubbleHeight)
		m.listLeft.SetSize(halfSize, bubbleHeight)

		m.status.SetSize(msg.Width)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, QuitKeys):
			return m, tea.Quit
		case key.Matches(msg, SwitchBox):
			m.toggleBox()
		}
	}

	m.updateStatusbar()

	return m, tea.Batch(cmds...)
}

func (b *model) updateStatusbar() {
	logoText := fmt.Sprintf("%s", "QuranTUI")
	b.status.SetContent(
		"",
		fmt.Sprintf("%v", "Search"),
		logoText,
		"",
	)
}

// toggleBox toggles between the two boxes.
func (b *model) toggleBox() {
	// if nothing is selected move to the right pane
	b.selected = (b.selected + 1) % 2
	if b.selected == 0 {
		b.listLeft.SetIsActive(false)
		b.listRight.SetIsActive(false)

		b.listLeft.SetBorderColor(b.theme.InactiveBoxBorderColor)
		b.listRight.SetBorderColor(b.theme.InactiveBoxBorderColor)

		b.listLeft.SetIsActive(true)
		b.listLeft.SetBorderColor(b.theme.ActiveBoxBorderColor)

	} else {
		b.listLeft.SetIsActive(false)
		b.listRight.SetIsActive(false)

		b.listRight.SetBorderColor(b.theme.InactiveBoxBorderColor)
		b.listLeft.SetBorderColor(b.theme.InactiveBoxBorderColor)

		b.listRight.SetIsActive(true)
		b.listRight.SetBorderColor(b.theme.ActiveBoxBorderColor)
	}
}
