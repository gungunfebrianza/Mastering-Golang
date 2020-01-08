package main

import (
	"fmt"
)

func main() {
	variableInteger := 5
	// Declare a pointer using existing variable
	var dataPointer = &variableInteger
	fmt.Printf("variableInteger: %#v\n", variableInteger)
	// Output : variableInteger: 5
	fmt.Printf("dataPointer: %#v\n", dataPointer)
	// Output : dataPointer: (*int)(0xc0000140b0)

	var pointerValue = dataPointer
	fmt.Println(pointerValue)
	// Output : 0xc0000140b0

	if pointerValue != nil {
		fmt.Printf("pointerValue: %#v\n", *pointerValue)
		// Output : pointerValue: 5
	}
}
