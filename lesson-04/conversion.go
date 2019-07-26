package main

import (
	"fmt"
)

type A struct {
	C string
}

type B struct {
	C string
}

func main() {
	a := B{"idris"}
	b := B{"idris"}

	fmt.Println(a == b)
}
