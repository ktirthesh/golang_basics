package main

import (
	"fmt"
	"net"
	"net/url"
)

func url_parsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("User:", u.User)
	fmt.Println("Username:", u.User.Username())
	P, _ := u.User.Password()
	fmt.Println("Password:", P)

	fmt.Println("Host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)
	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawFragment:", u.RawFragment)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("parsed query:", m)
	fmt.Println(m["k"][0])
}
