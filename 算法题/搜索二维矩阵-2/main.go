package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}},
		18))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, column := 0, len(matrix[0])-1
	for row < len(matrix) && column >= 0 {
		if matrix[row][column] == target {
			return true
		}
		if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}
