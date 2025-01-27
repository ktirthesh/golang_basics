package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0

}

func pointers_golng() {
	i := 1

	fmt.Println("intial", i)
	zeroval(i)
	fmt.Println("using zeroval for i", i)

	zeroptr(&i)
	fmt.Println("using the zeroptr for i ", i)
	fmt.Println("the locaton of i is ", &i)
}
