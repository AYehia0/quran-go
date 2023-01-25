// move between pages vim like

package bubbletea

import (
	"fmt"

	"github.com/AYehia0/quran-go/pkg/page"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// use this when you're processing stuff with
// complicated ANSI escape sequences
const useHighPerformanceRenderer = false

// the main UI update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, QuitKeys) {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		halfSize := msg.Width / 2
		bubbleHeight := msg.Height - page.Height - 2

		m.viewportLeft.SetBorderless(true)
		m.viewportRight.SetBorderless(true)

		m.viewportRight.SetSize(halfSize, bubbleHeight)
		m.viewportLeft.SetSize(halfSize, bubbleHeight)
		m.status.SetSize(msg.Width)

	}

	m.updateStatusbar()
	m.viewportLeft, cmd = m.viewportLeft.Update(msg)
	cmds = append(cmds, cmd)

	m.viewportRight, cmd = m.viewportRight.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (b model) updateStatusbar() {
	logoText := fmt.Sprintf("%s", "FM")
	b.status.SetContent(
		"",
		fmt.Sprintf("%v/%v", "Hi", "Hello"),
		logoText,
		"",
	)
}
