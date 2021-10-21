package main

import "fmt"

func main() {
	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a, 9))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil

}
