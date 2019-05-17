package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numbers struct {
	e   []int
	g   []int
	cnt int
}

func (n *numbers) lenE() int {
	return len(n.e)
}

func (n *numbers) lenG() int {
	return len(n.g)
}

func (n *numbers) appendG(i int) {
	n.g = append(n.g, 0)
	copy(n.g[1:], n.g[0:])
	n.g[0] = i
}

func (n *numbers) shellSort() {
	for i := 1; i <= n.lenE(); i = 3*i + 1 {
		n.appendG(i)
	}

	for j := 0; j < n.lenG(); j++ {
		n.insertionSort(n.g[j])
	}
}

func (n *numbers) insertionSort(g int) {
	for i := g; i < n.lenE(); i++ {
		v := n.e[i]
		j := i - g
		for j >= 0 && n.e[j] > v {
			n.e[j+g] = n.e[j]
			j = j - g
			n.cnt++
		}
		n.e[j+g] = v
	}
}

var s = bufio.NewScanner(os.Stdin)

func main() {
	n := 0
	if s.Scan() {
		n, _ = strconv.Atoi(s.Text())
	}

	e := make([]int, n)
	num := &numbers{nil, nil, 0}
	for i := range e {
		if !s.Scan() {
			break
		}
		num.e[i], _ = strconv.Atoi(s.Text())
	}

	num.shellSort()

	fmt.Println(num.lenG())
	fmt.Println(trimBracket(fmt.Sprint(num.g)))
	fmt.Println(num.cnt)
	t := trimBracket(fmt.Sprint(num.e))
	fmt.Println(strings.Replace(t, " ", "\n", -1))
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}
