package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	suit  string
	value int
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	e := make([]card, n)

	for x := 0; x < n; x++ {
		tmp := ""
		fmt.Fscan(r, &tmp)
		e[x].suit = string([]rune(tmp)[0])
		e[x].value, _ = strconv.Atoi(string([]rune(tmp)[1]))
	}

	bo := make([]card, len(e))
	_ = copy(bo, e)
	so := make([]card, len(e))
	_ = copy(so, e)

	b := bubbleSort(bo)
	print(b, b)
	print(b, selectionSort(so))
}

func bubbleSort(e []card) []card {
	n := len(e)
	flag := true
	for flag {
		flag = false
		for j := n - 1; j > 0; j-- {
			if e[j].value < e[j-1].value {
				tmp := e[j]
				e[j] = e[j-1]
				e[j-1] = tmp
				flag = true
			}
		}
	}
	return e
}

func selectionSort(e []card) []card {
	n := len(e)
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if e[j].value < e[minj].value {
				minj = j
			}
		}
		tmp := e[i]
		e[i] = e[minj]
		e[minj] = tmp
	}
	return e
}

// Bubble sort always stable
func isStable(bubble, out []card) string {
	for i := 0; i < len(bubble); i++ {
		if bubble[i].suit != out[i].suit {
			return "Not stable"
		}
	}
	return "Stable"
}

func print(in, out []card) {
	n := len(in)
	for x := 0; x < n; x++ {
		if x == n-1 {
			fmt.Printf("%v%d\n", out[x].suit, out[x].value)
		} else {
			fmt.Printf("%v%d ", out[x].suit, out[x].value)
		}
	}
	fmt.Printf("%v\n", isStable(in, out))
}
