package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	read1 := bufio.NewReader(os.Stdin)
	line1, _, err := read1.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}
	read2 := bufio.NewReader(os.Stdin)
	line2, _, err := read2.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}

	var n, _ = strconv.Atoi(string(line1))
	numsString := strings.Split(string(line2), " ")

	if n == 0 || n == 1 {
		return
	}

	var nums []int

	for _, s := range numsString {
		s1, _ := strconv.Atoi(s)
		nums = append(nums, s1)
	}

	mapInt := make(map[int]int)

	for i := 0; i < n; i++ {
		if _, ok := mapInt[nums[i]]; ok {
			mapInt[nums[i]]++
			continue
		}
		mapInt[nums[i]] = 1
	}

	maxV, minV := 1, 1

	for _, v := range mapInt {
		maxV = max(maxV, v)
		minV = min(minV, v)
	}

	var maxK, minK = nums[0], nums[0]

	for k, v := range mapInt {
		if v == maxV {
			maxK = max(maxK, k)
		}
		if v == minV {
			minK = min(minK, k)
		}
	}

	fmt.Println(maxK - minK)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
