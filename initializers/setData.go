package initializers

import (
	"bufio"
	"fmt"
	"io"
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

	reader := bufio.NewReader(file)

	// mapPool := []string{
	// 	"bank",
	// 	"coastline",
	// 	"border",
	// 	"stadim",
	// 	"chalet",
	// 	"clubhouse",
	// }

	currentOps := []string{
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

	type siteLocation struct {
		siteName    string
		opererators []string
	}

	var bank siteLocation

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err == io.EOF {
			// fmt.Println(line)
			if slices.Contains(currentOps, line) {
				bank.opererators = append(bank.opererators, line)
			}
			break
		} else if err != nil {
			fmt.Println("Error reading file", err)
			break
		}

		if strings.Contains(line, "Site") {
			tempSiteSLice := strings.Fields(line)
			site := strings.Join(tempSiteSLice[1:], " ")
			bank.siteName = site
		} else if slices.Contains(currentOps, line) {
			bank.opererators = append(bank.opererators, line)
		}

	}

	fmt.Println(bank)
}
