package main

import (
	"fmt"
)

var (
	str = "123"
)

func main() {
	arrange([]byte(str), func(a []byte) {
		fmt.Println(string(a))
	}, 0)
}

//打印第i开始的全排列
func arrange(a []byte, f func([]byte), i int) {
	if i > len(a) {
		f(a)
		return
	}
	//先打印一个第i+1开始的全排列
	arrange(a, f, i+1)
	for j := i + 1; j < len(a); j++ { //交换i和j元素后，打印一个第i+1开始的全排列，再交换i和j元素回去。
		a[i], a[j] = a[j], a[i]
		arrange(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

//arrange(a, f, 4) print a
//arrange(a, f, 3)   x
//arrange(a, f, 2)     x
//arrange(a, f, 1)        1<->2 arrange(a, f, 2) 1<->2
//arrange(a, f, 0)                                      0<->1 arrange(a, f, 1) 0<->1 0<->2 arrange(a, f, 1) 0<->2
