package main

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxJump := 0
	for index, value := range nums {
		if index > maxJump {
			return false
		}
		maxJump = max(maxJump, value+index)
	}
	return true
}
