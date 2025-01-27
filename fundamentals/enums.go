package main

import "fmt"

type ServerState int

const (
	StateIdel ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdel:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func enum_golang() {
	ns := transition(StateIdel)
	fmt.Println("state of ns is", ns)
	ns1 := transition(ns)
	fmt.Println("state of ns is", ns1)

}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdel:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdel
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("UnknownError :  %s", s))
	}
}
