package main

import (
	"fmt"
)

func main() {
	myCourses := make([]string, 5, 10)

	fmt.Printf("Length is %d Capacity is %d\n",
		len(myCourses), cap(myCourses))
}
