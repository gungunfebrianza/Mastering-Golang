package main

import "fmt"

func main() {
	var maxUint8 uint8 = 255
	fmt.Println(maxUint8) // Output : 255

	/* 	var anothermaxUint8 uint8 = 255 + 1
	fmt.Println(anothermaxUint8)  */
	// Output : constant 255 overflows uint32
}
