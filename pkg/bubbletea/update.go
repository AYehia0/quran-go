// move between pages vim like

package bubbletea

import (
	"fmt"

	"github.com/AYehia0/quran-go/pkg/page"
	"github.com/AYehia0/quran-go/pkg/quran"
	"github.com/AYehia0/quran-go/pkg/theme"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// the main UI update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		halfSize := msg.Width/2 + 4 // FIX: I don't know why 4 is the magical number that perfectly splits the screen but meh
		bubbleHeight := msg.Height - page.Height + 2

		m.viewportRight.SetSize(halfSize, bubbleHeight)
		m.viewportLeft.SetSize(halfSize, bubbleHeight)

		m.status.SetSize(msg.Width)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, QuitKeys):
			return m, tea.Quit
		case key.Matches(msg, SwitchBox):
			m.toggleBox()
		case key.Matches(msg, NextPage):
			m.movePage("next")
		case key.Matches(msg, PreviousPage):
			m.movePage("prev")
		}
	}

	m.updateStatusbar()

	m.viewportRight, cmd = m.viewportRight.Update(msg)
	cmds = append(cmds, cmd)

	m.viewportLeft, cmd = m.viewportLeft.Update(msg)
	cmds = append(cmds, cmd)

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

// move to next/previous pages
func (m *model) movePage(dir string) {

	// clearing all the Entries
	m.viewportLeft.Entries = nil
	m.viewportRight.Entries = nil

	// TODO: More human/natural viewing of ayaht
	m.viewportLeft.GotoTop()
	m.viewportRight.GotoTop()

	if dir == "next" {
		m.currentPage++
	} else {
		m.currentPage--
	}

	l, r := quran.GetPages(*m.ayaht, m.currentPage)
	m.viewportLeft.Entries, m.viewportRight.Entries = l, r

	// update the UI
	m.viewportLeft.PageTitle = theme.GetTitle(l)
	m.viewportRight.PageTitle = theme.GetTitle(r)

	// // TODO: clean that!!
	page.UpdateText(&m.viewportLeft)
	page.UpdateText(&m.viewportRight)
}

// toggleBox toggles between the two boxes.
func (b *model) toggleBox() {
	// if nothing is selected move to the right pane
	b.selected = (b.selected + 1) % 2
	if b.selected == 0 {
		b.viewportLeft.SetIsActive(false)
		b.viewportRight.SetIsActive(false)

		b.viewportLeft.SetBorderColor(b.theme.InactiveBoxBorderColor)
		b.viewportRight.SetBorderColor(b.theme.InactiveBoxBorderColor)

		b.viewportLeft.SetIsActive(true)
		b.viewportRight.SetBorderColor(b.theme.ActiveBoxBorderColor)

	} else {
		b.viewportLeft.SetIsActive(false)
		b.viewportRight.SetIsActive(false)

		b.viewportRight.SetBorderColor(b.theme.InactiveBoxBorderColor)
		b.viewportLeft.SetBorderColor(b.theme.InactiveBoxBorderColor)

		b.viewportRight.SetIsActive(true)
		b.viewportLeft.SetBorderColor(b.theme.ActiveBoxBorderColor)
	}
}
