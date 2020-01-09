package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	var y string
	var z bool
	var pointer *int
	fmt.Println(x)       // Output : 0
	fmt.Println(y)       // Output : ""
	fmt.Println(z)       // Output : false
	fmt.Println(pointer) // Output : <nil>

	var thisInteger int
	fmt.Printf("thisInteger : %#v \n", thisInteger)
	// Output : thisInteger : 0

	var thisFloat64 float64
	fmt.Printf("thisFloat64 : %#v \n", thisFloat64)
	// Output : thisFloat64 : 0

	var thisBool bool
	fmt.Printf("thisBool : %#v \n", thisBool)
	// Output : thisBool : false

	var thisString string
	fmt.Printf("thisString : %#v \n", thisString)
	// Output : thisString : ""

	var thisArrayString []string
	fmt.Printf("thisArrayString : %#v \n", thisArrayString)
	// Output : thisArrayString : []string(nil)

	var thisTime time.Time
	fmt.Printf("thisTime : %#v \n", thisTime)
	// Output : thisTime : time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}
}
