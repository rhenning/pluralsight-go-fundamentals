package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1, 4)

	fmt.Printf("Length is %d Capacity is %d\n",
		len(mySlice), cap(mySlice))

	for i := 1; i < 17; i++ {
		mySlice = append(mySlice, i)

		fmt.Printf("Capacity is %d\n", cap(mySlice))
	}
}
