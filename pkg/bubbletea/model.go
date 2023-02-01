package bubbletea

import (
	"log"

	"github.com/AYehia0/quran-go/pkg/config"
	"github.com/AYehia0/quran-go/pkg/page"
	"github.com/AYehia0/quran-go/pkg/quran"
	"github.com/AYehia0/quran-go/pkg/theme"
)

type model struct {
	viewportLeft  page.ViewportModel
	viewportRight page.ViewportModel
	status        page.Statusbar
	currentPage   int
	cursor        int
	selected      int // 0 for nothing selected, 1 for right, 2 for left
	ready         bool
	theme         theme.Theme
	ayaht         *map[int][]quran.Ayah
}

func New(ayaht *map[int][]quran.Ayah, bookmark quran.Bookmark) model {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	configTheme := theme.GetTheme(cfg.Theme.AppTheme)

	// left and right pages as list of ayaht
	l, r := quran.GetPages(*ayaht, bookmark.CurrentPage)

	// create a viewport
	viewportLeft, viewportRight := page.New(
		false,
		cfg.Settings.Borderless,
		theme.GetTitle(l),
		page.TitleColor{
			Background: configTheme.TitleBackgroundColor,
			Foreground: configTheme.TitleForegroundColor,
		},
		configTheme.InactiveBoxBorderColor,
		l,
		"left",
	), page.New(
		false,
		cfg.Settings.Borderless,
		theme.GetTitle(r),
		page.TitleColor{
			Background: configTheme.TitleBackgroundColor,
			Foreground: configTheme.TitleForegroundColor,
		},
		configTheme.InactiveBoxBorderColor,
		r,
		"right",
	)

	statusbarModel := page.NewStatus(
		page.ColorConfig{
			Foreground: configTheme.StatusBarSelectedFileForegroundColor,
			Background: configTheme.StatusBarSelectedFileBackgroundColor,
		},
		page.ColorConfig{
			Foreground: configTheme.StatusBarBarForegroundColor,
			Background: configTheme.StatusBarBarBackgroundColor,
		},
		page.ColorConfig{
			Foreground: configTheme.StatusBarTotalFilesForegroundColor,
			Background: configTheme.StatusBarTotalFilesBackgroundColor,
		},
		page.ColorConfig{
			Foreground: configTheme.StatusBarLogoForegroundColor,
			Background: configTheme.StatusBarLogoBackgroundColor,
		},
	)
	m := model{
		viewportLeft:  viewportLeft,
		viewportRight: viewportRight,
		status:        statusbarModel,
		currentPage:   bookmark.CurrentPage,
		theme:         configTheme,
		ayaht:         ayaht,
	}

	return m
}
