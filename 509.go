package main

func fib(N int) int {
	memo := make([]int, N+1)
	return helper(memo, N)
}

func helper(memo []int, n int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}

//
//func main() {
//	fmt.Println(fib(3))
//}

func fib3(N int) int {
	if N <= 1 {
		return N
	}
	if N == 2 {
		return 1
	}
	current, prev1, prev2 := 0, 1, 1
	for i := 3; i <= N; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return current
}
