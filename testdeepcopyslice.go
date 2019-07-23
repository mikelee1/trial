package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := [][]int{[]int{1, 2, 3}}
	adata, _ := json.Marshal(a)

	b := [][]int{}
	json.Unmarshal(adata, &b)
	b[0][0] = 4
	fmt.Println(a, b)
}
