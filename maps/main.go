package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	var colorMap map[string]string

	colorMap2 := make(map[string]string)

	fmt.Println(colors)
	fmt.Println(colorMap)
	fmt.Println(colorMap2)

	colors["white"] = "#ffffff"
	delete(colors, "green")
	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
