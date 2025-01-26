package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods_with_struct() {

	r := rect{height: 12, width: 4}
	fmt.Println("area is", r.area())
	fmt.Println("perimeter is ", r.perim())
	rp := &r
	fmt.Println("area is ", rp.area())
	fmt.Println("perimeter is ", rp.perim())
}
