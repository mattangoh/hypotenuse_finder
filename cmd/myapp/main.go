package main

import (
	"fmt"
	"math"
)

func main() {
  
  x := 3.0
  y := 4.0

  hyp := calculateHypotenuse(x, y)

  fmt.Printf("The hypotenuse is %v", hyp)

}

func calculateHypotenuse(x, y float64) float64 {
  return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
