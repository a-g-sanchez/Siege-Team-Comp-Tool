package initializers

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func SetData() map[string]map[string][]string {
	file, err := os.Open("R6-SIEGE-TEAM-COMP.csv")

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
		"theme park",
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

	strats := make(map[string]map[string][]string)

	var currentMap string
	var currentSite string
	var foundMap, foundSite bool

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if slices.Contains(mapPool, line) {
			currentMap = line
			strats[currentMap] = make(map[string][]string)
			foundMap = true
		}

		if strings.Contains(line, "Site") {
			if foundMap {
				tempSlice := strings.Fields(line)
				currentSite = strings.Join(tempSlice[1:], " ")
				strats[currentMap][currentSite] = nil
				foundSite = true
				continue
			}
		}

		if foundSite && slices.Contains(allOps, line) {
			if len(strats[currentMap][currentSite]) < 5 {
				strats[currentMap][currentSite] = append(strats[currentMap][currentSite], line)
			} else {
				foundSite = false
			}
		}

	}

	return strats
}
