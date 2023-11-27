package initializers

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type SiteLocation struct {
	Name      string
	Operators []string
}

type MapLocation struct {
	Name  string
	Sites []SiteLocation
}

func SetData() []MapLocation {
	file, err := os.Open("R6-SIEGE-TEAM-COMP.csv")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapPool := []string{
		"Bank",
		"Coastline",
		"Border",
		"Stadim",
		"Chalet",
		"Clubhouse",
		"Theme Park",
		"Nighthaven Labs",
		"Emerald Plains",
		"Consulate",
		"Kafe",
		"Kanal",
		"Oregon",
		"Outback",
		"Skyscraper",
		"Villa",
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

	var currentMap MapLocation
	var currentSite SiteLocation

	var strats []MapLocation

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if slices.Contains(mapPool, line) {
			currentMap = MapLocation{Name: line}
			if currentMap.Name != "" {
				strats = append(strats, currentMap)
			}
		}

		if strings.Contains(line, "Site") {
			tempSlice := strings.Fields(line)
			siteName := strings.Join(tempSlice[1:], " ")
			currentSite = SiteLocation{Name: siteName}

			for i, mapIndex := range strats {
				if mapIndex.Name == currentMap.Name {
					strats[i].Sites = append(strats[i].Sites, currentSite)
				}
			}
		}

		if slices.Contains(allOps, line) {
			currentSite.Operators = append(currentSite.Operators, line)
			for i, mapIndex := range strats {
				if mapIndex.Name == currentMap.Name {
					for j, sites := range mapIndex.Sites {
						if sites.Name == currentSite.Name {
							strats[i].Sites[j].Operators = append(strats[i].Sites[j].Operators, line)
						}
					}
				}

			}
		}

	}

	return strats
}
