package main

import "fmt"

func Sices_example() {
	fmt.Println("the slics in golang")
	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	// used to init slice map,chan object
	s = make([]string, 3)
	fmt.Println("emp", s, "len", len(s), "cap", cap(s))

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	fmt.Println("s", s)
	fmt.Println("get", s[2])

	s = append(s, "d")
	fmt.Println("s", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of data is ", c)

	l := s[2:5]
	fmt.Println("first silce of s is", l)

	l = s[:5]
	fmt.Println("second silce of s is", l)

	l = s[3:]
	fmt.Println("third silce of s is", l)

	t1 := []string{"a", "d", "z"}
	fmt.Println("hello", t1)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		fmt.Println("innerlen is ", innerlen)
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD is ", twoD)
}
