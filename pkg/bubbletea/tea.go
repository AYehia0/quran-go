package bubbletea

import (
	"github.com/AYehia0/quran-go/pkg/quran"
	tea "github.com/charmbracelet/bubbletea"
)

func InitModel(ayaht *map[int][]quran.Ayah, bookmark *quran.Bookmark) model {
	return New(ayaht, *bookmark)
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
