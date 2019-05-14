package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	var e = make([]int, n)

	for x := 0; x < n; x++ {
		fmt.Fscan(r, &e[x])
	}

	flag := true
	c := 0
	for flag {
		flag = false
		for j := n - 1; j > 0; j-- {
			if e[j] < e[j-1] {
				tmp := e[j]
				e[j] = e[j-1]
				e[j-1] = tmp
				flag = true
				c++
			}
		}
	}
	fmt.Printf("%v\n%d\n", strings.Trim(fmt.Sprint(e), "[]"), c)
}
