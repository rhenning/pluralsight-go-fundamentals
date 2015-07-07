package main

import (
	"fmt"
)

func main() {
	course := "Docker Deep Dive"

	fmt.Println("course is", course)

	changeCourse(&course) // pointer to course

	fmt.Println("course  is now", course)
}

func changeCourse(course *string) string {
	*course = "First look: native docker clustering"

	fmt.Println("Trying to change course to", *course)

	return *course
}
