package main

import (
	"fmt"
)

func main() {
	fmt.Println(pathWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}))
}

type Point struct {
	X     int
	Y     int
	Value int
}

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var res [][]int
	if len(obstacleGrid) == 0 {
		return res
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 {
		return res
	}


	stack := []Point{
		{0, 0, obstacleGrid[0][0]},
	}
	handled := make([][]bool, m)
	for k, _ := range handled {
		handled[k] = make([]bool, n)
	}

	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		if tmp.Value == 1 {
			stack = stack[:len(stack)-1]
			continue
		}

		if tmp.Y+1 < n && !handled[tmp.X][tmp.Y+1] {
			p := Point{
				tmp.X, tmp.Y + 1, obstacleGrid[tmp.X][tmp.Y+1],
			}
			if p.X == m-1 && p.Y == n-1 {
				break
			}
			stack = append(stack, p)
			handled[p.X][p.Y] = true
			continue
		}

		if tmp.X+1 < m && !handled[tmp.X+1][tmp.Y] {
			p := Point{
				tmp.X + 1, tmp.Y, obstacleGrid[tmp.X+1][tmp.Y],
			}
			if p.X == m-1 && p.Y == n-1 {
				break
			}
			stack = append(stack, p)
			handled[p.X][p.Y] = true
			continue
		}
		stack = stack[:len(stack)-1]

	}

	for _, v := range stack {
		res = append(res, []int{
			v.X, v.Y,
		})
	}
	if len(stack) > 0 {
		res = append(res, []int{m - 1, n - 1})
		return res
	}
	if obstacleGrid[0][0] == 0 && m == 1 && n == 1 {
		return [][]int{[]int{0, 0}}
	}

	return res
}
