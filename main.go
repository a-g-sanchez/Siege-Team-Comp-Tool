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

	// Getting data from the csv file
	stratData := initializers.SetData()

	// Setting up the tview display
	app := tview.NewApplication()

	title := tview.NewTextView().
		SetText("Select a map to view operators")

	list := tview.NewList().
		ShowSecondaryText(false)

	placeholder := tview.NewTextView()

	siteDisplay := tview.NewTextView()

	for i, strat := range stratData {
		list.AddItem(strat.Name, "", rune(i+'0'), nil)
	}

	list.AddItem("Quit", "", 'q', func() {
		app.Stop()
	})

	list.SetSelectedFunc(func(idx int, _ string, _ string, _ rune) {
		var str string
		if idx < len(stratData) {
			for _, site := range stratData[idx].Sites {
				str += fmt.Sprintf("%s\n %s\n\n", site.Name, site.Operators)
			}
		}
		siteDisplay.SetText(str)

	})

	// Building an terminal interface
	listFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 0, 1, false).
		AddItem(list, 0, 3, false)

	infoFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(placeholder, 0, 1, false).
		AddItem(siteDisplay, 0, 3, false)

	mainFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(listFlex, 0, 1, false).
		AddItem(infoFlex, 0, 1, false)

	if err := app.SetRoot(mainFlex, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
