package main_test

import (
	"testing"
	"fmt"
	"time"
)

var text = "uaauuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuaau"

func Test_main(t *testing.T) {
	start := time.Now()
	maxStr := ""
	for i := 0; i < 100000; i++ {
		maxLength := 0
		for i := 0; i+maxLength < len(text); i++ {
			for j := len(text) - 1; j > 0 && j > i && j-i+1 > maxLength; j-- {
				if check(text, i, j) {

					maxLength = j - i + 1
					maxStr = text[i : j+1]

				}
			}
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())
	fmt.Println(maxStr)

}

func check(input string, start, end int) bool {
	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}
