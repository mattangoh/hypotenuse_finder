package main

import (
	"fmt"
	"math"
)

func main() {
  
  var x,y float64

  fmt.Print("Please enter the value of the first side: ")
  fmt.Scan(&x)

  fmt.Print("Please enter the value of the second side: ")
  fmt.Scan(&y)

  hyp := calculateHypotenuse(x, y)

  fmt.Printf("The hypotenuse is %v!", hyp)

}

func calculateHypotenuse(x, y float64) float64 {
  return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
