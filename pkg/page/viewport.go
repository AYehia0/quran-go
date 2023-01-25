// create a viewport that could be used to create panes
package page

import (
	"fmt"

	"github.com/01walid/goarabic"
	"github.com/AYehia0/quran-go/pkg/quran"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 1
	keyWidth = 80
)

var (
	bubbleStyle = lipgloss.NewStyle().
			PaddingLeft(0).
			PaddingRight(0).
			BorderStyle(lipgloss.NormalBorder())
	inputStyle             = lipgloss.NewStyle().PaddingTop(0)
	statusMessageInfoStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
	statusMessageErrorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}).
				Render
)

type TitleColor struct {
	Background lipgloss.AdaptiveColor
	Foreground lipgloss.AdaptiveColor
}

// the settings of the viewport
type ViewportModel struct {
	Viewport    viewport.Model
	Entries     []quran.Ayah
	BorderColor lipgloss.AdaptiveColor
	PageTitle   string
	TitleColor  TitleColor
	IsActive    bool
	HasBorders  bool
}

// generateScreen generates the help text based on the title and entries.
// generete the text which is going to be displayed in the viewport based on the title and the content
func generateViewportText(title string, titleColor TitleColor, entries []quran.Ayah, width, height int) string {
	content := ""

	for _, ayah := range entries {
		text := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"}).
			Width(keyWidth).
			Render(goarabic.RemoveTashkeel(ayah.Text))

		row := lipgloss.JoinHorizontal(lipgloss.Top, goarabic.RemoveTashkeel(text))
		content += fmt.Sprintf("%s\n\n", row)
	}

	titleText := lipgloss.NewStyle().Bold(true).
		Background(titleColor.Background).
		Foreground(titleColor.Foreground).
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Italic(true).
		BorderBottom(true).
		BorderTop(false).
		BorderRight(false).
		BorderLeft(false).
		Render(title)

	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Render(lipgloss.JoinVertical(
			lipgloss.Top,
			titleText,
			content,
		))
}

// create a new instance
func New(
	active, borderless bool,
	title string,
	titleColor TitleColor,
	borderColor lipgloss.AdaptiveColor,
	entries []quran.Ayah,
) ViewportModel {
	viewPort := viewport.New(0, 0)
	border := lipgloss.NormalBorder()

	if borderless {
		border = lipgloss.HiddenBorder()
	}

	viewPort.Style = lipgloss.NewStyle().
		// PaddingLeft(padding).
		// PaddingRight(padding).
		Border(border).
		BorderForeground(borderColor)

	viewPort.SetContent(generateViewportText(title, titleColor, entries, 0, 0))

	return ViewportModel{
		Viewport:    viewPort,
		Entries:     entries,
		PageTitle:   title,
		IsActive:    active,
		HasBorders:  borderless,
		BorderColor: borderColor,
		TitleColor:  titleColor,
	}
}

// SetSize sets the size of the help bubble.
func (b *ViewportModel) SetSize(w, h int) {
	b.Viewport.Width = w - b.Viewport.Style.GetHorizontalFrameSize()
	b.Viewport.Height = h - b.Viewport.Style.GetVerticalFrameSize()

	b.Viewport.SetContent(generateViewportText(b.PageTitle, b.TitleColor, b.Entries, b.Viewport.Width, b.Viewport.Height))
}

// SetBorderColor sets the current color of the border.
func (b *ViewportModel) SetBorderColor(color lipgloss.AdaptiveColor) {
	b.BorderColor = color
}

// SetIsActive sets if the bubble is currently active.
func (b *ViewportModel) SetIsActive(active bool) {
	b.IsActive = active
}

// GotoTop jumps to the top of the viewport.
func (b *ViewportModel) GotoTop() {
	b.Viewport.GotoTop()
}

// SetTitleColor sets the color of the title.
func (b *ViewportModel) SetTitleColor(color TitleColor) {
	b.TitleColor = color

	b.Viewport.SetContent(generateViewportText(b.PageTitle, b.TitleColor, b.Entries, b.Viewport.Width, b.Viewport.Height))
}

// SetBorderless sets weather or not to show the border.
func (b *ViewportModel) SetBorderless(borderless bool) {
	b.HasBorders = borderless
}

// Update handles UI interactions with the help bubble.
func (b ViewportModel) Update(msg tea.Msg) (ViewportModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if b.IsActive {
		b.Viewport, cmd = b.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return b, tea.Batch(cmds...)
}

// View returns a string representation of the help bubble.
func (b ViewportModel) View() string {
	border := lipgloss.NormalBorder()

	if b.HasBorders {
		border = lipgloss.HiddenBorder()
	}

	b.Viewport.Style = lipgloss.NewStyle().
		PaddingLeft(padding).
		PaddingRight(padding).
		Border(border).
		BorderForeground(b.BorderColor)

	return bubbleStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			b.Viewport.View(),
			inputStyle.Render(""),
		))
}