package main

import "fmt"

func main() {
	var arr = []int{33, 2, 6, 4, 5}
	chooseSort(arr)
	fmt.Println(arr)
}

func chooseSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
