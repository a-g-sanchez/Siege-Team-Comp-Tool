package main

import (
	"fmt"

	"github.com/a-g-sanchez/hackathon-project/initializers"
)

func main() {
	stratData := initializers.SetData()

	// display the maps with the map name as the key and is value is a map datatype with the sites as a key and the operators for that site as its value
	// for key, value := range stratData {
	// 	fmt.Println(key, value)
	// }

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
		case "bank":
			fmt.Println(stratData[userInput])
		case "exit":
			loop = false
		default:
			fmt.Println("That is not a valid command")
		}

	}
}
