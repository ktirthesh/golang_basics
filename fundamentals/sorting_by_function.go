package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting_by_function_golang() {
	fruits := []string{"peach", "banana", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len((b)))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}
	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "Raj", age: 45},
		Person{name: "ravi", age: 34},
	}
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println("people", people)
}
