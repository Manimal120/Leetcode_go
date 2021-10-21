package main

//func main() {
//	s := "abc"
//	fmt.Println(s[0])	// 97
//}

// 1. judge nil
// 2. create an array to record freq of each letter
// 3. initiate result left right
// 4. when left less than array board
// 	4.1 right+1 less than array board, and the letter's freq is 0
//  	do: freq+1, right++
//  4.2 else: left move to the next letter, the

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	var freq [256]int
	left, right, res := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
