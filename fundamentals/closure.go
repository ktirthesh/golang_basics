package main

/*
the main purpose of closure is function which return anamous function
*/
import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func closures_golang() {
	nextfunc := intSeq()
	fmt.Println(nextfunc())
	fmt.Println(nextfunc())
	fmt.Println(nextfunc())
	newfunc := intSeq()

	fmt.Println("new function ", newfunc())
}
