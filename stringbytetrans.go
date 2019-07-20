package main

import "fmt"

func main() {
	a := [][]byte{[]byte("aaa"),[]byte("bbb")}
	b := []string{"bbb","ccc"}
	fmt.Println(stringtobyte(b))
	fmt.Println(bytetostring(a))

}

func stringtobyte(b []string) [][]byte {
	var res [][]byte
	for _,v := range b{
		byte1 := []byte(v)
		res = append(res,byte1)
	}
	return res
}

func bytetostring(b [][]byte) []string {
	var res []string
	for _,v := range b{
		s1 := string(v[:])
		fmt.Println(s1)
		res = append(res,s1)
	}
	return res
}
