package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func recover_golang() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recoverd error :\n", r)
		}
	}()
	mayPanic()
	fmt.Println("after maypanic")
}
