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

	fmt.Printf("%v\n", strings.Trim(fmt.Sprint(e), "[]"))

	for i := 1; i < n; i++ {
		v := e[i]
		j := i - 1
		for j >= 0 && e[j] > v {
			e[j+1] = e[j]
			j--
		}
		e[j+1] = v
		fmt.Printf("%v\n", strings.Trim(fmt.Sprint(e), "[]"))
	}
}
