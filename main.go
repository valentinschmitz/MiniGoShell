package main

import (
	"MiniGoShell/commands"
	"fmt"

	"github.com/fatih/color"
)

func check_command(input *string, last_exit_code *int) int {
	if *input == "quit" {
		commands.Quit_Command()
	}
	if *input == "$?" {
		commands.Get_Exit_Code(last_exit_code)
	}
	if *input == "ls" {
		return commands.Ls_Command()
	}
	return 84
}

func print_prefix() {
	color := color.New(color.FgCyan, color.Bold)
	color.Print("MiniGoShell> ")
}

func main() {
	var input string = ""
	var last_exit_code int = 0

	whileLoop := true
	for whileLoop {
		print_prefix()
		fmt.Scan(&input)
		last_exit_code = check_command(&input, &last_exit_code)
	}
}
