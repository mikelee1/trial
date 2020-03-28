package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}

func minimumLengthEncoding(words []string) int {
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if check(words[i], words[j]) {
				if len(words[i]) > len(words[j]) {
					words[j] = ""
				} else {
					words[i] = ""
				}
			}
		}
	}

	count := 0
	for _, word := range words {
		if word != "" {
			count += len(word) + 1
		}
	}

	return count
}

func check(a, b string) bool {
	var i, j = len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if a[i] != b[j] {
			return false
		}
	}
	return true
}
