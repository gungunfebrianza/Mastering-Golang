package main

import "fmt"

func main() {
	var maxUint32 uint32 = 4294967295
	fmt.Println(maxUint32) // Output : 4294967295

	/* 	var anotherMaxUint32 uint32 = 4294967295 + 1
	fmt.Println(anotherMaxUint32)  */
	// Output : constant 4294967296 overflows uint32
}
