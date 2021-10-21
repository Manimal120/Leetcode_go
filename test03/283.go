package main

func moveZeroes(nums []int) {

	temp := make([]int, len(nums))
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp[cur] = nums[i]
			cur++
		}
	}
	copy(nums, temp)
}

func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {

		}
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}
