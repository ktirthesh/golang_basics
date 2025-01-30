package main

import (
	"os"
	"text/template"
)

func text_template_golang() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("value: {{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"go", "Rust", "c++", "c#",
	})
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{"john doe"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "micky mouse",
	})
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not Empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4",
		"Range : {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"go",
			"Rust",
			"C++",
			"C#",
		})
}
