package main

import "fmt"

var scope = "Top"

func main() {
	fmt.Println("Main start :", scope)
	// Create a shadow variable
	scope := 99
	if true {
		fmt.Println("Block start :", scope)
		funcA()
	}
	fmt.Println("Main end :", scope)
}
func funcA() {
	fmt.Println("funcA start :", scope)
}
