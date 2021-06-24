package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"

	"encoding/gob"
)

const tmpfile = ".tmp.serial"

type Animal struct {
	Name string `json:"Name"`
	Legs int    `json:"Legs"`
	Wild bool   `json:"Wild"`
}

func plan(fp string) error {
	// test data
	x := &Animal{
		Name: "dog",
		Legs: 4,
		Wild: false,
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(x); err != nil {
		fmt.Print(err)
	}

	err := os.WriteFile(fp, buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func apply(fp string) error {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		return fmt.Errorf("no changes to apply. run 'plan' first")
	}
	// do stuff
	b, err := os.ReadFile(fp)
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile)

	var data Animal

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&data); err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func main() {
	// plan subcmd
	planCmd := flag.NewFlagSet("plan", flag.ExitOnError)
	outFile := planCmd.String("filepath", tmpfile, "file to apply")

	// apply subcmd
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	inFile := applyCmd.String("filepath", tmpfile, "file to apply")

	if len(os.Args) < 2 {
		fmt.Println("usage: a3 <command> [<args>]")
		fmt.Println("  plan\tprint changes to output without apply the changes.")
		fmt.Println("  apply\tapply the changes (as per the plan)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "plan":
		fmt.Println("** Plan **")
		err := plan(*outFile)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
	case "apply":
		fmt.Println("** Apply **")
		applyCmd.Parse(os.Args[2:])
		err := apply(*inFile)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Print with default helper functions
	// Create a new color object
	color.NoColor = false
	c := color.New(color.FgHiGreen).Add(color.Underline)
	fmt.Printf("👌👍👋 %s", c.Sprintf("Prints cyan text with an underline."))
}
