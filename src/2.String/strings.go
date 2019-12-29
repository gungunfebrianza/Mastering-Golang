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
	fmt.Println(strings.Count("Gun Gun Febrianza", "n"))  // Output : 3
	fmt.Println(strings.Count("Gun Gun Febrianza", "un")) // Output : 2

	//EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding.
	fmt.Println(strings.EqualFold("Gun", "gun")) // Output : true
	fmt.Println(strings.EqualFold("Gun", "Gan")) // Output : false

	//HasPrefix tests whether the string s begins with prefix.
	fmt.Println(strings.HasPrefix("Marketkoin", "mar")) // Output : false
	fmt.Println(strings.HasPrefix("Marketkoin", "Mar")) // Output : true
	fmt.Println(strings.HasPrefix("Marketkoin", "Mur")) // Output : false
	fmt.Println(strings.HasPrefix("Marketkoin", ""))    // Output : true

	//HasSuffix tests whether the string s ends with suffix.
	fmt.Println(strings.HasSuffix("Marketkoin", "koin")) // Output : true
	fmt.Println(strings.HasSuffix("Marketkoin", "Koin")) // Output : false
	fmt.Println(strings.HasSuffix("Marketkoin", "Mar"))  // Output : false
	fmt.Println(strings.HasSuffix("Marketkoin", ""))     // Output : true

	//Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	fmt.Println(strings.Index("Febrianza", "nza")) // Output : 6
	fmt.Println(strings.Index("Febrianza", "nzo")) // Output : -1
	fmt.Println(strings.Index("Febrianza", ""))    // Output : 0

	//LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
	fmt.Println(strings.Index("Gun Gun Febrianza", "Gun"))     // Output : 0
	fmt.Println(strings.LastIndex("Gun Gun Febrianza", "Gun")) // Output : 4
	fmt.Println(strings.LastIndex("Gun Gun Febrianza", "Gon")) // Output : -1
	fmt.Println(strings.LastIndex("Gun Gun Febrianza", ""))    // Output : 17

	//Repeat returns a new string consisting of count copies of the string s.
	fmt.Println(strings.Repeat("Gun", 2)) // Output : GunGun

	//Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

}
