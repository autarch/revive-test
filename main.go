// main is a test package.
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	v2 := "foo"
	if _, err := fmt.Println(v2); err != nil {
		panic(err)
	}
	if _, err := fmt.Println(rand.Int()); err != nil {
		panic(err)
	}
}
