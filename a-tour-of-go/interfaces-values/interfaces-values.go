package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	println(t.S)
}

type F float64

func (f F) M() {
	println(f)
}
func main() {
	var i I
	i = &T{"Hello, World!"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n", i, i)
}
