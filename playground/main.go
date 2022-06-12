package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	test := errors.New("test")

	test2 := errors.Wrap(test, "kocak lu")

	test3 := errors.Wrap(test, test2.Error())

	if errors.Cause(test2) == test {
		fmt.Print("masok")
	} else {
		fmt.Println("ga masok")
	}

	fmt.Print(test3)
}
