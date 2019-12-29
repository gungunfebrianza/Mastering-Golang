package main

import (
	"fmt"
	"strings"
)

func main() {
	// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
	fmt.Println(strings.Title("gun gun febrianza")) // Output : Gun Gun Febrianza

	//ToLower returns s with all Unicode letters mapped to their lower case.
	fmt.Println(strings.ToLower("Gun Gun Febrianza")) // Output : gun gun febrianza

	//ToUpper returns s with all Unicode letters mapped to their upper case.
	fmt.Println(strings.ToUpper("gun gun febrianza")) // GUN GUN FEBRIANZA

	//Contains reports whether substr is within s.
	fmt.Println(strings.Contains("Grenade", "ena")) // Output : true
	fmt.Println(strings.Contains("Grenade", "gun")) // Output : false
	fmt.Println(strings.Contains("Gun", ""))        // Output : true
	fmt.Println(strings.Contains("", ""))           // Output : true

	//Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	fmt.Println(strings.Compare("a", "a"))         // Output : 0
	fmt.Println(strings.Compare("a", "b"))         // Output : -1
	fmt.Println(strings.Compare("b", "a"))         // Output : 1
	fmt.Println(strings.Compare("5", "5"))         // Output : 0
	fmt.Println(strings.Compare("5", "6"))         // Output : -1
	fmt.Println(strings.Compare("6", "5"))         // Output : 1
	fmt.Println(strings.Compare("makan", "makan")) // Output : 0
	fmt.Println(strings.Compare("makan", "minum")) // Output : -1
	fmt.Println(strings.Compare("minum", "makan")) // Output : 1

	//Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
	fmt.Println(strings.Count("Gun Gun Febrianza", "n")) // Output :
}
