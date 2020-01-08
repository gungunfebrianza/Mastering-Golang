package main

import (
	"fmt"
)

func main() {
	variableInteger := 5
	// Declare a pointer using existing variable
	var dataPointer = &variableInteger
	fmt.Printf("variableInteger: %#v\n", variableInteger)
	// Output : pointerInteger: 5
	fmt.Printf("dataPointer: %#v\n", dataPointer)
	// Output : dataPointer: (*int)(0xc0000140b0)
	var pointerValue = dataPointer
	if pointerValue != nil {
		fmt.Printf("pointerValue: %#v\n", *pointerValue)
		// Output : pointerValue: 5
	}
}
