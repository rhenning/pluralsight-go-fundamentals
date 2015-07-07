package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.000
	b := 3

	fmt.Println("A is type", reflect.TypeOf(a))
	fmt.Println("B is type", reflect.TypeOf(b))

	c := a + float64(b)

	fmt.Println("C has value", c, "and is type", reflect.TypeOf(c))
}
