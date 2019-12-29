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
	fmt.Println(strings.ToUpper("gun gun febrianza")) // Output : GUN GUN FEBRIANZA
}
