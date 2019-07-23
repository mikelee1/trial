package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := 10
	b := fmt.Sprintf("PC%010d", a)

	id := ""
	lis := []string{}
	json.Unmarshal([]byte(b[2:]), &lis)
	fmt.Println(lis)
	sign := true
	for _, v := range lis {
		if v == "0" && sign {
			continue
		}
		sign = false

		id = id + v
	}

	fmt.Println(id)
}
