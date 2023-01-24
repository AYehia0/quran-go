package main

import (
	"fmt"
	"os"

	"github.com/AYehia0/quran-go/pkg/bubbletea"
	"github.com/AYehia0/quran-go/pkg/quran"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	bookmarkPath = "bookmark.json" // the path to the bookmark file
	quranData    = "quran.json"    // the path to the json file containing all the ayaht
)

func main() {
	// check the bookmark if it doesn't exist create a default one
	// if it exists do nothing
	quran.MakeBookmark(bookmarkPath)

	bookmark, err := quran.ReadBookmark(bookmarkPath)
	if err != nil {
		fmt.Printf("Something went wrong working with the bookmark : %v\n", err)
	}

	// read the surahs
	ayaht := quran.ParseQuranData(quranData)

	p := tea.NewProgram(
		bubbletea.InitModel(ayaht, bookmark),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
