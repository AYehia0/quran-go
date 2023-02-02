package theme

import (
	"fmt"
	"strings"

	"github.com/AYehia0/quran-go/pkg/quran"
	"github.com/charmbracelet/lipgloss"
)

// Theme represents the properties that make up a theme.
type Theme struct {
	SelectedTreeItemColor                lipgloss.AdaptiveColor
	UnselectedTreeItemColor              lipgloss.AdaptiveColor
	ActiveBoxBorderColor                 lipgloss.AdaptiveColor
	InactiveBoxBorderColor               lipgloss.AdaptiveColor
	StatusBarSelectedFileForegroundColor lipgloss.AdaptiveColor
	StatusBarSelectedFileBackgroundColor lipgloss.AdaptiveColor
	StatusBarBarForegroundColor          lipgloss.AdaptiveColor
	StatusBarBarBackgroundColor          lipgloss.AdaptiveColor
	StatusBarTotalFilesForegroundColor   lipgloss.AdaptiveColor
	StatusBarTotalFilesBackgroundColor   lipgloss.AdaptiveColor
	StatusBarLogoForegroundColor         lipgloss.AdaptiveColor
	StatusBarLogoBackgroundColor         lipgloss.AdaptiveColor
	TitleBackgroundColor                 lipgloss.AdaptiveColor
	TitleForegroundColor                 lipgloss.AdaptiveColor
}

var themeMap = map[string]Theme{
	"default": {
		SelectedTreeItemColor:                lipgloss.AdaptiveColor{Dark: "63", Light: "63"},
		UnselectedTreeItemColor:              lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"},
		ActiveBoxBorderColor:                 lipgloss.AdaptiveColor{Dark: "#F25D94", Light: "#F25D94"},
		InactiveBoxBorderColor:               lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#000000"},
		StatusBarSelectedFileForegroundColor: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
		StatusBarSelectedFileBackgroundColor: lipgloss.AdaptiveColor{Dark: "#F25D94", Light: "#F25D94"},
		StatusBarBarForegroundColor:          lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
		StatusBarBarBackgroundColor:          lipgloss.AdaptiveColor{Dark: "#3c3836", Light: "#3c3836"},
		StatusBarTotalFilesForegroundColor:   lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
		StatusBarTotalFilesBackgroundColor:   lipgloss.AdaptiveColor{Dark: "#A550DF", Light: "#A550DF"},
		StatusBarLogoForegroundColor:         lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
		StatusBarLogoBackgroundColor:         lipgloss.AdaptiveColor{Dark: "#6124DF", Light: "#6124DF"},
		TitleBackgroundColor:                 lipgloss.AdaptiveColor{Dark: "63", Light: "63"},
		TitleForegroundColor:                 lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
	},
}

// GetTheme returns a theme based on the given name.
func GetTheme(theme string) Theme {
	return themeMap["default"]
}

func stringInArray(s string, strings []string) bool {
	for _, str := range strings {
		if str == s {
			return true
		}
	}
	return false
}

// all the names in the page, MAX I think is 3
func getSurahNamesInPage(page []quran.Ayah, lang int) string {
	names := make([]string, 0)
	for _, ayah := range page {
		name := ""
		if lang == 1 {
			name = ayah.NameEn
		} else {
			name = ayah.NameAr
		}

		if !stringInArray(name, names) {
			names = append(names, name)
		}
	}

	return strings.Join(names, "|")
}

func GetTitle(page []quran.Ayah) string {
	return fmt.Sprintf("P%d - %d - %s", page[0].Page, page[0].Juz, getSurahNamesInPage(page, 1))
}
