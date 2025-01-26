package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type react struct {
	width, height float64
}

type circle struct {
	radious float64
}

func (r react) area() float64 {
	return r.width * r.height
}
func (r react) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radious
}

func measures(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radious", c.radious)
	}
}

func interfaces_golang() {
	r := react{width: 10.001, height: 20.001}
	c := circle{radious: 40.9034}
	measures(r)
	measures(c)
	detectCircle(r)
	detectCircle(c)
}
