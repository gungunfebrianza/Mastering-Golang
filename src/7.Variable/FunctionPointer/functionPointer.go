package main

import (
	"fmt"
)

func add5Value(count int) {
	count += 5
	fmt.Println("Add 5 Value : ", count)
}

func add5Point(count *int) {
	//Dereference the value and add 5 to it
	*count += 5
	fmt.Println("Add 5 Point : ", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)

	// use & to pass a pointer to the variable
	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
