package main

import "fmt"

func main() {
	var unsignedInteger8 uint8 = 255
	fmt.Println(unsignedInteger8) // Output : 255

	var unsignedInteger16 uint16 = 65535
	fmt.Println(unsignedInteger16) // Ouput : 65535

	var unsignedInteger32 uint32 = 4294967295
	fmt.Println(unsignedInteger32) // Ouput : 4294967295

	var unsignedInteger64 uint64 = 18446744073709551615
	fmt.Println(unsignedInteger64) // Ouput : 18446744073709551615

	// fmt.Println(unsignedInteger16 + unsignedInteger8)
	// invalid operation: unsignedInteger16 + unsignedInteger8 (mismatched types uint16 and uint8)
}
