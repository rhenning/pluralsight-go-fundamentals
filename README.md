# Overview

* Simple (only 25 keywords)
* Compiled
* Statically linked
* Statically typed
* Supports type inference

## Reference

* This is excellent: https://github.com/a8m/go-lang-cheat-sheet
* Idiomatic Go: https://golang.org/doc/effective_go.html

# Some conventions

* packages should be single-word lowercase names
* functions and vars should be mixedCase or MixedCase

# Install it

# Set up workspace

* Create `~/gocode/{src,pkg,bin}`
* Set `$GOPATH` env var to workspace dir

# Hello World

```
package main  // means this is an executable not library

import "fmt"

// entry point of executable (as opposed to shared lib)
// takes no arguments has no return values
func main() {
  fmt.Println("hi!")
}
```

# Compiling and Running

`go run file.go` compiles and immediately runs.

# Comments

C-style line and block comments.  Line comments are idiomatic, even when spanning multiple lines.

```
// line
// comments

/*
 * block
 * comments
 */
 ```

# Function visibility

To export a function from its package declare it with a capital letter

```
func Something() {
  // public
}

func other() {
  // private
}
```

# Variables and Constants

## Declaring and initializing variables

* Globals must start with `var` keyword
* Uninitialized variable is "zero" or empty string

```
package main

import(
  "fmt"
)

var(
  name, course  string
  module        float64
)

func main() {
  fmt.Println("name is set to", name)
  fmt.Println("module is set to", module)
}
```

* Declaring vars inside functions can use type inference

```
function this() {
  name := "Bill"
  age := 26
}
```

## Checking types

```
import "reflect"
...
reflect.TypeOf(thing)
```

## Typecasting

```
c := float64(b) + a  // c initialized as float64
```

## Pointers

Go passes arguments by value not reference

* Variables passed in function arguments are copied onto the stack
* `&somevar` means "memory address of somevar" or "pointer to somevar"
* `*somevar` means dereference a pointer to somevar

```
func main() {
  module := 3.2
  modulePtr := &module

  fmt.Println("Memory address of module is ",
      modulePtr, "and value is", *modulePtr)
}
```

* Pointers must be used to pass arguments by reference

```
func main() {
  this := "this"
  fmt.Println("this is", this)
  changeValue(&this)
  fmt.Println("this is now", this)
}

func changeValue(thingy *string) string {
  *thingy = "that"

  return *thingy
}
```

## Constants

* Constants are declared with the `const` keyword

```
const pi = 3.1415927
```

## Environment Variables

```
package main

import(
  "fmt"
  "os"
)

func main() {
  for _, env := range os.Environ() {
    fmt.Println(env)
  }
}
```

# Functions

* Multiple return values
* Return values can be named in which case vars are created

```
//declaration syntax
func funcName(paramName paramType) returnType {
  return somethingOfReturnType
}

// example
func myFunc(s string) string {
  return strings.ToUpper(s)
}

// with named multiple params
func f(s1, s2 string) (this, that string) {
  // do stuff
}

// named variables of different types
func f2(s1 string, i int) (this string, that int) {
  // do things
}
```

## Variadic Functions
A *variadic function* is a function with a variable 
number of parameters. To declare one use an ellipsis 
before the type.

```
func f(s1 ...int) int {
  // stuff here
}
```

## Anonymous self-executing functions

```
func() {
  // do something
}()
```

# Conditionals and Branching

* if
* switch/case

## if

* No parens necessary
* *Must* evaluate to a boolean
  * Cannot use ints/strings in expressions

```
if <expression> {
  doSomething()
} else if <expression> {
  doSomethingElse()
} else {
  doTheDefaultThing()
}
```

Can include simple initialization statements
for scope of the if block

```
if firstRank, secondRank := 39, 614 ; firstRank < secondRank {
  // do something if firstRank < secondRank
  // firstRank, secondRank are GC'd upon block exit
}
```

## switch

* Can also take an optional simple initialization ("simple statement" in go-speak)
* No fall-through - first match is run then block exited

```
switch <simpleStatement> ; <expression> {
  case <expression>: <code>
  case <expression>: <code>
  default: <code>
}
```

```
switch "Docker Deep Dive" {
  case "Docker Deep Dive": <code>
  case "Go Fundamentals": <code>
  default: <code>
}
```

Handling fall-through:

```
switch "Docker Deep Dive" {
  case "Docker Deep Dive":
    <code>
    fallthrough
  case "Go Fundamentals": <code>
  default: <code>
}
```

Idiom is to match multiple cases:

```
func main() {
  switch tmpNum := myGimmeRandom(); tmpNum {
  case 0, 2, 4, 6, 8:
    fmt.Println("even", tmpNum)
  case 1, 3, 5, 7, 9:
    fmt.Println("odd", tmpNum)
  }
}
```

## The role of *if* in error handling

It is idiomatic to return `error` as the last
return from functions and methods. `nil` is
returned for success.

```
func testConn(target string) (rspTime float64, err error) {
  // do stuff
  return respTime, err
}
```

Common error handling:

```
f, err = os.Open(filename)

if err != nil {
  //handle errors
  fmt.Println("error opening file", filename)
  return
}

// normal execution resumes
// skip if/else for error handling

```

# Loops

* Only one `for` loop, overloaded with many incantations

```
for <expression>
```

Expression can be
* Blank for infinite loop
* Boolean expression
* A range

Examples:

```
for {
  // infinite. `break` to exit
}

for i < 10 {
  // loop while i < 10 
}

for i := range courseList {
  //assign current val of courseList to i in turn
}

for i := 0; i < 10; i++ {
  // standard C-style iterative loop
}
```

## for ... range

```
func main() {
  coursesInProg := []string{  // a "slice" of strings
    "Docker Deep Dive",
    "Docker Clustering",
    "Docker and Kubernetes",
  }

  for _, i := range coursesInProg {
    fmt.Println(i)
  }
}
```

## break and continue

* Use `break` to break out of the innermost loop
* Use `continue` to immediately jump to the next iteration in a loop
* To break out of an outer loop use a label:


```
for <expr> {
breakPoint // arbitrary label
  for <expr> {
    break breakPoint
  }
}
```

# Arrays and slices

* All items must be of same type
* Arrays are fixed length
* Slices are variable length
  * Internally they are built on top of arrays
  * Internally a list of references to an array element
  * Because they are references they are *passed by reference by default*
  * Variable up to size of backing arary
* Slices are the preferred Go idiom

* Use `make` to create a slice:

```
myCourses := make(
  []string, // type
  5,        // length now
  10        // capacity (max len, sz of backing array)
)
```

* `len(mySlice)` to get length
* `cap(mySlice)` to get capacity
* Initialized with zero values by default
* Omitting `cap` from `make` creates slice with cap same as length

* Or declare a literal with same len/cap:

```
myCourses := []string{"this", "that", "other"}
```

* When getting ranges of slices the left hand side is inclusive but the right hand side is exclusive.

```
sliceOfSlice := mySlice[2:5] // 2-4!! not 2-5
```

* Use `append` to add a value to a slice
  * Slice is grown automatically when running over capacity


* Iterate a slice with a ```for ... range``` loop
  * Index and value are returned

```
for index, data := range mySlice {
  // do something
}
```

* Concatenate two slices with ellipses like so:

```
mySlice := []int{1, 2, 3, 4, 5}
newSlice := []int{10, 20, 30}

mySlice = append(mySlice, newSlice...)
```

* To create a multidimensional slice just use slices as slice elements as expected

# Maps

* aka "hash table", "dictionary", "associative array"
* **They are unordered**
  * Seriously. The starting offset is randomized.
* Dynamically resizable
* Also passed by reference like slices
* Declare like `make(map[keyType]valueType)`
* Or as a literal

```
for k, v := range map[string]string {
  "foo": "bar",
  "baz": "biddy",
} {
  // do something with k and v
}
```

## Iteration and ordering

* Iterate with a `for ... range` loop
  * Returns key, value instead of index, value
  * **Start offset is totally random**

## Access, insert, update, delete

* As expected:

```
myMap = make(map[string]int)
myMap["X"] = 100
delete(myMap, "Y")  // Y is key
```

## References and performance

* Remember that maps and slices are *not* copied when passed to a function
* Unsafe for concurrency
* Specify size for large maps if known to avoid realloc
  * make(map[string]string, 1000)

# Structs

* Custom data types
* Basically a collection of named fields

```
type Circle struct {
  r float64 // radius
  d float64 // diameter
  c float64 // circumference
  // yeah we wouldn't need all these in reality
}
```

* Structs are how Go "does OOP"
  * A bunch of named fields/functions
    * Fields can be of different types
  * Not really OO
    * No objects
    * No classes
    * No inheritance

* When initializing a struct you can leave out
field names if you remember the order

```
DockerDeepDive4 := courseMeta{
  "Nigel Poulton",
  "Intermediate",
  5,
}
```

* Use periods to access fields for read or write

```
DockerDeepDive.Rating = 0
fmt.Println(DockerDeepDive.Rating)
```

* Methods on structs: https://gobyexample.com/methods

# Concurrency

Go's implementation is based on the Actor model, aka "concurrent sequential processes".

* Concurrency is about creating multiple processes that execute *independently*
  * This does *not* mean in parallel or simultaneously

Rob Pike says "Concurrency is about *dealing* with lots of
things (events). Parallelism is about *doing* lots of things."

## Concurrency in Go

* Go uses goroutines instead of OS threads
  * Lighter weight
  * Scheduled by Go runtime rather than OS kernel
  * Faster start-up times
  * Safe communication (no IPC necessary)
* Actor model
* Communicating Sequential Processes (CSP)
* Messages pass between Actors via Channels
* Channels are safe communication pipes between goroutines

## Writing a concurrent Go program

* Putting `go` in front of a function call causes it
to become a goroutine and execute in non-blocking mode.
* If `main()` exits **all GoRoutines immediately quit.**
  * Use the `sync` package's `WaitGroup` to make parent wait
  on backgrounded goroutines.  (See `concur.go`)

## Channels

* Can be buffered or unbuffered.
* To create use `make` with the `chan` keyword.

```
// unbuffered (blocking)
myChannel := make(chan int)

// buffered (non-blocking to size of buf)
myChannel2 := make(chan int, 5) // size of buf
```

* GoRoutines block when reading from a channel that contains
no data whether buffered or unbuffered

## More on channels

* Check out Mike van Sickle's "Concurrent Programming in Go" on Pluralsight
