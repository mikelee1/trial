package main

func main() {
	//merge([]int{1, 2, 4, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
	merge([]int{2, 0}, 1, []int{1}, 1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	aTail := m - 1
	bTail := n - 1
	if m == 0 {
		for k, v := range nums2 {
			nums1[k] = v
		}
	}

	for aTail >= 0 && bTail >= 0 {
		if nums1[aTail] < nums2[bTail] {
			nums1[aTail+bTail+1] = nums2[bTail]
			bTail--
		} else {
			nums1[aTail+bTail+1] = nums1[aTail]
			aTail--
		}
	}
	if bTail >= 0 {
		for k, v := range nums2[:bTail+1] {
			nums1[k] = v
		}
	}
}
