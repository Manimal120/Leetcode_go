package main

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	used, inProcess, res := make([]bool, len(nums)), make([]int, 0), make([][]int, 0)

	dfsProcess(nums, 0, inProcess, &res, &used)

	return res
}

func dfsProcess(nums []int, index int, inProcess []int, res *[][]int, used *[]bool) {
	if len(inProcess) == len(nums) {
		temp := make([]int, len(inProcess))
		copy(temp, inProcess)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			inProcess = append(inProcess, nums[i])
			dfsProcess(nums, index+1, inProcess, res, used)
			inProcess = inProcess[:len(inProcess)-1]
			(*used)[i] = false
		}
	}
	return
}
