package main

import "fmt"

func main() {
	var exponentPart1 float32 = 1.88e2
	fmt.Println(exponentPart1) // Output : 188

	var exponentPart2 float32 = 188e2
	fmt.Println(exponentPart2) // Output : 18800

	var exponentPart3 float32 = 188e-2
	fmt.Println(exponentPart3) // Output : 1.88
}
