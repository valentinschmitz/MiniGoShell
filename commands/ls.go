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
		if object.IsDir() {
			color := color.New(color.FgBlue, color.Bold)
			color.Print(object.Name() + " ")
			continue
		}
		print(object.Name() + " ")
	}
	println()
	return 0
}
