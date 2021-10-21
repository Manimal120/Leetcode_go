package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	in := []byte(input)
	res := make([]byte, 0)

	for i := 0; i < len(in); i++ {
		res = append(res, in[len(in)-1-i])
	}
	fmt.Print(string(res))
}
