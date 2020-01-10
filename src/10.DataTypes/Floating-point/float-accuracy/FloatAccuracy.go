package main

import "fmt"

func main() {
	var intVariable int = 100
	var float32Variable float32 = 100
	var float64Variable float64 = 100
	fmt.Println(intVariable / 3)     // Output : 33
	fmt.Println(float32Variable / 3) // Output : 33.333332
	fmt.Println(float64Variable / 3) // Output : 33.333333333333336
}
