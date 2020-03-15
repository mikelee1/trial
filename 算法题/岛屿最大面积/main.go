package main

import "fmt"

//var island = [][]int{
//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
//}

func main() {
	island := [][]int{
		{0, 1}, {1, 1},
	}
	fmt.Println(maxAreaOfIsland(island))
}

type Point struct {
	X int
	Y int
}

func maxAreaOfIsland(island [][]int) int {
	toHandle := []Point{}
	handled := map[Point]bool{}
	var max = 0
	for row, v := range island {
		for column, v1 := range v {
			if v1 == 0 {
				continue
			}
			p := Point{X: row, Y: column}
			if ok := handled[p]; ok {
				continue
			}

			toHandle = append(toHandle, p)
			handled[p] = true
			count := 0
			for len(toHandle) > 0 {
				count++
				p := toHandle[0]
				toHandle = toHandle[1:]
				pp := Point{p.X + 1, p.Y}
				if p.X+1 < len(island) && island[p.X+1][p.Y] == 1 && !handled[pp] {
					toHandle = append(toHandle, pp)
					handled[pp] = true
				}
				pp = Point{p.X, p.Y - 1}
				if p.Y-1 >= 0 && island[p.X][p.Y-1] == 1 && !handled[pp] {
					toHandle = append(toHandle, pp)
					handled[pp] = true
				}
				pp = Point{p.X, p.Y + 1}
				if p.Y+1 < len(island[0]) && island[p.X][p.Y+1] == 1 && !handled[pp] {
					toHandle = append(toHandle, pp)
					handled[pp] = true
				}

				pp = Point{p.X - 1, p.Y}
				if p.X-1 >= 0 && island[p.X-1][p.Y] == 1 && !handled[pp] {
					toHandle = append(toHandle, pp)
					handled[pp] = true
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
