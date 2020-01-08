package main

import "fmt"

func main() {
	fmt.Println(4 > 1)  // Output : true
	fmt.Println(4 == 2) // Output : false
	fmt.Println(4 != 2) // Output : true

	//Lexicography
	fmt.Println('Z' > 'C')       // Output : true
	fmt.Println("Glow" > "Glee") // Output : true
}
