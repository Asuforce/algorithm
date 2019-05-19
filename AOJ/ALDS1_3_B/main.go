package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	p    []process
	head int
	tail int
	size int
}

type process struct {
	name string
	time int
}

func newQueue(n int) *queue {
	p := make([]process, n)
	return &queue{p, 0, 0, n - 1}
}

func (q *queue) enqueue(p process) {
	q.p[q.tail] = p
	if q.tail == q.size {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *queue) dequeue() process {
	p := q.p[q.head]
	if q.head == q.size {
		q.head = 0
	} else {
		q.head++
	}
	return p
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	e := []string{}
	if s.Scan() {
		e = strings.Fields(s.Text())
	}
	n, _ := strconv.Atoi(e[0])
	q, _ := strconv.Atoi(e[1])

	Q := newQueue(n)
	for i := 0; i < n; i++ {
		if !s.Scan() {
			break
		}
		tmp := strings.Fields(s.Text())
		t, _ := strconv.Atoi(tmp[1])
		p := process{tmp[0], t}
		Q.enqueue(p)
	}

	cnt := 0
	for n > 0 {
		p := Q.dequeue()
		t := p.time - q
		if t > 0 {
			p.time = t
			Q.enqueue(p)
			cnt = cnt + q
		} else {
			cnt = cnt + p.time
			n--
			fmt.Printf("%v %d\n", p.name, cnt)
		}
	}
}
