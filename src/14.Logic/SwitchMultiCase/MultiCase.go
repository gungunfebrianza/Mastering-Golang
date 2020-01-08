package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Tuesday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday:
		fmt.Println("This is Saturday")
	default:
		fmt.Println("Day, is not valid! Error!")
	}
}
