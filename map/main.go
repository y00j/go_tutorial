package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "this is red",
		"white": "this is white",
		"blue":  "blue dabadeedaba",
	}

	printMap(colors)
	// fmt.Println(colors)
}

func printMap(m map[string]string) {
	for color,desc:=range m {
		fmt.Println("color is:", color, ", description is:", desc)
	}
} 
