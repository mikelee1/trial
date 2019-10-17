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
	for j := i + 1; j < len(a); j++ {//交换i和j元素后，打印一个第i+1开始的全排列，再交换i和j元素回去。
		same := false
		for k := i + 1; k < j; k++{
			if a[i] == a[j]||a[k]==a[j]{
				same = true
			}
		}
		if same || a[i] == a[j]{
			continue
		}


		a[i], a[j] = a[j], a[i]
		arrange(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
