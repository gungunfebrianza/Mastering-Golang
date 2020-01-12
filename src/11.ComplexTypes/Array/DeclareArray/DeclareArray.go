package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names [3]string

	names[0] = "Gun Gun Febrianza"
	names[1] = "Nikolaj Vestorp"
	names[2] = "Racheysa Hazna"

	fmt.Println(names[0])
	fmt.Println(names)
	fmt.Println(reflect.TypeOf(names))
}
