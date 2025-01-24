package main

import (
	"fmt"
	"math"
)

func constants() {
	fmt.Println("constants")
	const n = 200330
	const d = 3e20 / n
	fmt.Println("n and d", n, d)
	// fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
