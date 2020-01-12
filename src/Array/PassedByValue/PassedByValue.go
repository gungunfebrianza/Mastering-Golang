package main

import (
	"fmt"
)

func changeArray(numbers [5]int) {
	numbers[0] = 5
	fmt.Println("Change Number : ", numbers)
}

func main() {
	countries := [...]string{"USA", "UK", "AU"}
	mycountries := countries
	fmt.Println(mycountries)

	numbers := [...]int{1, 2, 3, 4, 5}
	changeArray(numbers)
	fmt.Println("Original Numbers : ", numbers)
}
