package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

const (
	MAX_PACKAGES = 100000
	MAX_TRACKS   = 100000
)

type allocation struct {
	n int
	k int
	w []int
}

func (a *allocation) check(p int) int {
	i := 0
	for j := 0; j < a.k; j++ {
		s := 0
		for s+a.w[i] <= p {
			s += a.w[i]
			i++
			if i == a.n {
				return a.n
			}
		}
	}
	return i
}

func (a *allocation) binarySearch() int {
	left, mid, right := 0, 0, MAX_PACKAGES*MAX_TRACKS
	for right-left > 1 {
		mid = (left + right) / 2
		v := a.check(mid)
		if v >= a.n {
			right = mid
			continue
		}
		left = mid
	}
	return right
}

func main() {
	s.Split(bufio.ScanWords)

	a := &allocation{}
	a.n = nextInt(s)
	a.k = nextInt(s)
	a.w = make([]int, a.n)

	for i := range a.w {
		a.w[i] = nextInt(s)
	}
	fmt.Println(a.binarySearch())
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return n
}
