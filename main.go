package main

import (
	"fmt"

	"github.com/a-g-sanchez/hackathon-project/initializers"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func configStyles() {
	tview.Styles = tview.Theme{
		// PrimitiveBackgroundColor:    tcell.ColorBlack,
		ContrastBackgroundColor:     tcell.ColorDarkBlue,
		MoreContrastBackgroundColor: tcell.ColorGreen,
		BorderColor:                 tcell.ColorWhite,
		TitleColor:                  tcell.ColorWhite,
		GraphicsColor:               tcell.ColorWhite,
		PrimaryTextColor:            tcell.ColorGhostWhite,
		SecondaryTextColor:          tcell.ColorSalmon,
		TertiaryTextColor:           tcell.ColorGreen,
		InverseTextColor:            tcell.ColorDeepSkyBlue,
		ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
	}
}

func init() {
	configStyles()
}

func main() {
	stratData := initializers.SetData()
	// fmt.Println(stratData)

	app := tview.NewApplication()

	title := tview.NewTextView().SetText("Welcome! Select a map")
	list := tview.NewList().ShowSecondaryText(false)

	flex := tview.NewFlex().
		SetDirection(0).
		AddItem(title, 1, 2, false).
		AddItem(list, 0, 1, false)

	for i, strat := range stratData {
		idx := i + 1
		fmt.Println(idx)
		list.AddItem(strat.Name, "", 0, nil)
	}

	list.AddItem("Exit", "", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
