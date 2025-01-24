package main

import (
	"fmt"
	"time"
)

func switch_golang() {
	fmt.Println("the switch statement is called")
	// switch case with int
	i := 2
	switch i {
	case 1:
		fmt.Println("switch 1 -> i is ", i)
	case 2:
		fmt.Println("switch 1 -> i is ", i)
	case 3:
		fmt.Println("switch 1 -> i is ", i)
	}

	// switch case with possible cases

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("today is Monday")
	case time.Monday:
		fmt.Println("today is Monday")
	case time.Tuesday:
		fmt.Println("today is Tuesday")
	case time.Wednesday:
		fmt.Println("today is Wednesday")
	case time.Thursday:
		fmt.Println("today is Thursday")
	case time.Friday:
		fmt.Println("today is Friday")
	case time.Saturday:
		fmt.Println("today is Saturday")
	}

	// switch case  condition and default
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("1.before the noon")
		fmt.Println("2.before the noon")
	default:
		fmt.Println("after the noon")
	}

	// switch case with switch has  multiple cases or is present
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("this is weekend ")
	default:
		fmt.Println("this is not weekend")
	}

	// switch case work as function
	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("ths value is bool", i, t)
		case int:
			fmt.Println("this is integer", i, t)
		default:
			fmt.Printf("the data type i is -> %T\n", t)
		}

	}

	whatami("tirthesh")
	whatami(1)
}
