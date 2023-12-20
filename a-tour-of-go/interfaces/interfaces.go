package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X float64
	Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	return (math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2)))
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	var a Abser
	v := Vertex{2, 3}
	f := MyFloat(-math.Sqrt2)

	a = f
	a = &v
	fmt.Println(a.Abs())
}
