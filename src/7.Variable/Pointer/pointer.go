package main

import (
	"fmt"
	"time"
)

func main() {
	var thisInteger int
	fmt.Printf("thisInteger : %#v \n", thisInteger)

	var thisFloat64 float64
	fmt.Printf("thisFloat64 : %#v \n", thisFloat64)

	var thisBool bool
	fmt.Printf("thisBool : %#v \n", thisBool)

	var thisString string
	fmt.Printf("thisString : %#v \n", thisString)

	var thisArrayString []string
	fmt.Printf("Emails : %#v \n", thisArrayString)

	var thisTime time.Time
	fmt.Printf("thisTime : %#v \n", thisTime)

	fmt.Println()

	var pointerInteger *int //Output : nil
	fmt.Printf("pointerInteger: %#v\n", pointerInteger)

	count2 := new(int)
	fmt.Printf("count2: %#v\n", count2)

	countTemp := 5
	count3 := &countTemp
	fmt.Printf("count3: %#v\n", count3)
	fmt.Println(count3)

	t := &time.Time{}
	fmt.Printf("time : %#v\n", t)
}
