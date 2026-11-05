package main

import (
	quit "MiniGoShell/commands"
	"fmt"

	"github.com/fatih/color"
)

func check_command(input *string) {
	if *input == "quit" {
		quit.Quit_Command()
	}
}

func print_prefix() {
	color := color.New(color.FgCyan, color.Bold)
	color.Print("MiniGoShell> ")
}

func main() {
	var input string = ""

	whileLoop := true
	for whileLoop {
		print_prefix()
		fmt.Scan(&input)
		check_command(&input)
	}
}
