package main

import (
	"fmt"

	"github.com/smooth-55/go-circle-ci/something"
)

func main() {
	var operation something.Something = something.Something{
		NumberA: 10,
		NumberB: 20,
	}
	fmt.Println(operation.Add())

}
