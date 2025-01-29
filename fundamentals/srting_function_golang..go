package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func string_function_golang() {
	p("contains:", s.Contains("test", "es"))
	p("count:", s.Count("tests", "t"))
	p("hasprefix:", s.HasPrefix("tests", "te"))
	p("HasSuffix:", s.HasSuffix("tests", "ts"))
	p("Index", s.Index("test", "s"))
	p("JOin", s.Join([]string{"a", "b", "v"}, "-"))
	p("reapeat", s.Repeat("a", 5))
	p("replace", s.Replace("foo", "o", "s", -1))
	p("replace", s.Replace("foo", "o", "d", 1))
	p("split", s.Split("a-b-c-d-e", "-"))
	p("tolower", s.ToLower("TEST"))
	p("toupper", s.ToUpper("test"))

}
