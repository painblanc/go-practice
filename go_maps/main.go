package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}

	var colors map[string]string

	colours["white"] = "#FFF"
	fmt.Println(colors)
	fmt.Println(colours)

	printMap(colours)
	delete(colours, "white")
	fmt.Println(colours)

}

func printMap(m map[string]string) {
	for colour, hex := range m {
		fmt.Println("The colour", colour, "has a hex value of", hex)
	}
}
