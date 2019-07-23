package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	grid := [][]byte{
		[]byte{1, 1, 0, 0, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 1, 0, 0},
		[]byte{0, 0, 0, 1, 1},
	}

	grid = [][]byte{
		[]byte{1, 1, 1, 1, 0},
		[]byte{1, 1, 0, 1, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0},
	}

	//grid = [][]byte{
	//	[]byte{1,1,1},
	//	[]byte{0,1,0},
	//	[]byte{1,1,1},
	//}
	fmt.Println(numIslands2(grid))
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	copyed := [][]byte{}
	gridData, _ := json.Marshal(grid)
	json.Unmarshal(gridData, &copyed)
	for k, v := range grid {
		for k1, _ := range v {
			fmt.Println(grid[k][k1])
			if grid[k][k1] == 0 {
				continue
			}
			if k1+1 < len(v) && grid[k][k1+1] == 1 {
				copyed[k][k1+1] = 0
			}
			if k+1 < len(grid) && grid[k+1][k1] == 1 {
				copyed[k+1][k1] = 0
			}
		}
	}
	sum := 0
	for _, v := range copyed {
		for _, v := range v {
			sum += int(v)
		}
	}

	return sum
}

func numIslands2(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	copyed := [][]byte{}
	gridData, _ := json.Marshal(grid)
	json.Unmarshal(gridData, &copyed)
	solvedMap := map[string]bool{}
	for k, v := range grid {
		for k1, _ := range v {
			fmt.Println(grid[k][k1])
			solvedMap[string(k)+"-"+string(k1)] = true

			if grid[k][k1] == byte(0) {
				continue
			}

			if k1+1 < len(v) && grid[k][k1+1] == 1 {
				copyed[k][k1+1] = 0
			}
			if k+1 < len(grid) && grid[k+1][k1] == 1 {
				copyed[k+1][k1] = 0
			}
			if k1-1 > -1 && grid[k][k1-1] == 1 && !solvedMap[string(k)+"-"+string(k1)] {
				copyed[k][k1-1] = 0
			}

		}
	}
	sum := 0
	for _, v := range copyed {
		fmt.Println(v)
		for _, v := range v {

			sum += int(v)
		}
	}

	return sum
}
