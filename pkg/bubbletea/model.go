package bubbletea

import (
	"log"
	"strconv"

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
	ready         bool
	selected      map[int]struct{}
}

func New(ayaht *[][]quran.Surah, bookmark quran.Bookmark) model {

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	theme := theme.GetTheme(cfg.Theme.AppTheme)

	// left and right pages as list of ayaht
	l, r := quran.GetPages(ayaht, bookmark.CurrentPage)

	// create a viewport
	viewportLeft, viewportRight := page.New(
		false,
		cfg.Settings.Borderless,
		"Page - "+strconv.Itoa(l[0].Page),
		page.TitleColor{
			Background: theme.TitleBackgroundColor,
			Foreground: theme.TitleForegroundColor,
		},
		theme.InactiveBoxBorderColor,
		l,
	), page.New(
		false,
		cfg.Settings.Borderless,
		"Page - "+strconv.Itoa(r[0].Page),
		page.TitleColor{
			Background: theme.TitleBackgroundColor,
			Foreground: theme.TitleForegroundColor,
		},
		theme.InactiveBoxBorderColor,
		r,
	)

	statusbarModel := page.NewStatus(
		page.ColorConfig{
			Foreground: theme.StatusBarSelectedFileForegroundColor,
			Background: theme.StatusBarSelectedFileBackgroundColor,
		},
		page.ColorConfig{
			Foreground: theme.StatusBarBarForegroundColor,
			Background: theme.StatusBarBarBackgroundColor,
		},
		page.ColorConfig{
			Foreground: theme.StatusBarTotalFilesForegroundColor,
			Background: theme.StatusBarTotalFilesBackgroundColor,
		},
		page.ColorConfig{
			Foreground: theme.StatusBarLogoForegroundColor,
			Background: theme.StatusBarLogoBackgroundColor,
		},
	)
	m := model{
		viewportLeft:  viewportLeft,
		viewportRight: viewportRight,
		status:        statusbarModel,
		currentPage:   bookmark.CurrentPage,
		selected:      make(map[int]struct{}),
	}

	return m
}
