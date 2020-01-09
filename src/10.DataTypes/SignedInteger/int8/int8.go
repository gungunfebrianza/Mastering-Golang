package main

import "fmt"

func main() {
	var maxInt8 int8 = 127
	fmt.Println(maxInt8) // Output : 127

	/* 	var anothermaxInt8 int8 = 127 + 1
	fmt.Println(anothermaxInt8)  */
	// Output : constant 128 overflows int8
}
