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

	cars := [2]string{"Pajero","Lamborghini Veneno"}
	fmt.Println(cars)

	/* Compiler determines the length of array */
  scores := [...]int{12, 78, 50, 100, 2, 34, 56, 7, 88}
  fmt.Println(scores, len(scores))
}
