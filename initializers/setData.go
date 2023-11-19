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
		"bank",
		"coastline",
		"border",
		"stadim",
		"chalet",
		"clubhouse",
		"theme park",
	}

	// allOps := []string{
	// 	"Glaz",
	// 	"Fuze",
	// 	"IQ",
	// 	"Blitz",
	// 	"Twitch",
	// 	"Montagne",
	// 	"Thermite",
	// 	"Ash",
	// 	"Thatcher",
	// 	"Sledge",
	// 	"Buck",
	// 	"Blackbeard",
	// 	"Capitao",
	// 	"Hibana",
	// 	"Jackel",
	// 	"Ying",
	// 	"Zofia",
	// 	"Dokkaebi",
	// 	"Finka",
	// 	"Lion",
	// 	"Maverick",
	// 	"Nomad",
	// 	"Gridlock",
	// 	"Nokk",
	// 	"Amaru",
	// 	"Kali",
	// 	"Iana",
	// 	"Ace",
	// 	"Zero",
	// 	"Flores",
	// 	"Osa",
	// 	"Sens",
	// 	"Grim",
	// 	"Brava",
	// 	"Ram",
	// }

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

	}

	fmt.Println(strats)

	return strats
}
