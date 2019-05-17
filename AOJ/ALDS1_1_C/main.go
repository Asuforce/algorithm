package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	n, cnt := 0, 0
	if sc.Scan() {
		n = nextInt()
	}
	for i := 0; i < n; i++ {
		if !sc.Scan() {
			break
		}
		if isPrime(nextInt()) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	} else if x < 2 || x%2 == 0 {
		return false
	}
	max := int(math.Sqrt(float64(x)))
	for i := 3; i < max+1; i = i + 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
