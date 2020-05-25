package main

import (
	"fmt"
	"math"
)

func MySqrtWrap(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Println(MySqrtWrap(2))
}
