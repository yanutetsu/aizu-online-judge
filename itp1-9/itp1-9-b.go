package main

import (
	"fmt"
)

func main() {
	result := []string{}

	for {
		var s string
		fmt.Scan(&s)

		if s == "-" {
			break
		}

		var m int
		fmt.Scan(&m)

		for i := 0; i < m; i++ {
			var h int
			fmt.Scan(&h)

			head := s[:h]
			tail := s[h:]

			s = tail + head
		}

		result = append(result, s)
	}

	for _, v := range result {
		fmt.Println(v)
	}
}
