package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("ops! something went wrong")
	if err != nil {
		fmt.Print(err)
	}
}
