package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	name := os.Getenv("USER")
	// declared but unused var throws compiler error
	// course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("name is", name, "and is type", reflect.TypeOf(name))
	fmt.Println("module is", module, "and is type", reflect.TypeOf(module))
	fmt.Println("memory address of *module* variable is",
		ptr, "and value is", *ptr)
}
