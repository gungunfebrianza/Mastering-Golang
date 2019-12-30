package main

import (
	"fmt"
	"reflect"
)

func main() {
	nama := "Gun Gun Febrianza"
	fmt.Println(reflect.TypeOf(nama)) // Output : string
	umur := 27
	fmt.Println(reflect.TypeOf(umur)) // Output : int
	tagihan := 11.33
	fmt.Println(reflect.TypeOf(tagihan)) // Output : float64
	nganggur := true
	fmt.Println(reflect.TypeOf(nganggur)) // Output : bool
}
