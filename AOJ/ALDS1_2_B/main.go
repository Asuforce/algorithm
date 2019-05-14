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

	a, c := selectionSort(e, n)
	fmt.Printf("%v\n%d\n", strings.Trim(fmt.Sprint(a), "[]"), c)
}

func selectionSort(e []int, n int) ([]int, int) {
	cnt := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if e[j] < e[minj] {
				minj = j
			}
		}
		if i != minj {
			tmp := e[i]
			e[i] = e[minj]
			e[minj] = tmp
			cnt++
		}
	}
	return e, cnt
}
