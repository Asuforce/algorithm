package main

import (
	"fmt"
	"math"
)

const th = math.Pi * 60.0 / 180.0

type point struct {
	x, y float64
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	a := point{0, 0}
	b := point{100, 0}

	fmt.Printf("%.8f %.8f\n", a.x, a.y)
	koch(n, a, b)
	fmt.Printf("%.8f %.8f\n", b.x, b.y)
}

func koch(d int, p1, p2 point) {
	if d == 0 {
		return
	}

	var (
		s point
		t point
		u point
	)

	s.x = (2.0*p1.x + 1.0*p2.x) / 3.0
	s.y = (2.0*p1.y + 1.0*p2.y) / 3.0
	t.x = (1.0*p1.x + 2.0*p2.x) / 3.0
	t.y = (1.0*p1.y + 2.0*p2.y) / 3.0
	u.x = (t.x-s.x)*math.Cos(th) - (t.y-s.y)*math.Sin(th) + s.x
	u.y = (t.x-s.x)*math.Sin(th) + (t.y-s.y)*math.Cos(th) + s.y

	koch(d-1, p1, s)
	fmt.Printf("%.8f %.8f\n", s.x, s.y)
	koch(d-1, s, u)
	fmt.Printf("%.8f %.8f\n", u.x, u.y)
	koch(d-1, u, t)
	fmt.Printf("%.8f %.8f\n", t.x, t.y)
	koch(d-1, t, p2)
}
