package main_test

import (
	"testing"
	"fmt"
)

func TestAppend(t *testing.T) {
	for _, v := range [1000]int{}{
		go func(v int) {
			Fuck(v)
		}(v)
	}
}
//全局变量在并发时会有bug，要加锁
//var a []int
func Fuck(i int) []int {
	//自己命名空间内的变量不存在并发写入的问题
	a := []int{}
	for _, v := range []int{1,2,3}{
		a = append(a, v)
	}
	if len(a) != 3{
		fmt.Println(a)
	}
	return a
}