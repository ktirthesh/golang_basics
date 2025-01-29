package main

import (
	"fmt"
	"slices"
)

func slices_sorting_golang() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings", strs)
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("ints:", ints)
	s := slices.IsSorted(ints)
	fmt.Println("isSorted", s)
}
