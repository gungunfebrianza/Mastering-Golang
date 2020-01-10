package main

import "fmt"

func main() {
	var SignedInteger8 int8 = 127
	fmt.Println(SignedInteger8) // Output : 127

	var SignedInteger16 int16 = 32767
	fmt.Println(SignedInteger16) // Ouput : 32767

	var SignedInteger32 int32 = 2147483647
	fmt.Println(SignedInteger32) // Ouput : 2147483647

	var SignedInteger64 int64 = 9223372036854775807
	fmt.Println(SignedInteger64) // Ouput : 9223372036854775807

	// fmt.Println(SignedInteger32 + SignedInteger64)
	// invalid operation: SignedInteger32 + SignedInteger64 (mismatched types int32 and int64)
}
