package main

import (
	"fmt"
	"math"
)

func main() {
  
  var x,y float64
  
  for {

    fmt.Print("Please enter the value of the first side: ")
    _, err := fmt.Scan(&x)

    if err != nil {
      fmt.Println("Error... Please enter a valid float")
      var discard string
      fmt.Scanln(&discard)
    } else if x <= 0{
       fmt.Println("Error... Please enter a positive float") 
    } else {
        break
    }
  }
  for {
      fmt.Print("Please enter the value of the second side: ")
      _, err := fmt.Scan(&y)

      if err != nil {
        fmt.Println("Error... Please enter a valid float")
        var discard string
        fmt.Scanln(&discard)
      } else if x <= 0{
          fmt.Println("Error... Please enter a positive float") 
      } else {
          break
      }
  }    
  
  hyp := calculateHypotenuse(x, y)

  fmt.Printf("The hypotenuse is %v!", hyp)

}

func calculateHypotenuse(x, y float64) float64 {
  return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
