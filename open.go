package main

import(
  "fmt"
  "os"
)

func main() {
  f, err := os.Open("file.txt")

  if err != nil {
    fmt.Println("couldn't open file")
    return
  }

  fmt.Println(f)
}
