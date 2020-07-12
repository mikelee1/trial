package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []Person{
		Person{Name:"lee",Age:1},
		Person{Name:"mike",Age:3},
		Person{Name:"mike",Age:2},
	}
	fmt.Println(a)
	sort.Sort(PersonSlice(a))
	fmt.Println(a)
	fmt.Println(a)
}

type Person struct {
	Name string
	Age int
	Num float64
}


type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
//按年龄降序排列
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	if a[i].Age > a[j].Age{
		return true
	}
	return false
}