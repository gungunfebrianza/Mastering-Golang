package main

import (
	"fmt"
	"reflect"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(2, 3)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result)
}
