package cmd

import (
	"fmt"

	models "github.com/omgupta1608/brino/models"
	utils "github.com/omgupta1608/brino/utils"
)

//Explore - used for explore command
func Explore(bp []models.BoilerPlate, args []string) {

	ok, hasMore, flags := utils.ValidateCommand("explore", args)
	if !ok {
		fmt.Println("Invalid Command")
		return
	}

	if hasMore {
		if flags[0] == "-d" || flags[0] == "-D" {
			for i := 0; i < len(bp); i++ {
				fmt.Printf("%s - %s\n", bp[i].ID, bp[i].Details)
			}
		}
	} else {
		for i := 0; i < len(bp); i++ {
			fmt.Println(bp[i].ID)
		}
	}
}
