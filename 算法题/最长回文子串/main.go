package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		maxLength := 0
		maxStr := ""
		for i := 0; i+maxLength < len(text); i++ {
			for j := len(text) - 1; j >= 0 && j >= i && j-i+1 > maxLength; j-- {
				if check(text, i, j) {
					maxLength = j - i + 1
					maxStr = text[i : j+1]
				}
			}
		}
		fmt.Println(maxStr)
	}
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
