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

	a, c := bubbleSort(e, n)
	fmt.Printf("%v\n%d\n", strings.Trim(fmt.Sprint(a), "[]"), c)
}

func bubbleSort(e []int, n int) ([]int, int) {
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
	return e, c
}
