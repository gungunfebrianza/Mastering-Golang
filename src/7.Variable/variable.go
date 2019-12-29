package main

import "fmt"

func main() {
	var name string            // Declare Variable
	var height, weight float64 // Declare Variable
	var age int                // Declare Variable

	name = "Gun Gun Febrianza"
	height, weight = 168, 55
	age = 27

	fmt.Println(name)
	fmt.Println("Nama : ", name, "usia", age)
	fmt.Println("Tinggi Badan :", height)
	fmt.Println("Berat Badan :", weight)
	/* Output :
	Gun Gun Febrianza
	Nama :  Gun Gun Febrianza  usia 27
	Tinggi Badan :  168
	Berat Badan :  55
	*/
}
