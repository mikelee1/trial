package main

import "fmt"

func main() {
	fmt.Println(long([]string{
		"a123c",
		"a1",
		"a123d",
	}))
}

func long(strs []string) string {
	var res string
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res

}
