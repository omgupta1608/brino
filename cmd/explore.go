package cmd

import (
	"fmt"

	models "github.com/omgupta1608/brino/models"
	utils "github.com/omgupta1608/brino/utils"
)

//Explore - used for explore command
func Explore(bp map[string]models.BoilerPlate, args []string) {

	ok, hasMore, flags := utils.ValidateCommand("explore", args)
	if !ok {
		fmt.Println("Invalid Command Syntax")
		return
	}

	if hasMore {
		if flags[0] == "-d" || flags[0] == "-D" {
			for _, b := range bp {
				fmt.Printf("%s - %s\n", b.ID, b.Details)
			}
		}
	} else {
		for _, b := range bp {
			fmt.Printf("%s\n", b.ID)
		}
	}
}
