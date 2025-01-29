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
}
