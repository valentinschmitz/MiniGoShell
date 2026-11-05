package commands

import (
	"github.com/fatih/color"
)

func Help_Command() int {
	color.Yellow("Available commands:")
	color.Yellow("  quit - Exit the shell")
	color.Yellow("  $? - Display the last exit code")
	color.Yellow("  ls - List directory contents")
	color.Yellow("  help - Show this help message")
	return 0
}
