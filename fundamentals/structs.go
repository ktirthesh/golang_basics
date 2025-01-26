package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

func struct_golang() {
	fmt.Println(person{name: "tirthesh"})
	fmt.Println(person{name: "amit", age: 22})
	fmt.Println(person{"fdfdf", 23})
	fmt.Println(&person{name: "coco", age: 21})
	fmt.Println(newPerson("john"))

	s := person{name: "alen", age: 34}
	fmt.Println("name is s", s)
	s1 := &s
	fmt.Println("age is ", s1.age)
	s1.age = 23
	fmt.Println("age is ", s1.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex", true,
	}
	fmt.Println("dog default is ", dog)
}
