package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	m := make([]int, 26)

	for {
		sc.Scan()
		s := sc.Text()
		if len(s) == 0 {
			for k, v := range m {
				fmt.Printf("%c : %d\n", k+97, v)
			}
			return
		}
		for _, r := range s {
			if 'A' <= r && r <= 'Z' {
				m[r-65]++
			} else if 'a' <= r && r <= 'z' {
				m[r-97]++
			}
		}
	}
}
