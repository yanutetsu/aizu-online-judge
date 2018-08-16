package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for i := 1; ; i++ {
		sc.Scan()
		s := sc.Text()

		ns := strings.Split(s, " ")

		a, _ := strconv.Atoi(ns[0])
		b, _ := strconv.Atoi(ns[1])

		if a == 0 && b == 0 {
			return
		}

		if a < b {
			fmt.Printf("%d %d\n", a, b)
		} else {
			fmt.Printf("%d %d\n", b, a)
		}
	}
}
