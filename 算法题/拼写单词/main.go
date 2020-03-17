package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters(words []string, chars string) int {
	charsMap := map[string]int{}
	charsList := strings.Split(chars, "")
	for _, char := range charsList {
		charsMap[char] += 1
	}
	count := 0
	for _, word := range words {
		if check(word, charsMap) {
			count += len(word)
		}
	}
	return count
}

func check(word string, charsMap map[string]int) bool {
	cm := map[string]int{}
	for _, w := range strings.Split(word, "") {
		if _, ok := charsMap[w]; !ok || cm[w]+1 > charsMap[w] {
			return false
		}
		cm[w] += 1
	}
	return true
}
