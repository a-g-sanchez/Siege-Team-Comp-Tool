package initializers

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetTheme() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorBlack,
		ContrastBackgroundColor:     tcell.ColorDarkBlue,
		MoreContrastBackgroundColor: tcell.ColorGreen,
		BorderColor:                 tcell.ColorWhite,
		TitleColor:                  tcell.ColorWhite,
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorGhostWhite,
		SecondaryTextColor:          tcell.ColorSalmon,
		TertiaryTextColor:           tcell.ColorGreen,
		InverseTextColor:            tcell.ColorDarkSlateBlue,
		ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
	}
}
