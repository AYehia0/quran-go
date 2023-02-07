package page

import (
	"fmt"
	"strings"

	"github.com/AYehia0/quran-go/pkg/quran"
	"github.com/AYehia0/quran-go/pkg/theme"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	bubbleStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			BorderStyle(lipgloss.NormalBorder())
	inputStyle = lipgloss.NewStyle().PaddingTop(1)
)

type ListModel struct {
	List        list.Model
	SearchInput textinput.Model
	width       int
	height      int
	Entries     []quran.Ayah
	BorderColor lipgloss.AdaptiveColor // TODO: Style isn't really needed if bordercolor is used!
	Style       lipgloss.Style
	PageTitle   string
	IsActive    bool
	HasBorders  bool
	delegate    list.DefaultDelegate
}

func (b *ListModel) SetSize(width, height int) {
	horizontal, vertical := bubbleStyle.GetFrameSize()

	b.List.Styles.StatusBar.Width(width - horizontal)
	b.List.SetSize(
		width-horizontal-vertical,
		height-vertical-lipgloss.Height(b.SearchInput.View())-inputStyle.GetVerticalPadding(),
	)
}

// SetIsActive sets if the bubble is currently active.
func (b *ListModel) SetIsActive(active bool) {
	b.IsActive = active
}

// SetBorderColor sets the color of the border.
func (b *ListModel) SetBorderColor(color lipgloss.AdaptiveColor) {
	b.Style = bubbleStyle.Copy().BorderForeground(color)
}

/*
Item section :
  - the item to be shown in the listModel view
*/

type getAyahtListingMsg []list.Item

type item struct {
	title, desc  string
	ayah         quran.Ayah
	surahsInPage string
	ayahNumber   int
	endOfSurah   bool
}

func (i item) Title() string {
	ayahWidth := 80
	text := lipgloss.NewStyle().Width(ayahWidth).Render(fmt.Sprintf("%s", i.ayah.Text))
	return text
}
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

/*
ListModel section :
	- the list works the same as the viewport but it displays list of Items of type AyahItem
*/
// create a new instance
func NewList(
	active, borderless bool,
	title string,
	titleColor TitleColor,
	borderColor, selectedItemColor, titleForegroundColor, titleBackgroundColor lipgloss.AdaptiveColor,
	entries []quran.Ayah,
) ListModel {

	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.Copy().
		Foreground(selectedItemColor).
		BorderLeftForeground(selectedItemColor)
	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedTitle.Copy()

	listModel := list.New([]list.Item{}, listDelegate, 0, 0)
	listModel.Title = title
	listModel.Styles.Title = listModel.Styles.Title.Copy().
		Bold(true).
		Italic(true).
		Background(titleBackgroundColor).
		Foreground(titleForegroundColor)
	listModel.DisableQuitKeybindings()

	input := textinput.NewModel()
	input.Prompt = "❯ "
	input.Placeholder = "Search Ayah ..."
	input.CharLimit = 250
	input.Width = 50

	if borderless {
		bubbleStyle = bubbleStyle.Copy().Border(lipgloss.HiddenBorder())
	} else {
		bubbleStyle = bubbleStyle.Copy().BorderForeground(borderColor)
	}

	return ListModel{
		List:        listModel,
		SearchInput: input,
		Style:       bubbleStyle,
		BorderColor: borderColor,
		PageTitle:   title,
		IsActive:    active,
		HasBorders:  borderless,
		Entries:     entries,
		delegate:    listDelegate,
	}

}

func (b ListModel) View() string {
	var inputView string = ""

	return b.Style.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			b.List.View(),
			inputStyle.Render(inputView),
		))
}

func getListOfAyaht(title string, entries []quran.Ayah) []list.Item {

	var items []list.Item
	var isEnd bool

	for _, ayah := range entries {

		surahs := strings.Split(title, theme.SurahTitleSep)
		if len(surahs) > 1 && ayah.NumberInSurah == ayah.NumberAyaht {
			isEnd = true
		}

		items = append(items, item{
			title:      ayah.Text,
			ayah:       ayah,
			endOfSurah: isEnd,
		})
	}
	return items
}

func getAyahtListingCmd(title string, entries []quran.Ayah) tea.Cmd {
	return func() tea.Msg {
		return getListOfAyaht(title, entries)
	}
}

func (b *ListModel) ShowList() tea.Cmd {
	return getAyahtListingCmd(b.PageTitle, b.Entries)
}

// SetSelectedItemColors sets the foreground of the selected item.
func (b *ListModel) SetSelectedItemColors(foreground lipgloss.AdaptiveColor) {
	b.delegate.Styles.SelectedTitle = b.delegate.Styles.SelectedTitle.Copy().
		Foreground(foreground).
		BorderLeftForeground(foreground)
	b.delegate.Styles.SelectedDesc = b.delegate.Styles.SelectedTitle.Copy()

	b.List.SetDelegate(b.delegate)
}

// get the current page ayaht into a list of Items
func (b ListModel) Init() tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	cmd = getAyahtListingCmd(b.PageTitle, b.Entries)
	cmds = append(cmds, cmd, textinput.Blink)

	return tea.Batch(cmds...)
}

func (b ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.width = msg.Width
		b.height = msg.Height

		ayahtItems := getListOfAyaht(b.PageTitle, b.Entries)
		cmd = b.List.SetItems(ayahtItems)
		cmds = append(cmds, cmd)
	}
	if b.IsActive {
		b.List, cmd = b.List.Update(msg)
		cmds = append(cmds, cmd)
	}

	return b, tea.Batch(cmds...)
}
