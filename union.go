package main

import (
	"log"
	"sort"
)

func main() {
	array1 := []int{1, 5, 2, 8, 9, 4, 10}
	array2 := []int{6, 3, 5, 1}

	sort.Ints(array1)
	sort.Ints(array2)

	// 并集
	var union []int
	i, j := 0, 0
	for i < len(array1) || j < len(array2) {
		if i < len(array1) && j < len(array2) && array1[i] == array2[j] {
			union = append(union, array1[i])
			i++
			j++
		} else if j >= len(array2) || array1[i] < array2[j] {
			union = append(union, array1[i])
			i++
		} else {
			union = append(union, array2[j])
			j++
		}
	}

	log.Print(union)
}