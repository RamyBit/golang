package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}
func main() {
	v := Vertex{2, 3}
	v.Scale(10)
	fmt.Println(v.Abs())
}
