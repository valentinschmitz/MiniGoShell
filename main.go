package main

import (
	"fmt"
	"os"
)

func check_command(input *string) {
	if *input == "quit" {
		fmt.Println("Bye ! Bye !")
		os.Exit(0)
	}
}

func main() {
	var input string = ""

	whileLoop := true
	for whileLoop {
		fmt.Scan(&input)
		check_command(&input)
	}
}
