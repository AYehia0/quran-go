package bubbletea

import "github.com/charmbracelet/bubbles/key"

var (
	QuitKeys     = key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"))
	SwitchBox    = key.NewBinding(key.WithKeys("tab"))
	NextPage     = key.NewBinding(key.WithKeys("h", "left"))  // go left
	PreviousPage = key.NewBinding(key.WithKeys("l", "right")) // go left
)
