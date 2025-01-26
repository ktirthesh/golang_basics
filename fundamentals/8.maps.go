package main

import (
	"fmt"
	"maps"
)

func maps_golang() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 133
	m["k3"] = 123
	fmt.Println("value of m is ", m)

	v1 := m["k1"]
	fmt.Println("value of k1 is", v1)

	v4 := m["k4"]
	fmt.Println("value of k4 is", v4)

	fmt.Println("length of m is ", len(m))

	delete(m, "k2")
	fmt.Println("now m is  ", m)

	clear(m)
	fmt.Println("m is cleared :", m)

	_, prs := m["k2"]
	fmt.Println("prs is ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n is ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n2 is ", n2)

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
