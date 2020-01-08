package main

import "fmt"

func main() {
	names := []string{"Gun", "Gon", "Gin"}
	//The built-in function, len, is used to get the length of any collection.
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
