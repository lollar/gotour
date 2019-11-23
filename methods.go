package main

import (
  "fmt"
  "math"
  "time"
  "strings"
  "io"
)

// Methods
// Go does not have classes but you can define methods on types
// receive appears in it's own argument list

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Methods continued
// A method can be declared on non-struct types
// You can only declare a method with a receiver whose type is
// defined in the same package as the method. (No monkeypatching <3)

type MyFloat float64

func (f MyFloat) FloatAbs() float64 {
  if f < 0 { return float64(-f) }

  return float64(f)
}

// Pointer receivers
// Methods can be declared with pointer recievers. Meaning the receiver type
// has the literal syntax *T where T cannot also be a pointer.  Methods with
// pointer receivers can modify the value. With a value receiver, a method
// operates on a copy of the original value. Pointer receivers must be used
// to operate on the original declared value being passed to the function

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

// Pointers and functions
// Pointers can also be used in an argument list for a function

func AbsFunc(v Vertex) float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

// Choosing a value or pointer receiver
// There are two reasons to use a pointer receiver:
// 1. So the method can modify the value that it's receiver points to
// 2. Avoid copying the value on each method call. Making it more efficient
//
// Methods should not have both a value receiver or a pointer receiver,
// only one

// Methods and pointer indirection
// Functions with a pointer argument must take a pointer
// methods with pointer receivers take either a value or a pointer
// The same can be said for the reverse direction. 
// Methods that take value types must receive a value of that type
// while methods with value receivers take either a value or a pointer

// Interfaces
// Interface type is a set of method signatures. A value of interface
// type can hold any value that implements those methods.

// Interfaces are implemented implicity
// A type implements an interface by implementing its methods. No explicit
// "implements" keyword is needed.

type I interface {
  M()
}

type T struct {
  S string
}
func (t *T) M() {
  fmt.Println(t.S)
}

// Interface values
// Can be thought of as a type of value and concrete type `(value, type)`

type F float64

func (f F) M() {
  fmt.Println(f)
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

// Interface values with nil underlying values
// If concrete value is nil the method wil lbe called with
// a nil receiver. Calling a method on a nil interface is
// a run-time error because there is no concrete type

// The empty interface
// An interface that specifies zero methods.

// Type assertions
// Provides access to an interface value's underlying concrete value.
// `t := i.(T)`
// Inteface value `i` assigns undlerying `T` value to the `t` variable
// You can test a type asserition
// `t, ok := i.(T)` // this assigns the underlying value and a boolean value
// that reports wheter the assertion succeeded.

func typeAssertions() {
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  // panic!
  // f = i.(float64)
}

// Type switches
// Construct that permits several type assertions in series.
// Like a case statement but specify types not values

func typeChecker(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("int: %v\n", v)
  case string:
    fmt.Printf("string: %q\n", v)
  default:
    fmt.Printf("man... I don't know %T\n", v)
  }
}

func do() {
  typeChecker(21)
  typeChecker("hello")
  typeChecker(true)
}

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// Errors
// Error state expressed with `error` values. The `error` type is a
// built-in interface. A nil `error` denotes success.
type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
  return &MyError {
    time.Now(),
    "you fucked up",
  }
}

// Readers
// the `io` package specifies the `io.Reader` interface
// the following code creates a `strings.Reader` and consumes
// its output 8 bytes at a time.
func readersExample() {
  r := strings.NewReader("Hello, Reader!")

  b := make([]byte, 8)

  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])
    if err == io.EOF {
      break
    }
  }
}

// Add methods from each section here to execute code
func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.FloatAbs())

  v.Scale(10)
  fmt.Println(v.Abs())

  v2 := Vertex{3,4}
  ScaleFunc(&v2, 10)
  fmt.Println(AbsFunc(v2))

  var i I = &T{"hello"}
  i.M()

  var j I

  j = &T{"Hello"}
  describe(j)
  j.M()

  j = F(math.Pi)
  describe(j)
  j.M()

  typeAssertions()
  do()

  a := Person{"Arthur Dent", 42}
  b := Person{"Bill Clinton", 67}
  fmt.Println(a,b)

  if err := run(); err != nil {
    fmt.Println(err)
  }

  readersExample()
}
