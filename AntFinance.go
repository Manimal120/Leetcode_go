package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	a := make([]int, n)
	br := bufio.NewScanner(os.Stdin)
	br.Scan()
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(strings.Split(br.Text(), " ")[i])
	}
}
