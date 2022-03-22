package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	/*fmt.Println("before insert",colors)
	colors["yellow"] = "laksjdf"
	//delete(colors, "yellow")
	fmt.Println("after insert",colors)
	delete(colors, "yellow")
	fmt.Println("after deletion", colors)
	*/
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
