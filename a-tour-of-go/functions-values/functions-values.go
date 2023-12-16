package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64) float64 {
	return fn(2, 3)
}

func main() {
	fmt.Println(compute(math.Pow))
}
