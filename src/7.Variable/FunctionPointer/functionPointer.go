package main

import (
	"fmt"
)

func passingValue(count int) {
	count += 10
	fmt.Println("Passing Value Result : ", count)
}

func passingPointer(count *int) {
	//Dereference the value and add 5 to it
	*count += 10
	fmt.Println("Passing Pointer Result : ", *count)
}

func main() {
	var count int
	passingValue(count)
	fmt.Println("Function passingValue:", count)
	fmt.Println()
	// use & to pass a pointer to the variable
	passingPointer(&count)
	fmt.Println("Function passingPointer post:", count)
}
