package main

//func main() {
//	a := []int{2, 7, 11, 15}
//	fmt.Println(twoSum(a, 9))
//}
// 1. make a map
// 2. find another by target
// 3. if m[another] exists, return directly, else put in map
// 4. record map
// 5. final return nil

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
