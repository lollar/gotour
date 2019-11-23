package main

import(
  "fmt"
  "math"
  "strings"
)

// Pointers
// Zero value is nil, the `&` operatores generates a pointer to its operand
// The `*` operator denotes the pointer's underlying value
func pointers() {
  i, j := 42, 2701

  p := &i          // point to i
  fmt.Println(*p)  // read i through the pointer
  *p = 21          // set i through the pointer
  fmt.Println(i)   // see the new value of i

  p = &j           // point to j
  *p = *p / 37     // divide j through the pointer
  fmt.Println(j)   // see the new value of j
}

// Structs
type Vertex struct {
  X int
  Y int
}

// Struct Fields
// Accessed using a `.`
func vertex() {
  v := Vertex{1,2}
  v.X = 4
  fmt.Println(v.X)
}

// Pointers to structs
// Fields can be added via a pointer
// long hand notation `(*p).X` but the language allows us to write `p.X`
func pointersToStructs() {
  v := Vertex{1,2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}

// Struct Literals
// Denotes a newly allocated struct value by listing the values of it's fields

var(
  v1 = Vertex{1, 2}
  v2 = Vertex{X: 1} // Y:0 is implict
  v3 = Vertex{}     // X:0, Y:0 is implicit
  p  = &Vertex{1, 2}
)

// Arrays
// Denoted as `[n]T` where n is size and T is type
// Arrays length is part of it's type so cannot be resized

func arr() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)
}

// Slices
// Much more common than arrays. Dynamically sized and flexible.
// Formed by specifying two indices, a low and high bound.
// Similar to references to arrays, it does not store any data, just
// describes the defined section of an array. Changing elements of a
// slice will update the corresponding elements of the array.

func slices() {
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println(s)
}

func arrSlice() {
  names := [4]string{ "John", "Paul", "George", "Ringo" }
  fmt.Println(names)

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)
}

// Slice literals
// Creates the same array, the builds a slice that references it

func literallySliced() {
  q := []int{2,3,5,7,11,13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)
}

// Slice defaults
// can omit the high or low bounds to use defaults instead
// defaults: low bound - 0 & high bound - length of slice

func sliceDefaults() {
  s := []int{2,3,5,7,11,13}

  s = s[:]
  fmt.Println(s)

  s = s[:2]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)

}

// Slice length & capacity
// Lenghth = number of elements it contains.
// Capacity = Number of elements in the underlying array
// `len(slice)` & `cap(slice)` do determine values

// Nil slices
// Zero value of a slic eis nil. It has a lenght and capacity of 0 and no underlying array

// Creating a slice with make
// Slices can be created with the built-in fuction `make`. This is how you create dynamically-sized arrays.
// `make` allocates a zeroed array and returns a slice that refers to the array.
// `a := make([]int, 5) // len(a) = 5
func makeSlice() {
  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Slice of slices
// Slices can contain any type, including other slices
func sliceOfSlices() {
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}

// Appending to a slice
// Use built-in append function. First param is a slice of type T,
// the rest of the params are T values to append. A newly allocated
// array will be provided if original array is too small
func appendToSlice() {
  var s []int
  printSlice("s", s)

  // append works on nil slices
  s = append(s, 0)
  printSlice("nil slice", s)

  s = append(s, 2, 3, 4)
  printSlice("more than one", s)
}

// Range
// Range form of the for loop iterates over a slice or map
// When ranging over a slice two values are returned the
// index, and the underlying value

func ranging() {
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}

// Range cont'd
// You can skip index or value by assigning to `_`
// And if you only want index you can omit the second variable

func ratm() {
  pow := make([]int, 10)
  for i := range pow {
    pow[i] = 1 << uint(i)
  }

  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }
}

// Maps
// Maps keys to values. Zero value is nil, and keys cannot be added.
// The make function returns a map of type T that is initialized

type Coordinates struct {
  Lat, Long float64
}

func mappyMap() {
  m := make(map[string]Coordinates)
  m["Bell Labs"] = Coordinates{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}

// Map literals
// Like struct literals but require keys
// You can also omit type from elements of the literal

var ma = map[string]Coordinates{
  "Bell Labs": Coordinates{
    40.68433, -74.39967,
	},
	"Google": Coordinates{
		37.42202, -122.08408,
	},
}

func printCoordinates() {
  fmt.Println(ma)
}

// Mutating maps
// Insertion `m[key] = elem`
// Retrieval `elem = m[key]`
// Deletion  `delete(m, key)`
// Presence  `elem, ok := m[key]`

func mapMutation() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("Value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("Value:", m["Answer"])

  delete(m, "Answer")
  fmt.Println("Value:", m["Answer"])

  v, ok := m["Answer"]
  fmt.Println("Value:", v, "Present?", ok)
}

// Function values
// Functions can be passed around like any other value
// They be used as arguments and return values

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func functioning() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  fmt.Println(hypot(5, 12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}

// Function closures
// A closure is a function value that references variables
// from outside its body.

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func adderCaller() {
  pos, neg := adder(), adder()

  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-2*i))
  }
}

// Add methods from each section here to execute code
func main() {
  pointers()
  fmt.Println(Vertex{1,2})
  vertex()
  pointersToStructs()
  fmt.Println(v1, p, v2, v3)
  arr()
  slices()
  arrSlice()
  fmt.Println("\n")
  literallySliced()
  sliceDefaults()
  makeSlice()
  sliceOfSlices()
  appendToSlice()
  ranging()
  ratm()
  mappyMap()
  printCoordinates()
  mapMutation()
  functioning()
  adderCaller()
}
