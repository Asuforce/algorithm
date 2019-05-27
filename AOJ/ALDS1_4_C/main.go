package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// M is Magic number
const M = 123456

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type hash struct {
	T []string
}

func newHash() *hash {
	s := make([]string, M)
	return &hash{T: s}
}

func (ha *hash) insert(str string) {
	key := getInt(str)
	for i := 0; ; i++ {
		j := h(key, i)
		if ha.T[j] == str {
			return
		} else if len(ha.T[j]) == 0 {
			ha.T[j] = str
			break
		}
	}
}

func (ha *hash) find(str string) bool {
	key := getInt(str)
	for i := 0; i < len(ha.T); i++ {
		j := h(key, i)
		if ha.T[j] == str {
			return true
		}
	}
	return false
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	ha := newHash()
	for i := 0; sc.Scan() && i < n; i++ {
		cmd := strings.Fields(sc.Text())
		switch cmd[0] {
		case "find":
			if ha.find(cmd[1]) {
				fmt.Fprintln(wr, "yes")
			} else {
				fmt.Fprintln(wr, "no")
			}
		case "insert":
			ha.insert(cmd[1])
		}
	}
	wr.Flush()
}

func atoi(str byte) int {
	switch str {
	case 'A':
		return 1
	case 'C':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	default:
		return 0
	}
}

func getInt(str string) int {
	cnt, p := 0, 1
	for i := 0; i < len(str); i++ {
		cnt += p * atoi(str[i])
		p *= 5
	}
	return cnt
}

func h1(key int) int {
	return key % M
}

func h2(key int) int {
	return 1 + key%(M-1)
}

func h(key, i int) int {
	return (h1(key) + i*h2(key)) % M
}
