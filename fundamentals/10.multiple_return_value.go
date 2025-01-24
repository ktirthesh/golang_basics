package main

import "fmt"

func val() (int, int) {
	return 2, 3
}

func multiple_return_value_golang() {
	x, y := val()
	fmt.Println("multiple return values are x,y", x, y)
	_, z := val()
	fmt.Println("value of z and _ is ", z)
}
