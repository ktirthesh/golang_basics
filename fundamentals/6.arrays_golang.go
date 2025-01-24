package main

import "fmt"

func Arays() {
	fmt.Println("the array in golang")
	var a [5]int
	fmt.Println("the intial value of a ", a)
	a[4] = 12
	fmt.Println("the updated value of a is ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("initla value of b is ", b)

	b = [...]int{1, 2, 4, 6, 3}
	fmt.Println("the updated value of b is ", b)

	c := [...]int{1, 6: 12, 300, 1, 3, 4}
	fmt.Println("the updated value of c is ", c)
	var two [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println("position and value of two", i, j, two[i][j])
		}
	}

	fmt.Println("value of two is ", two)

	twoD := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("twoD is ", twoD)
}
