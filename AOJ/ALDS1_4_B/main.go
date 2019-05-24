package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	cnt, n := 0, nextInt(sc)

	S := make([]int, n)
	for i := range S {
		S[i] = nextInt(sc)
	}
	q := nextInt(sc)
	for i := 0; i < q; i++ {
		key := nextInt(sc)
		if binarySearch(S, key) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func binarySearch(x []int, key int) bool {
	left, mid, right := 0, 0, len(x)
	for left < right {
		mid = (left + right) / 2
		if x[mid] == key {
			return true
		} else if key > x[mid] {
			left = mid + 1
			continue
		}
		right = mid
	}
	return false
}
