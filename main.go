package main

import (
	"fmt"

	"github.com/a-g-sanchez/hackathon-project/initializers"
)

func main() {
	stratData := initializers.SetData()

	commands := map[string]string{
		"[help]":       "Display available commands",
		"[exit]":       "Exit the application",
		"[maps]":       "Display availabe maps to query",
		"[*map name*]": "Enter a map name and see the sites with operators",
	}

	availableMaps := []string{
		"bank",
		"coastline",
		"border",
		"stadim",
		"chalet",
		"clubhouse",
		"theme park",
	}

	var userInput string

	loop := true
	for loop {

		// Look to change this to use args
		fmt.Scan(&userInput)

		switch userInput {
		case "help":
			for key, val := range commands {
				fmt.Println(key, " - ", val)
			}
		case "maps":
			fmt.Println("Available maps are: ")
			for _, val := range availableMaps {
				fmt.Println(val)
			}
		case "border":
			var selectedMap initializers.MapLocation
			for _, mapChoices := range stratData {
				if mapChoices.Name == userInput {
					selectedMap = mapChoices
				}
			}
			fmt.Println(selectedMap)
		case "exit":
			loop = false
		default:
			fmt.Println("That is not a valid command")
		}

	}
}
