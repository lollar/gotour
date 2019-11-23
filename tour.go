// Getting Started
package main

import(
  "fmt"
  "math"
  "math/rand"
  "math/cmplx"
)

func helloWorld() {
  fmt.Printf("hello, world\n")
}

// Package
// The Package name is the same as the last element of the import path
// i.e. "math/rand" (see import above) is made up of files with `package rand`

func randomNumber() {
  fmt.Println("My favorite number is", rand.Intn(10))
}

// Imports
// Imports written as above are called "factored" import statements.
// They can also be written as:
// import "fmt"
// import "math"
//
// But the factored import statements are considered good style.

func myProblems() {
  fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}

// Exported names
// Name is exported if it begins with a capital letter.
// When importing a package, only exported names can be referenced, for example:
// ```
// func main() {
//   fmt.Println(math.pi)
// }
// ```
// This will raise an error `cannot refer to unexported name math.pi`

// Functions
// A function can take zero arguments or arguments with type. The type comes after the variable name.
// The return type is listed after the function signature.

func goAdd(x int, y int) int {
  return x + y
}

// Functions continue
// When consecutive parameters have the same type you can omit the type from all but the last. For example:
// ```
// func goAdd(x, y int) int { ... }

// Multiple results
// Functions can return any number of results.

func lastNameFirst(firstName, lastName string) (string, string) {
  return lastName, firstName
}

// Named return values
// return values in the function signature maybe named. Setting the variables within the function body
// will allow for a "naked" return. As a best practice nake returns should only be used in short functions

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

// Variables
// `var` declatres list of variables with type listed last

var c, python, java bool

// Variables with Initializers
// See example in main below

var k, l int = 1, 2

// Short variable declarations
// The `:=` symbol can be used in place of `var` with implicit type in a function body.
// Outside of a function it is not available.

func shortVariable() {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i,j,k,c,python,java)
}

// Basic Types
// bool string int[8-64] uint[8-64] byte(alias for uint8)
// rune(alias for int32, represents Unicode code point)
// float[32-64] complex[64-128]
// int uint and uintptr will be 32 bits on 32 bit system and 64 on 64 bit system

var (
  ToBe   bool       = false
  MaxInt uint64     = 1<<64-1
  z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func printBasicTypes() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Zero values
// When a variable is declared without an inital value it will be given it's "zero value"

func zeroValues() {
  var i int     // 0
  var f float64 // 0
  var b bool    // false
  var s string  // ""
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Type conversions
// T(v) converts value `v` to type `T`

func typeConversions() {
  x, y := 3, 4
  f    := math.Sqrt(float64(x*x + y*y))
  z    := uint(f)

  fmt.Println(x, y, z)
}

// Type inference
// When using var or := and no type is provided, the type will be inferred from the value on the right hand side

// Constants
// Constants are declared like variables but with the `const` keyword. They can consist of char, string, bool
// or numeric values. They cannot be declared using `:=` syntax

const MyPi = 3.14

func tourOfConstants() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", MyPi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}

// Numeric Constants
// High-precision values.

const (
  Big = 1 << 100
  Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func printNumConsts() {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}

// Add methods from each section here to execute code
func main() {
  helloWorld()
  randomNumber()
  myProblems()
  fmt.Println(goAdd(42,13))
  fmt.Println(lastNameFirst("mike", "lollar"))
  fmt.Println(split(17))

  // Variables example
  var i, j int
  fmt.Println(i,j,c,python,java)

  // Variables with Initializers
  var golang, ruby, perl = true, false, "hell no!"
  fmt.Println(k, l, golang, ruby, perl)

  shortVariable()
  printBasicTypes()
  zeroValues()
  typeConversions()
  tourOfConstants()
  printNumConsts()
}
