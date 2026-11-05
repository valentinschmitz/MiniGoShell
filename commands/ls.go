package commands

import (
	"os"
)

func Ls_Command() {
	list, err := os.ReadDir("./")

	if err != nil {
		panic(err)
	}

	for _, object := range list {
		print(object.Name() + " ")
	}
	println()
}
