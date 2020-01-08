package main

import (
	"fmt"
	"time"
)

func main() {
	var pointerInteger *int //Output : nil
	fmt.Printf("pointerInteger: %#v\n", pointerInteger)
	// Output : pointerInteger: (*int)(nil)

	if pointerInteger != nil {
		fmt.Printf("pointerInteger: %#v\n", *pointerInteger)
	}

	count2 := new(int)
	fmt.Printf("count2: %#v\n", count2)
	// Output : count2: (*int)(0xc0000140f0)
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}

	pointerVault := 5
	count3 := &pointerVault
	fmt.Printf("count3: %#v\n", count3)
	// Output : count3: (*int)(0xc0000140f8)
	fmt.Println(count3)
	// Output :  0xc0000140f8
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}

	t := &time.Time{}
	fmt.Printf("time : %#v\n", t)
	// Output :  &time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}
}
