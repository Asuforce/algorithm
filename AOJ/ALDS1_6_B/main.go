package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	s.Split(bufio.ScanWords)
	n := nextInt(s)
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt(s)
	}

	q := partition(a, 0, n-1)
	for i := range a {
		if i == q {
			fmt.Printf("[%d] ", a[i])
			continue
		} else if i == len(a)-1 {
			fmt.Printf("%d\n", a[i])
			break
		}
		fmt.Printf("%d ", a[i])
	}
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	r, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return r
}

func partition(A []int, p, r int) int {
	x, y, i := A[r], A[:r], p-1
	for j := p; j < len(y); j++ {
		if t := A[j]; t <= x {
			i++
			A[j] = A[i]
			A[i] = t
		}
	}
	A[r] = A[i+1]
	A[i+1] = x

	return i + 1
}
