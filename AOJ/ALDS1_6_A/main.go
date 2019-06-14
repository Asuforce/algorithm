package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VMAX = 10e4

var s = bufio.NewScanner(os.Stdin)

func main() {
	s.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	fmt.Println(trimBracket(fmt.Sprint(countingSort(a))))
}

func nextInt() int {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n
}

func countingSort(A []int) []int {
	C := make([]int, VMAX+1)
	for _, v := range A {
		C[v]++
	}
	for j := 1; j <= VMAX; j++ {
		C[j] = C[j] + C[j-1]
	}
	B := make([]int, len(A))
	for _, v := range A {
		B[C[v]-1] = v
		C[v]--
	}
	return B
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}
