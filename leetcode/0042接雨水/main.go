package main

import "fmt"

func main() {
	data := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//data := []int{5,2,1,2,1,5}
	fmt.Println(trap(data))
}
func intTrim0(a []int) []int {
	i := 0
	for ; i < len(a); i++ {
		if a[i] == 0 {
			continue
		}
		break
	}
	j := len(a) - 1
	for ; j > -1; j-- {
		if a[j] == 0 {
			continue
		}
		break
	}
	fmt.Println("intreim:", a[i:max(j+1, len(a))])
	return a[i:max(j+1, len(a))]
}

func LocalMax(l []int) []int {
	lm := []int{}
	lmdata := []int{}
	for i := 0; i < len(l); i++ {
		if i == 0 {
			if l[i] > l[i+1] {
				lmdata = append(lmdata, l[i])
				lm = append(lm, i)
			}
			continue
		}
		if i == len(l)-1 {
			if l[i] > l[i-1] {
				lmdata = append(lmdata, l[i])
				lm = append(lm, i)
			}
			continue
		}
		if l[i] >= l[i+1] && l[i] >= l[i-1] {
			lmdata = append(lmdata, l[i])
			lm = append(lm, i)
		}
	}
	fmt.Println("localmax:", lm)
	fmt.Println("lmdata:", lmdata)
	lm = check(lm, lmdata)
	return lm
}

func check(lm, lmdata []int) []int {
	res := []int{}
	tmpres := []int{}
	big := 0
	for key, value := range lmdata {
		big = max(value, big)
		if value == big {
			tmpres = []int{}
			res = append(res, lm[key])
			continue
		} else {
			tmpres = append(tmpres, lm[key])
		}
	}
	res = append(res, tmpres...)
	fmt.Println("check:", res)
	return res
}

func trap(height []int) int {
	height = intTrim0(height)
	if len(height) < 3 {
		return 0
	}
	localMax := LocalMax(height)

	res := countRain(height, localMax)
	return res
}

func countRain(height, localMax []int) int {
	c := 0
	if len(localMax) < 2 {
		return 0
	}
	for i := 0; i < len(localMax)-1; i++ {
		min := min(height[localMax[i]], height[localMax[i+1]])
		fmt.Println(localMax[i], localMax[i+1])
		for j := localMax[i]; j < localMax[i+1]; j++ {
			c += max(min-height[j], 0)
			fmt.Println("add:", max(min-height[j], 0))
		}
	}
	return c
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
