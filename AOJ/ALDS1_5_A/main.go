package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	s    = bufio.NewScanner(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
	a, n = []int{}, 0
)

func main() {
	n = nextInt(s)
	a = atoi(s)
	_ = nextInt(s)
	m := atoi(s)

	for _, v := range m {
		i := 0
		if solve(i, v) {
			fmt.Fprintln(w, "yes")
			continue
		}
		fmt.Fprintln(w, "no")
	}
	w.Flush()
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	r, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return r
}

func atoi(s *bufio.Scanner) []int {
	s.Scan()
	a := strings.Fields(s.Text())
	i := []int{}

	for _, v := range a {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		i = append(i, j)
	}
	return i
}

func solve(i, m int) bool {
	if m == 0 {
		return true
	} else if i >= n {
		return false
	}
	return solve(i+1, m) || solve(i+1, m-a[i])
}
