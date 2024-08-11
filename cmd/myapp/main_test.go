package main

import (
  "testing"
)

func TestCalculateHypotenuse(t *testing.T) {
  tests := []struct {
    a, b float64
    expected float64
  }{
      {3, 4, 5},
      {5, 12, 13},
      {8, 15, 17},
      {7, 24, 25},
  }

  for _, tt := range tests {
    result := CalculateHypotenuse(tt.a, tt.b)
    if result != tt.expected {
      t.Errorf("CalculateHypotenuse(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
    }
  }

}
