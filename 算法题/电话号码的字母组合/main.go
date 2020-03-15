package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCombinations("234"))
}

var m = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := []string{}
	di := strings.Split(digits, "")
	if len(di) == 0 {
		return res
	}
	head := di[0]
	di = di[1:]
	res = m[head]

	for len(di) > 0 {
		head = di[0]
		di = di[1:]
		res = generate(res, m[head])
	}
	return res
}

func generate(a, b []string) (res []string) {
	for _, v1 := range a {
		for _, v2 := range b {
			res = append(res, v1+v2)
		}
	}
	return
}
