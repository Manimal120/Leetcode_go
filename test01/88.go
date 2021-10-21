package main

// 延长l1的屁股， 谁大加后面
// 把剩下的n放进开头去

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[p] = nums1[m-1]
			m--
		} else {
			nums1[p] = nums2[n-1]
			n--
		}
	}
	for i := 0; i < n; i++ {
		nums1[n-1] = nums2[n-1]
	}
}
