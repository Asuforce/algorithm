package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	max := -9999999999
	min := 9999999999
	for x := 0; x < n; x++ {
		var e int
		fmt.Fscan(r, &e)
		fmt.Printf("max: %d\n", max)
		if max < e-min {
			max = e - min
		}
		if min > e {
			min = e
		}
	}
	fmt.Println(max)
}
