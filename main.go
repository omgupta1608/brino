package main

import (
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
			cmd.Explore(bp)
		}
	}
}
