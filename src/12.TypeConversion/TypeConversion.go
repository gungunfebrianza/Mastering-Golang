package main

import (
	"fmt"
	"reflect"
)

func main() {
	var usia int = 65
	fmt.Println(float64(usia))                 // Output : 65
	fmt.Println(reflect.TypeOf(float64(usia))) // Output : float64
	var berat float64 = 4.8
	fmt.Println(int(berat))                 // Output : 4 (0.8 removed)
	fmt.Println(reflect.TypeOf(int(berat))) // Output : int
	fmt.Println("\r s")
}
