package main

import (
	"fmt"
)

func getLowerCase(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}

	return string(b)
}

func main() {
	var w string
	fmt.Scan(&w)
	w = getLowerCase(w)

	count := 0
	for {
		var t string
		fmt.Scan(&t)

		if t == "END_OF_TEXT" {
			break
		}

		t = getLowerCase(t)
		if w == t {
			count++
		}
	}
	fmt.Println(count)
}
