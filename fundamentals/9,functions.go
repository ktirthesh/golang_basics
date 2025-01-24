package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func function_call_golang() {
	res := plus(12, 32)
	fmt.Println("result is plus is ", res)

	res = plusplus(13, 23, 11)
	fmt.Println("plusplus result is ", res)
}
