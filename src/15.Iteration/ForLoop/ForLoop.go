package main

import (
	"fmt"
)

func main() {
	for index := 0; index < 5; index++ {
		fmt.Println(index)
	}
	for index := 0; index <= 5; index++ {
		fmt.Println(index)
	}
}
