package main

import (
	"fmt"
)

func passingByValue(count int) {
	count += 10
	fmt.Println("Passed Value : ", count)
}

func passingByPointer(count *int) {
	//Dereference the value and add 5 to it
	*count += 10
	fmt.Println("Passed Pointer : ", *count)
}

func main() {
	var count int
	passingByValue(count) // passing value
	fmt.Println("Function passingByValue Result:", count)

	fmt.Println()

	// use & to pass a pointer to the variable
	passingByPointer(&count) // passing value by pointer
	fmt.Println("Function passingByPointer Result:", count)
}
