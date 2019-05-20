package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	key  int
	next *node
	prev *node
}

type linkedList struct {
	head *node
}

func newLinkedList() *linkedList {
	n := &node{}
	n.prev, n.next = n, n
	return &linkedList{n}
}

func (ll *linkedList) insert(x int) {
	next := ll.head.next
	n := &node{}
	n.key = x
	n.next = next
	n.prev = ll.head
	ll.head.next = n
	next.prev = n
}

func (ll *linkedList) delete(n *node) {
	next := n.next
	prev := n.prev
	prev.next = next
	next.prev = prev
}

func (ll *linkedList) deleteFromIndex(x int) {
	for n := ll.head.next; n != ll.head; n = n.next {
		if n.key == x {
			ll.delete(n)
			return
		}
	}
}

func (ll *linkedList) deleteFirst() {
	ll.delete(ll.head.next)
}

func (ll *linkedList) deleteLast() {
	ll.delete(ll.head.prev)
}

func main() {
	n := 0
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	ll := newLinkedList()
	for i := 0; i < n; i++ {
		s.Scan()
		cmd := s.Text()
		switch cmd {
		case "deleteFirst":
			ll.deleteFirst()
		case "deleteLast":
			ll.deleteLast()
		default:
			s.Scan()
			x, _ := strconv.Atoi(s.Text())
			if cmd == "insert" {
				ll.insert(x)
				break
			}
			ll.deleteFromIndex(x)
		}
	}

	e := []int{}
	for n := ll.head.next; n != ll.head; n = n.next {
		e = append(e, n.key)
	}
	fmt.Println(trimBracket(fmt.Sprint(e)))
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}
