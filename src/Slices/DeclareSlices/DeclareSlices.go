package main

import "fmt"

func main() {
	slices := []string{"Gun", "Gan", "Gin", "Gon", "Gen"}
	fmt.Println(slices)

	array := [7]int{22, 34, 56, 78, 88, 44, 33}
	var slice []int = array[1:4]
	fmt.Println(slice)
}
