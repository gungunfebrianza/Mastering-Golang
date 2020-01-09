package main

import "fmt"

func main() {
	var maxInt32 int32 = 2147483647
	fmt.Println(maxInt32) // Output : 2147483647

	/* 	var anotherMaxInt32 int32 = 2147483647 + 1
	fmt.Println(anotherMaxInt32)  */
	// Output : constant 2147483648 overflows int32
}
