package main

func subsets(nums []int) [][]int {
	inProcess, res := make([]int, 0), make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		dfsProcess1(nums, i, 0, inProcess, &res)
	}
	return res
}

func dfsProcess1(nums []int, i, index int, inProcess []int, res *[][]int) {
	if len(inProcess) == i {
		temp := make([]int, len(inProcess))
		copy(temp, inProcess)
		*res = append(*res, temp)
		return
	}

	for k := index; k < len(nums)-(i-len(inProcess))+1; k++ {
		inProcess = append(inProcess, nums[k])
		dfsProcess1(nums, i, index+1, inProcess, res)
		inProcess = inProcess[:len(inProcess)-1]
	}
	return
}
