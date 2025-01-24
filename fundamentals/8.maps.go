package main

import (
	"fmt"
	"maps"
)

func Maps_golang() {

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 12
	fmt.Println("map in golang ->", m)
	v1 := m["k1"]
	fmt.Println("value of v1 is ", v1)
	v3 := m["k3"]
	fmt.Println("value of k3 is ", v3)

	delete(m, "k2")
	fmt.Println("m", m)
	clear(m)
	fmt.Println("after clear m", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n is ", n)

	n1 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n) {
		fmt.Println("n and n1 are equal")
	}
}
