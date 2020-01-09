package main

import (
	"fmt"
	"strings"
)

func main() {
	//Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	// Output : oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// Output : moo moo moo
}
