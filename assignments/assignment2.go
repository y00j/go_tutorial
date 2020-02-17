package main

import "fmt"

type triangle struct{
	base float64
	height float64
}

type square struct{
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square {sideLength:1.9,}
	t := triangle {
		base: 2.0,
		height: 3.001,
	}

	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}