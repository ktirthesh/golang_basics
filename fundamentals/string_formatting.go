package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func string_formatting_golang() {
	p := point{1, 2}
	fmt.Printf("struct1 %v\n", p)
	fmt.Printf("struct2 %+v\n", p)
	fmt.Printf("struct3 %#v\n", p)
	fmt.Printf("struct 4 : %T\n", p)
	fmt.Printf("bool  : %t\n", true)
	fmt.Printf("int : %d\n", 123)
	fmt.Printf("bin: %b\n", 12)
	fmt.Printf("char: %c\n", 33)
	fmt.Printf("hex: %x\n", 456)
	fmt.Printf("float: %e\n", 1234.23)
	fmt.Printf("flaot2: %E\n", 1234.23)
	fmt.Printf("string: %s\n", "\"string\"")
	fmt.Printf("string2 with double quote: %q\n", "\"string\"")
	fmt.Printf("str in hex: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 2334)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.4556)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.455677)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "bar")
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "bar")
	s := fmt.Sprintf("sprintf a: %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "io: as %s \n", "error")
}
