package main

import "github.com/fatih/color"

func main() {
	// Print with default helper functions
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("👌👍👋 Prints cyan text with an underline.")
}
