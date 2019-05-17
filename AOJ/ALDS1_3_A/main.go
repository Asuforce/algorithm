package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []int

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Push(v int) stack {
	return append(s, v)
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	e := []string{}
	if s.Scan() {
		e = strings.Split(s.Text(), " ")
	}

	l, r := 0, 0
	f := make(stack, 0)

	for _, ele := range e {
		if num, err := strconv.Atoi(ele); err == nil {
			f = f.Push(num)
		} else {
			f, l = f.Pop()
			f, r = f.Pop()
			switch ele {
			case "+":
				f = f.Push(l + r)
			case "-":
				f = f.Push(r - l)
			case "*":
				f = f.Push(l * r)
			}
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(f), "[]"))
}
