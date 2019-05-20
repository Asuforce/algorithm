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
	tail *node
}

func (ll *linkedList) insert(x int) {
	new := &node{x, nil, ll.head}
	ll.head.next = new
	ll.head = new
	if ll.tail == nil {
		ll.tail = ll.head
	}
}

func (ll *linkedList) delete(x int) {
	n := ll.searchList(x)
	if n.next == nil {
		ll.tail = n.prev
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n = nil
}

func (ll *linkedList) deleteFirst() {
	ll.delete(ll.head.key)
}

func (ll *linkedList) deleteLast() {
	ll.delete(ll.tail.key)
}

func (ll *linkedList) searchList(x int) *node {
	n := ll.head
	for n.key != x {
		n = n.prev
	}
	return n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	n := 0
	if s.Scan() {
		n, _ = strconv.Atoi(s.Text())
	}

	ll := &linkedList{}
	for i := 0; i < n; i++ {
		if !s.Scan() {
			break
		}

		com := strings.Fields(s.Text())
		switch com[0] {
		case "insert":
			ll.insert(stoi(com[1]))
		case "delete":
			ll.delete(stoi(com[1]))
		case "deleteFirst":
			ll.deleteFirst()
		case "deleteLast":
			ll.deleteLast()
		}
	}

	node := ll.head
	e := []int{}
	for node != nil {
		e = append(e, node.key)
		node = node.next
	}
	fmt.Println(e)
	// t := trimBracket(fmt.Sprint(e))
	// fmt.Println(strings.Replace(t, " ", "\n", -1))
}

func stoi(x string) int {
	i, _ := strconv.Atoi(x)
	return i
}

func trimBracket(s string) string {
	return strings.Trim(s, "[]")
}
