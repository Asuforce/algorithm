package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	e := []string{}
	if s.Scan() {
		e = strings.Split(s.Text(), " ")
	}

	fmt.Println(gcd(conv(e[0]), conv(e[1])))
}

func conv(e string) int {
	r, _ := strconv.Atoi(e)
	return r
}

func gcd(x, y int) int {
	if x < y {
		tmp := y
		y = x
		x = tmp
	}
	for y > 0 {
		r := x % y
		x = y
		y = r
	}
	return x
}
