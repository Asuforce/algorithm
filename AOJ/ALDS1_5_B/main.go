package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// INFTY is over max num
const INFTY = 100000 * 100000

var (
	s   = bufio.NewScanner(os.Stdin)
	cnt = 0
)

func main() {
	s.Split(bufio.ScanWords)

	n := nextInt(s)
	S := make([]int, n)

	for i := range S {
		S[i] = nextInt(s)
	}
	mergeSort(S, 0, n)
	fmt.Println(trimBracket(fmt.Sprint(S)))
	fmt.Println(cnt)
}

func nextInt(s *bufio.Scanner) int {
	s.Scan()
	r, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return r
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}

func mergeSort(A []int, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}

	L[n1], R[n2] = INFTY, INFTY
	i, j := 0, 0
	for k := left; k < right; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
			continue
		}
		A[k] = R[j]
		j++
	}
}
