package main

import "fmt"

func greet(name *string) string {
	n := "stranger"
	if name != nil {
		n = *name
	}

	return fmt.Sprintf("hello %s", n)
}

func main() {
	message := greet(nil)
	fmt.Println(message)

	name := "bob"
	message = greet(&name)
	fmt.Println(message)
}
