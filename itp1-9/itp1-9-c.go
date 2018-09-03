package main

import (
	"fmt"
)

func main() {
	tp := 0
	hp := 0

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var t, h string
		fmt.Scan(&t, &h)

		l := len(t)
		if len(h) < len(t) {
			l = len(h)
		}

		r := 0
		for j := 0; j < l; j++ {
			if t[j] == h[j] {
				continue
			} else if h[j] < t[j] {
				r = 1
				break
			} else if t[j] < h[j] {
				r = 2
				break
			}
		}
		if r == 0 {
			if len(t) == len(h) {
				tp++
				hp++
			} else if len(h) < len(t) {
				tp += 3
			} else if len(t) < len(h) {
				hp += 3
			}
		} else if r == 1 {
			tp += 3
		} else if r == 2 {
			hp += 3
		}
	}
	fmt.Printf("%d %d\n", tp, hp)
}
