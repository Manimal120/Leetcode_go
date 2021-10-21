package main

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	maxJump := 0

	for k, v := range nums {
		if k > maxJump {
			return false
		}
		maxJump = max(maxJump, k+v)
	}
	return true
}
