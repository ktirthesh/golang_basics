package main

import (
	"fmt"
	"math"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %v", b.num)
}

type Container struct {
	base
	str string
}

func struct_embeding_golang() {
	co := Container{
		base: base{
			num: 12,
		},
		str: "some name",
	}
	fmt.Printf("co -{num %v str %v }\n", co.num, co.str)
	fmt.Printf("also the base num in co %v\n", co.base.num)
	fmt.Println("describe of num is ", co.describe())

	type describer interface {
		describe() string
	}
	fmt.Println(math.MaxInt64)
	var d describer = co
	fmt.Println("describer", d.describe())
}
