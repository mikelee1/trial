package main

type a interface {
	Add() int
}


type b interface {
	Add() int
}

type Test struct {
	A int
	B int
}

func (t Test)Add() int {
	return t.A + t.B
}
