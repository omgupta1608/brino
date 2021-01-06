package cmd

import (
	"fmt"

	"github.com/omgupta1608/brino/utils"
)

//Explore - used for explore command
func Explore(obj map[string]interface{}) {
	data, ok := utils.FindOne(obj, "boilerplates")
	if !ok {
		fmt.Println("Something wrong")
		return
	}
	fmt.Println(data)
}
