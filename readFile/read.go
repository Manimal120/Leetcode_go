package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var a int
var b int

//写这玩意儿是因为用Go刷牛客的人特别少，很多题解都没看到go语言的
//然后面对机试的ACM核心模式，足以让常年用Go刷LC的人很不适应
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		a, _ = strconv.Atoi(strings.Split(input.Text(), " ")[0])
		b, _ = strconv.Atoi(strings.Split(input.Text(), " ")[1])
	}
	fmt.Println(a + b)

}
