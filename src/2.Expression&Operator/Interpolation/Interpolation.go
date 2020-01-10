package main

import "fmt"

func main() {
	name := "Gun"
	age := 28
	message := fmt.Sprintf("%s is %d years old", name, age)

	fmt.Println(message)
}
