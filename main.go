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

	infoDisplay := tview.NewTextView()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 1, 2, false).
		AddItem(list, 0, 1, false).
		AddItem(infoDisplay, 0, 1, false)

	for i, strat := range stratData {
		list.AddItem(strat.Name, "", rune(i+'0'), nil)

		list.SetSelectedFunc(func(idx int, _ string, _ string, _ rune) {
			ind := fmt.Sprintf("%d", idx)
			text := fmt.Sprintf("selected item %s", ind)
			infoDisplay.SetText(text)
		})

	}

	list.AddItem("exit", "", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
