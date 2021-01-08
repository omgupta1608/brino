package main

import (
	"fmt"
	"os"

	cmd "github.com/omgupta1608/brino/cmd"
	"github.com/omgupta1608/brino/data"
)

func main() {
	bp := data.LoadData()
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			cmd.Help()
		case "explore":
			cmd.Explore(bp, os.Args)
		case "visit":
			cmd.Visit(bp, os.Args)
		case "make":
			cmd.Make()
		default:
			fmt.Println("That is not a valid brino command")
		}
	}
}
