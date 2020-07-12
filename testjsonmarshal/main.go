package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	a := A{
		"lee",
		20,
	}
	aByte, _ := json.Marshal(a)
	b := B{}
	json.Unmarshal(aByte, &b)
	fmt.Println(b)

	r := Receipt{
		ID:        "iii",
		OwnerComp: "kdjf",
	}
	rBytes, _ := json.Marshal(r)
	fmt.Println(string(rBytes))
	c := "{\"ID\":\"iii\",\"ownerComp\":\"kdjf\"}"
	receipt := Receipt{}
	json.Unmarshal([]byte(c), &receipt)
	fmt.Println(receipt)
}

type A struct {
	FirstName string `json:"first_name"`
	Age       int
}

type B struct {
	FirstName string `json:"first_name"`
	Age       int
}

type Receipt struct {
	ID        string //编号
	OwnerComp string //货主单位
}
