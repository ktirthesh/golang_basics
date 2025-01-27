package main

import "fmt"

func sum(num ...int) {
	fmt.Println("parameters receiverd in sum function", num)
	total := 0
	for _, num := range num {
		total += num
	}
	fmt.Println("total", total)
}

func vardic_function_golang() {
	sum(12, 34)
	numsss := []int{1, 3, 45, 5, 3, 2, 1}
	sum(numsss...)
}
