package main

import "github.com/fatih/color"

func main() {
	// Print with default helper functions
	// Create a new color object
	color.NoColor = false
	c := color.New(color.FgCyan)
	c.Println("👌👍👋 Prints cyan text with an underline.")
}
