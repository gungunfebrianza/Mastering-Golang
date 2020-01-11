package main

import "fmt"

func main() {
	var decimalPoint1 float32 = 1.23
	fmt.Println(decimalPoint1) // Ouput : 1.23

	var decimalPoint2 float32 = 01.23
	fmt.Println(decimalPoint2) // Ouput : 1.23

	var decimalPoint3 float32 = .23
	fmt.Println(decimalPoint3) // Ouput : 0.23

	var decimalPoint4 float32 = 2.
	fmt.Println(decimalPoint4) // Ouput : 2

}
