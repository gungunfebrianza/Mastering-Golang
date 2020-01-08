package main

import (
	"fmt"
)

func main() {
	var pointerInteger *int //Output : nil
	fmt.Printf("pointerInteger: %#v\n", pointerInteger)
	// Output : pointerInteger: (*int)(nil)

	if pointerInteger != nil {
		fmt.Printf("pointerInteger: %#v\n", *pointerInteger)
		// Output : no output
	}

}
