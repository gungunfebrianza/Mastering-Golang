package main

import (
	"fmt"
)

func main() {

	input := 7
	if input < 0 {
		fmt.Println("Input must be positive integer")
	} else if input%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
