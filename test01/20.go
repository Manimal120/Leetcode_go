package main

//func main() {
//	input := "{[]}"
//	fmt.Println(isValid(input))
//}

// 1. judge nil
// 2. create a stack
// 3. range string, ([{ put in
// 4. )]}, stack exist values, and () [] {} can be eliminate, stack - 1
// 5. else return false
// 6. len(stack) = 0 ?

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
