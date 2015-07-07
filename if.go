package main

import (
	"fmt"
	"reflect"
)

func main() {
	//variables to store course rank
	// I think this is a bug in the course
	// these are strings not ints
	// but string comparison with <> is ok
	firstRank := "39"   // docker deep dive
	secondRank := "614" // docker clustering

	if firstRank < secondRank {
		fmt.Println("first course < second course")
	} else if firstRank > secondRank {
		fmt.Println("second course > first course")
	} else {
		fmt.Println("courses are the same",
			"or something weird is going on.")
	}

	fmt.Println("type of firstRank is",
		reflect.TypeOf(firstRank))
}
