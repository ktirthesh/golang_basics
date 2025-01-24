package main

import "fmt"

func for_loop() {
	fmt.Println("for loop ")
	i := 1
	for i < 3 {
		fmt.Println("i is ", i)
		i++
	}

	for j := 0; j < 8; j++ {
		fmt.Println("val of j is ", j)
	}

	for i := range 3 {
		fmt.Println("val od i", i)
	}
	for {
		fmt.Println("singe loop break  ")
		break
	}
	for n := range 8 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("value of n is ", n)
	}
}
