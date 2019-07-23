package main

import "fmt"

func main() {
	fmt.Println(getE(10000000000))
}

func getE(n int) float64 {
	fmt.Println(float64(1) + float64(1)/float64(n))
	return pow(float64(1)+float64(1)/float64(n), n)
}

func pow(x float64, n int) float64 {
	ret := float64(1) // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
