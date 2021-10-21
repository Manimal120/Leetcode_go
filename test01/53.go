package main

//dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-1], nums[i])
		res = max(res, dp[i])
	}
	return res
}
