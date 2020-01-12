package main

import "fmt"

func main() {
	mynames := [2]string{"Gun Gun Febrianza", "Nikolaj Vestorp"}
	for i := 0; i < len(mynames); i++ {
		fmt.Println(mynames[i])
	}

	for i, v := range mynames {
		fmt.Println(i,v)
	}

	for _, v := range mynames {
		fmt.Println(v)
	}
}
