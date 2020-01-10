package main

import (
	"fmt"
)

type Foo struct {
	Item string
}

func XFoo(value string) *Foo {
	return &Foo{
		Item: value,
	}
}

func (f *Foo) getItem() string {
	return f.Item
}

func (f *Foo) setItem(value string) {
	f.Item = value
}

func main() {
	foox := XFoo("test")
	fmt.Println(foox.Item)
	foox.setItem("Skill!")
	fmt.Println(foox.Item)
	item := foox.getItem()
	fmt.Println(item)
}
