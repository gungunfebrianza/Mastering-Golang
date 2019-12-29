package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2))            // Return : int
	fmt.Println(reflect.TypeOf(2.4))          // Return : float64
	fmt.Println(reflect.TypeOf("Hello Gun!")) // Return : string
	fmt.Println(reflect.TypeOf('g'))          // Return : int32
	fmt.Println(reflect.TypeOf(true))         // Return : bool
}
