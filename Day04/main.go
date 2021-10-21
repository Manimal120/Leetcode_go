package main

import "fmt"

func main() {
	arr := [4]int{0, 0}
	a := arr[0:2]
	a1 := append(a, 1)
	a2 := append(a1, 2)
	a3 := append(a1, 3)
	fmt.Println(a1, a2, a3)
}
