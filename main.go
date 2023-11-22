package main

import (
	"fmt"

	"github.com/a-g-sanchez/siege-team-comp-tool/initializers"
	"github.com/rivo/tview"
)

func init() {
	initializers.SetTheme()
}

func main() {
	stratData := initializers.SetData()

	app := tview.NewApplication()

	title := tview.NewTextView().SetText("Welcome! Select a map")
	list := tview.NewList().ShowSecondaryText(false)

	infoDisplay := tview.NewTextView()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 1, 1, false).
		AddItem(list, 0, 1, false).
		AddItem(infoDisplay, 0, 1, false)

	for i, strat := range stratData {
		list.AddItem(strat.Name, "", rune(i+'0'), nil)
	}

	list.AddItem("exit", "", 'q', func() {
		app.Stop()
	})

	list.SetSelectedFunc(func(idx int, _ string, _ string, _ rune) {
		var str string
		if idx < len(stratData) {
			for _, site := range stratData[idx].Sites {
				str += fmt.Sprintf("%s\n %s\n\n", site.Name, site.Operators)
			}
		}
		infoDisplay.SetText(str)

	})

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
