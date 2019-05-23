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
	cnt := 0
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	S := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		S = append(S, v)
	}

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	for i := 0; i < q; i++ {
		sc.Scan()
		key, _ := strconv.Atoi(sc.Text())
		S = append(S, key)
		if search(S, key) {
			cnt++
		}
		S = S[:len(S)-1]
	}

	fmt.Println(cnt)
}

func search(s []int, key int) bool {
	for i := 0; i != len(s)-1; i++ {
		if s[i] == key {
			return true
		}
	}
	return false
}
