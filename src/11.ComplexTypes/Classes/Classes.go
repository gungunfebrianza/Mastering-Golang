package main

import (
	"fmt"
)

type Hooman struct {
	name string
	age  int
}

func XHooman(name string, age int) *Hooman {
	return &Hooman{
		name: name,
		age:  age,
	}
}

func (f *Hooman) getName() string {
	return f.name
}

func (f *Hooman) setName(value string) {
	f.name = value
}

func main() {
	human := XHooman("Nikolaj Vestorp", 33)
	fmt.Println(human.name)
	human.setName("Gun Gun Febrianza")
	fmt.Println(human.name)
	fmt.Println(human.getName())
}
