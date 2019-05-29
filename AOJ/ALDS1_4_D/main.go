package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	s       = bufio.NewScanner(os.Stdin)
	n, k, w = 0, 0, []int{}
)

func check(p int) int {
	i := 0
	for j := 0; j < k; j++ {
		s := 0
		for s+w[i] <= p {
			s += w[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}

func binarySearch() int {
	left, right := 0, 100000*100000
	for right-left > 1 {
		mid := (left + right) / 2
		v := check(mid)
		if v >= n {
			right = mid
			continue
		}
		left = mid
	}
	return right
}

func main() {
	s.Split(bufio.ScanWords)

	n = nextInt(s)
	k = nextInt(s)
	w = make([]int, n)

	for i := range w {
		w[i] = nextInt(s)
	}
	fmt.Println(binarySearch())
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	r, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return r
}
