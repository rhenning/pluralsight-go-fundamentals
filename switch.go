package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "linux":
		fmt.Println("Linux")
	case "docker":
		fmt.Println("Docker")
	default:
		fmt.Println("Sorry, no match!")
	}
}
