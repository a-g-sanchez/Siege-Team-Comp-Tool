package initializers

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func SetData() {
	file, err := os.Open("R6-SIEGE-TEAM-COMP-BANK.csv")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapPool := []string{
		"bank",
		"coastline",
		"border",
		"stadim",
		"chalet",
		"clubhouse",
	}

	allOps := []string{
		"Glaz",
		"Fuze",
		"IQ",
		"Blitz",
		"Twitch",
		"Montagne",
		"Thermite",
		"Ash",
		"Thatcher",
		"Sledge",
		"Buck",
		"Blackbeard",
		"Capitao",
		"Hibana",
		"Jackel",
		"Ying",
		"Zofia",
		"Dokkaebi",
		"Finka",
		"Lion",
		"Maverick",
		"Nomad",
		"Gridlock",
		"Nokk",
		"Amaru",
		"Kali",
		"Iana",
		"Ace",
		"Zero",
		"Flores",
		"Osa",
		"Sens",
		"Grim",
		"Brava",
		"Ram",
	}

	type mapData map[string]interface{}

	maps := make(mapData)

	var currentMap string
	var currentSite string
	var ops []string
	var foundSite bool

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if slices.Contains(mapPool, line) {
			currentMap = line
			maps["mapName"] = currentMap
		}

		if strings.Contains(line, "Site") {
			tempSlice := strings.Fields(line)
			currentSite = strings.Join(tempSlice[1:], " ")
			foundSite = true
			ops = nil
			continue
		}

		if foundSite && slices.Contains(allOps, line) {
			ops = append(ops, line)

			if len(ops) == 5 {
				maps[currentSite] = ops
				foundSite = false
				currentMap = ""
			}
		}

	}

	// fmt.Println(maps["CEO - Executive Lounge"])
	// if slice, ok := maps["Lockers - CCTV"].([]string); ok {
	// 	for i, op := range slice {
	// 		fmt.Println(i, op)
	// 	}
	// }
}
