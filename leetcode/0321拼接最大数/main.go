package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{
		3, 4, 6,
	}
	nums2 := []int{
		1, 2, 5, 8, 3, 9,
	}
	fmt.Println(maxNumber(nums1, nums2, 5))
}

func remainAndOrderCheck(v, k int, res []int, head1index, head2index *int, nums1, nums2 Nums) bool {
	if len(nums1)-*head1index+len(nums2)-*head2index < k-len(res) {
		return false
	}
	index1 := nums1.Indexof(v)
	index2 := nums2.Indexof(v)
	if index1 == -1 && index2 == -1 {
		return false
	}
	if index1 != -1 {
		*head1index = index1
		return true
	}
	if index2 != -1 {
		*head2index = index2
	}
	return true

}
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := []int{}
	head1index := 0
	head2index := 0
	//nums1 nums2整体排序
	var totalnums Nums
	totalnums = append(nums1, nums2...)

	sort.Sort(totalnums)
	fmt.Println(totalnums)

	//从大到小依次检查
	for _, v := range totalnums {
		//检查该位选择以后，会不会位数不足;检查该位的位置是否符合顺序
		if !remainAndOrderCheck(v, k, res, &head1index, &head2index, nums1, nums2) {
			continue
		}

		fmt.Println(v)
		res = append(res, v)
		//如果位数达到要求则退出
		if len(res) > k-1 {
			break
		}
	}
	return res
}

type Nums []int

func (n Nums) Len() int {
	return len(n)
}

func (n Nums) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n Nums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nums) Indexof(v int) int {
	for k, v1 := range n {
		if v1 == v {
			return k
		}
	}
	return -1
}
