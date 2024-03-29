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
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	println(f)
}
func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()
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
