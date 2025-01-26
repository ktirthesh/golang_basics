package main

import (
	"fmt"
	"unicode/utf8"
)

func string_and_rums() {
	const s = "สวัสดี"
	fmt.Println("s is ", s)
	fmt.Println("length of s is ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count", utf8.RuneCountInString(s))
	for idx, runeValue := range s {
		fmt.Println("idx ad runevalue is - ", idx, runeValue)
	}
	fmt.Println("using DecodeRuneinString")

	for i, w := 0, 0; i < len(s); i += w {
		runevalue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runevalue, i)
		w = width

		exmineRune(runevalue)
	}
}

func exmineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
