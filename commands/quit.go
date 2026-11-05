package commands

import (
	"os"

	"github.com/fatih/color"
)

func Quit_Command() {
	color.Red("Bye ! Bye !")
	os.Exit(0)
}
