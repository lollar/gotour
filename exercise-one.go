package main

import "fmt"

func Sqrt(x float64) float64 {
  z  := 1.0
  zz := 0.0

  for z != zz {
    zz = z
    z -= (z*z - x) / (2*x)
  }

  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
