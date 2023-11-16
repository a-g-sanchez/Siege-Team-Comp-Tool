package main

import (
	"fmt"

	"github.com/a-g-sanchez/hackathon-project/initializers"
)

func main() {
	stratData := initializers.SetData()

	fmt.Println(stratData)
	// fmt.Println(stratData["border"])
	// fmt.Println(stratData["coastline"]["Hookah - Billiards"])
}
