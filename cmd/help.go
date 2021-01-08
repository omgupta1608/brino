package cmd

import (
	"fmt"
)

//Help - For the help command
func Help() {
	fmt.Println("")
	fmt.Println("====================BRINO HELP===========================")
	fmt.Println("")
	fmt.Println("following are the available commands")
	fmt.Println("")
	fmt.Println("brino make <boilerplate_name>                 used for generating a new boilerplate of the specified name")
	fmt.Println("brino explore                                 lists all the available biolerplates. Use the (-d) flag for detailed information")
	fmt.Println("brino visit <boilerplate_name>                visit the official website or documentation for the specified biolerplate (-d) to open docs")
	fmt.Println("brino help                                    for help")
}
