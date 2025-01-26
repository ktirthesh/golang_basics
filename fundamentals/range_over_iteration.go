package main

import (
	"fmt"
	"iter"
	"slices"
)

type Lisst[T any] struct {
	head, tail *elements[T]
}
type elements[T any] struct {
	next *elements[T]
	val  T
}

func (lst *Lisst[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &elements[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &elements[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst Lisst[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func range_over_iteration_golang() {
	lst := Lisst[int]{}
	lst.Push(12)
	lst.Push(23)
	lst.Push(45)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all", all)
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}

}
