package commands

import "github.com/fatih/color"

func Get_Exit_Code(last_exit_code *int) {
	if *last_exit_code == 0 {
		color.Green("0")
	} else {
		color.Red("%d", *last_exit_code)
	}
}
