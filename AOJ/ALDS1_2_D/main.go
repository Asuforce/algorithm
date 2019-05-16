package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := 0
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		n, _ = strconv.Atoi(s.Text())
	}

	e := make([]int, n)
	for i := range e {
		if !s.Scan() {
			break
		}
		e[i], _ = strconv.Atoi(s.Text())
	}

	g, c := shellSort(e)

	fmt.Println(len(g))
	fmt.Println(trimBracket(fmt.Sprint(g)))
	fmt.Println(c)
	t := trimBracket(fmt.Sprint(e))
	fmt.Println(strings.Replace(t, " ", "\n", -1))
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}

func insertionSort(e []int, g int) int {
	cnt := 0
	for i := g; i < len(e); i++ {
		v := e[i]
		j := i - g
		for j >= 0 && e[j] > v {
			e[j+g] = e[j]
			j = j - g
			cnt++
		}
		e[j+g] = v
	}
	return cnt
}

func shellSort(e []int) ([]int, int) {
	cnt := 0
	g := []int{}
	for i := 1; i <= len(e); i = 3*i + 1 {
		g = append(g, 0)
		copy(g[1:], g[0:])
		g[0] = i
	}

	for j := 0; j < len(g); j++ {
		cnt += insertionSort(e, g[j])
	}

	return g, cnt
}
