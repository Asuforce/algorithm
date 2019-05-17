package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	t   *node
	len int
}

type node struct {
	v    int
	prev *node
}

func (s *stack) Len() int {
	return s.len
}

func (s *stack) Peek() int {
	if s.len == 0 {
		return -1
	}
	return s.t.v
}

func (s *stack) Pop() int {
	if s.len == 0 {
		return -1
	}
	n := s.t
	s.t = n.prev
	s.len--
	return n.v
}

func (s *stack) Push(v int) {
	s.t = &node{v: v, prev: s.t}
	s.len++
}

func (s *stack) calc(op string) {
	l, r := s.Pop(), s.Pop()
	switch op {
	case "+":
		s.Push(l + r)
	case "-":
		s.Push(r - l)
	case "*":
		s.Push(l * r)
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	e := []string{}
	if s.Scan() {
		e = strings.Fields(s.Text())
	}
	f := &stack{nil, 0}
	for _, ele := range e {
		if num, err := strconv.Atoi(ele); err == nil {
			f.Push(num)
		} else {
			f.calc(ele)
		}
	}
	fmt.Println(f.Peek())
}
