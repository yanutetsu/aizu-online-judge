package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var c string
		fmt.Scan(&c)

		switch c {
		case "print":
			var a, b int
			fmt.Scan(&a, &b)
			fmt.Printf("%s\n", str[a:b+1])
		case "reverse":
			var a, b int
			fmt.Scan(&a, &b)
			buf := make([]rune, len(str))
			j := b
			for i, v := range str {
				if a <= i && i <= b {
					buf[j] = v
					if j != 0 {
						j--
					}
				} else {
					buf[i] = v
				}
			}
			str = string(buf)
		case "replace":
			var a, b int
			var p string
			fmt.Scan(&a, &b, &p)
			str = str[0:a] + p + str[b+1:len(str)]
		}
	}
}
