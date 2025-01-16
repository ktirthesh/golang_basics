package main

import "fmt"

func if_else() {
	fmt.Println("-------------------the if else---------------------------- ")

	if len("hello") > 1 {
		fmt.Println("contion is called (only if is called)")
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("the 8 or 7 is even ")
	} else {
		fmt.Println("the 8 or 7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println("0 is less than ", num)
	} else if num < 5 {
		fmt.Println("5 is less than  ", num)
	} else if num < 10 {
		fmt.Println("10 is less than  ", num)
	}

}
