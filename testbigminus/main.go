package main

import (
	"fmt"
	"math/big"
)

func main() {
	res, err := BigMinus("2100010100000000000000", "2020011011373000932851")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func BigMinus(a, b string) (string, error) {

	n := new(big.Int)
	n, ok := n.SetString(a, 10)
	if !ok {
		fmt.Println("SetString: error")
		return "", fmt.Errorf("Fail to set string ")
	}

	n1 := new(big.Int)
	n1, ok = n1.SetString(b, 10)
	if !ok {
		fmt.Println("SetString: error")
		return "", fmt.Errorf("Fail to set string ")
	}

	n2 := new(big.Int)
	n2.Sub(n, n1)
	return n2.String(), nil
}
