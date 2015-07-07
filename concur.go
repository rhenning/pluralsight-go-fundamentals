package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // max # of parallel threads

	var waitGrp sync.WaitGroup
	waitGrp.Add(2) // size of wait group

	go func() { // anonymous function
		defer waitGrp.Done()

		time.Sleep(5 * time.Second) // 5 seconds
		fmt.Println("Hello")
	}() // self-execute immediately

	go func() { // this is non-blocking
		defer waitGrp.Done()

		fmt.Println("Pluralsight")
	}()

	// WARN: if main() exits all goroutines immediately quit!

	waitGrp.Wait()
}
