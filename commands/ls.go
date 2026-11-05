package commands

import (
	"os"

	"github.com/fatih/color"
)

func Ls_Command() int {
	list, err := os.ReadDir("./")

	if err != nil {
		color.Red("Error: %v", err)
		return 84
	}

	for _, object := range list {
		print(object.Name() + " ")
	}
	println()
	return 0
}
