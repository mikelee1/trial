package main

import (
	"fmt"
	"sort"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"bytes"
)

type Person struct {
	Name string // 姓名
	Age  int    // 年纪
}

// 按照 Person.Age 从大到小排序
type PersonSlice []Person

func (a PersonSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a PersonSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a PersonSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	if a[j].Age < a[i].Age{
		return true
	}

	aa, _ := UTF82GBK(a[i].Name)
	bb, _ := UTF82GBK(a[j].Name)
	bLen := len(bb)
	for idx, chr := range aa {
		if idx > bLen-1 {
			return false
		}
		if chr != bb[idx] {
			return chr < bb[idx]
		}
	}
	return true
}

func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

func main() {
	people := []Person{
		{"1", 12},
		{"2", 30},
		{"3", 52},
		{"黄飞鸿", 26},
		{"李世民", 26},
		{"牛根", 26},
		{"龙天下", 26},
	}
	//fmt.Println(people)

	sort.Sort(PersonSlice(people)) // 按照 Age 的逆序排序
	fmt.Println(people)

	//sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	//fmt.Println(people)
	//mapa = map[string]int{"a": 1, "b": 2}
	//for _, v := range mapa {
	//	fmt.Println(v)
	//}
	strslice := []string{"中国", "美国", "美人", "中国人"}
	fmt.Println(sort.StringSlice(strslice))
}

var mapa = make(map[string]int)
