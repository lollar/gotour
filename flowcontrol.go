package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

// For
// The only looping construct in Go

func forExample() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
    fmt.Println(sum)
  }
}

// For continued
// The init and post statements are optional :)

func forExampleCont() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}

// For is Go's "while"
// Semicolons in previous example aren't actually needed

// For(ever)
// `for { ... }`
// will enter an infinite loop

// If
// parentheses are not required but braces are
func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func powagain(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't access v here
  return lim
}

// Switch
// Only runs the selected case not all cases that follow. No `break` is needed

func switchySwitch() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    fmt.Printf("%s. \n", os)
  }
}

func untilSaturday() {
  fmt.Println("\nWhen's Saturday?")
  today := time.Now().Weekday()

  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away. :(")
  }
  fmt.Println("")
}

// Defer
// Defers the execution of a function until the surrounding function returns.
// When defers are stacked they are executed in LIFO order

func logger(x, y int) {
  fmt.Println(x + y)
}

// Defer, Panic, and Recover - Blog Post
// A deferred function's arguments are evaluated when the defer statement is evaluated.
// Deferred statements happen in LIFO order
// Deferred functions may read and assign to the return functions named return values
func c() (i int) {
  defer func() { i++ } ()
  return 1
}

// Panic is a built-in function that stops the ordinary flow of control and begins _panicking_
// When panic is called the function will stop, execute any deferred fucntions, and return

// Recover regains control of a panicking goroutine. It's only useful inside deferred fuctions.

func f() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("recovered in f", r)
    }
  }()
  fmt.Println("Calling g.")
  g(0)
  fmt.Println("Returned normally from g.")
}

func g(i int) {
  if i > 3 {
    fmt.Println("Panicking!")
    panic(fmt.Sprintf("%v", i))
  }

  defer fmt.Println("Defer in g", i)
  fmt.Println("Printing in g", i)
  g(i + 1)
}

// Add methods from each section here to execute code
func main() {
  forExample()
  forExampleCont()
  fmt.Println(sqrt(16), sqrt(-4))
  fmt.Println(pow(3,2,10), pow(3,3,20))
  fmt.Println(powagain(3,2,10), powagain(3,3,20))
  switchySwitch()
  untilSaturday()

  x := 9
  y := 10

  fmt.Println(c())

  f()
  fmt.Println("Returned from f normally.")

  // Defer section (has to run last)
  defer logger(x, y)
  fmt.Printf("%d + %d = ", x, y)
}
