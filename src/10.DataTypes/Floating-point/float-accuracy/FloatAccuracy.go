	package main

	import (
		"fmt"
		"math"
	)

	func main() {
		var intVariable int = 100
		var float32Variable float32 = 100
		var float64Variable float64 = 100
		fmt.Println(intVariable / 3)     // Output : 33
		fmt.Println(float32Variable / 3) // Output : 33.333332
		fmt.Println(float64Variable / 3) // Output : 33.333333333333336

		var pi64 = math.Pi
		var pi32 float32 = math.Pi
		fmt.Println(pi64)
		fmt.Println(pi32)
		/*
			3.141592653589793
			3.1415927
		*/
	}
