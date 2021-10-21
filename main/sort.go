package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := make([]int, 0, 128)
	for i := 0; i < 5; i++ {
		s = append(s, rand.Intn(1000))
	}
	fmt.Printf("%v\n", s)
	//fmt.Println(sort.IntsAreSorted(s))
	quicksort(s)
	fmt.Printf("%v\n", s)
	//fmt.Println(sort.IntsAreSorted(s))
}

func quicksort(nums []int) {
	if len(nums) < 2 {
		return
	}
	base := nums[0]
	l, r := 0, len(nums)-1
	for l != r {
		for ; l != r && nums[r] > base; r-- {
		}
		nums[l] = nums[r]
		for ; l != r && nums[l] <= base; l++ {
		}
		nums[r] = nums[l]
	}
	nums[l] = base

	quicksort(nums[:l])
	quicksort(nums[l+1:])
}
