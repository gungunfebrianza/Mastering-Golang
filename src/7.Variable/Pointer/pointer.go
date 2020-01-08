package main

import (
	"fmt"
	"time"
)

func main() {
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

	fmt.Println()

	var pointerInteger *int //Output : nil
	fmt.Printf("pointerInteger: %#v\n", pointerInteger)
	// Output : pointerInteger: (*int)(nil)

	count2 := new(int)
	fmt.Printf("count2: %#v\n", count2)
	// Output : count2: (*int)(0xc0000140f0)

	countTemp := 5
	count3 := &countTemp
	fmt.Printf("count3: %#v\n", count3)
	// Output : count3: (*int)(0xc0000140f8)
	fmt.Println(count3)
	// Output :  0xc0000140f8

	t := &time.Time{}
	fmt.Printf("time : %#v\n", t)
	// Output :  &time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}
}
