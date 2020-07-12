package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Test struct {
	String  *string `json:"string,omitempty"`
	Integer *int    `json:"integer,omitempty"`
}

func check(js []byte) {
	var t Test
	if err := json.Unmarshal(js, &t); err != nil {
		log.Fatal(err)
	}

	if t.Integer != nil {
		*t.Integer += 1
	}
	t.Integer = nil

	newJS, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(newJS))
}

func main() {
	check([]byte(`{"string":"this is a string","integer":0}`))
	check([]byte(`{"string":"this is a string"}`))
}
