package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bfks",
		"white": "#fffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// // add values
	// colors["white"] = "#ffff"

	// // remove values
	// delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
