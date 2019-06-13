package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX   = 10e5
	INFTY = 10e9 + 1
)

var s = bufio.NewScanner(os.Stdin)

type card struct {
	suit  string
	value int
}

func main() {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	a, b := make([]card, n), make([]card, n)

	for x := 0; x < n; x++ {
		s.Scan()
		tmp := strings.Fields(s.Text())
		a[x].suit = string(tmp[0])
		a[x].value, _ = strconv.Atoi(tmp[1])
	}

	_ = copy(b, a)
	mergeSort(a, 0, n)
	quickSort(b, 0, n-1)

	fmt.Println(isStable(a, b))
	for _, v := range b {
		fmt.Printf("%v %d\n", v.suit, v.value)
	}
}

func isStable(a, b []card) string {
	for y := 0; y < len(a); y++ {
		if a[y].suit != b[y].suit {
			return "Not stable"
		}
	}
	return "Stable"
}

func merge(A []card, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	L := make([]card, MAX/2+2)
	R := make([]card, MAX/2+2)

	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}

	L[n1].value, R[n2].value = INFTY, INFTY
	i, j := 0, 0
	for k := left; k < right; k++ {
		if L[i].value <= R[j].value {
			A[k] = L[i]
			i++
			continue
		}
		A[k] = R[j]
		j++
	}
}

func mergeSort(A []card, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func partition(A []card, p, r int) int {
	x, y, i := A[r], A[:r], p-1
	for j := p; j < len(y); j++ {
		if t := A[j]; t.value <= x.value {
			i++
			A[j] = A[i]
			A[i] = t
		}
	}
	A[r] = A[i+1]
	A[i+1] = x

	return i + 1
}

func quickSort(A []card, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}
