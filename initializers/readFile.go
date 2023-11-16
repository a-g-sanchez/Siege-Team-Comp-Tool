package initializers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func ReadFile() {
	file, err := os.Open("R6-SIEGE-TEAM-COMP-BANK.csv")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	mapPool := []string{"bank", "coastline", "border"}
	currentOps := []string{"Hibana", "Thermite", "Fenrir", "Bandit", "Buck"}

	type siteLocation struct {
		siteName    string
		opererators []string
	}

	var bank siteLocation

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if err == io.EOF {
			fmt.Println(line)
			if slices.Contains(currentOps, line) {
				bank.opererators = append(bank.opererators, line)
			}
			break
		} else if err != nil {
			fmt.Println("Error reading file", err)
			break
		}

		fmt.Println(line)
		if slices.Contains(mapPool, line) {
			bank.siteName = line
		} else if slices.Contains(currentOps, line) {
			bank.opererators = append(bank.opererators, line)
		}
	}

	fmt.Println(bank)
}
