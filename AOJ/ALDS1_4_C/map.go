package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)
var w = bytes.Buffer{}

func main() {
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	T := make(map[string]int)
	for i := 0; s.Scan() && i < n; i++ {
		cmd := s.Text()
		if cmd[0] == 'f' {
			s.Scan()
			if _, ok := T[s.Text()]; ok {
				w.WriteString("yes\n")
				continue
			}
			w.WriteString("no\n")
			continue
		}
		s.Scan()
		T[s.Text()] = 0
	}
	fmt.Println(strings.TrimRight(w.String(), "\n"))
}
