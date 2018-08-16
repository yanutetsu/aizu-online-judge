package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() (int, error) {
	sc.Scan()
	s := sc.Text()
	return strconv.Atoi(s)
}

func main() {
	for i := 1; ; i++ {
		a, _ := nextInt()
		if a == 0 {
			return
		}

		fmt.Printf("Case %d: %d\n", i, a)
	}
}
