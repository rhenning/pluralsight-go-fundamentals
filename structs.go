package main

import "fmt"

func main() {
  type courseMeta struct {
    Author string
    Level string
    Rating float64
  }

  var DockerDeepDive courseMeta
  
  // or for a pointer:
  DockerDeepDive2 := new(courseMeta)
  
  // composite literal:
  DockerDeepDive3 := courseMeta{
    Author: "Nigel Poulton",
    Level: "Intermediate",
    Rating: 5,
  }

  // and without field names:
  DockerDeepDive4 := courseMeta{
    "Nigel Poulton",
    "Intermediate",
    5,
  }

  fmt.Println(DockerDeepDive)
  fmt.Println(DockerDeepDive2)
  fmt.Println(DockerDeepDive3)
  fmt.Println(DockerDeepDive4)
}
