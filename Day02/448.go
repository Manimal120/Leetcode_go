package main

func findDisappearedNumbers(nums []int) []int {
	var res []int
	for _, v := range nums {
		if v < 0 {
			v = -v
		}

		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}

	for _, v := range nums {
		if v < 0 {
			res = append(res, v)
		}
	}

	return res
}
