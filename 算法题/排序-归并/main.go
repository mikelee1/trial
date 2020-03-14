package main

import (
	"fmt"
)

func main() {
	var arr = []int{6, 2, 22, 45, 1, 6, 8, 200, 56, 111}

	fmt.Println(mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return Combine(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func Combine(arr1, arr2 []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	if i == len(arr1) {
		res = append(res, arr2[j:]...)
	}
	if j == len(arr2) {
		res = append(res, arr1[i:]...)
	}
	return res
}
