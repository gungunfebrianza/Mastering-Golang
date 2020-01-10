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

	var exponentPart1 float32 = 1.88e2
	fmt.Println(exponentPart1) // Output : 188

	var exponentPart2 float32 = 188e2
	fmt.Println(exponentPart2) // Output : 18800

	var exponentPart3 float32 = 188e-2
	fmt.Println(exponentPart3) // Output : 1.88
}
