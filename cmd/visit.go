package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	models "github.com/omgupta1608/brino/models"
	utils "github.com/omgupta1608/brino/utils"
)

//Visit - used for visit command
func Visit(bp map[string]models.BoilerPlate, args []string) {
	ok, hasMore, flags := utils.ValidateCommand("visit", args)
	if !ok {
		fmt.Println("Invalid Command Syntax")
		return
	}
	var url string
	var err error
	if hasMore {
		url = bp[args[0]].Docs
		if flags[1] == "-d" || flags[1] == "-D" {
			switch runtime.GOOS {
			case "linux":
				err = exec.Command("xdg-open", url).Start()
			case "windows":
				err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			case "darwin":
				err = exec.Command("open", url).Start()
			default:
				err = fmt.Errorf("unsupported platform")
			}
			if err != nil {
				fmt.Printf("%s\n", err)
			}
		}
	} else {
		url = bp[args[0]].Website
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}
