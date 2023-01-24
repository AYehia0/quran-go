package bubbletea

import "github.com/charmbracelet/bubbles/key"

var QuitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
)
