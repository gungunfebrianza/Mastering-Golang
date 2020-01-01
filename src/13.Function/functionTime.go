package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var sekarang time.Time = time.Now()
	var tahun int = sekarang.Year()
	fmt.Println(tahun) // Output : 2020
	broken := "gun gun"
	replacer := strings.NewReplacer("u", "n")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
