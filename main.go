package main

import (
	"fmt"
)

func main() {
	var input string = ""

	whileLoop := true
	for whileLoop {
		fmt.Scan(&input)
		fmt.Println("User said : " + input)
	}
}
