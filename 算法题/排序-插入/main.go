package main

import "fmt"

func main() {
	var arr = []int{6, 2, 22, 45, 1, 6, 8, 200, 56, 111}
	insertSort(arr)
	fmt.Println(arr)
}

//打麻将理牌
func insertSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
