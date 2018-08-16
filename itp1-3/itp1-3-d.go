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
	sc.Scan()
	s := sc.Text()

	ns := strings.Split(s, " ")

	a, _ := strconv.Atoi(ns[0])
	b, _ := strconv.Atoi(ns[1])
	c, _ := strconv.Atoi(ns[2])

	count := 0
	for i := a; i < b+1; i++ {
		if c%i == 0 {
			count++
		}
	}

	fmt.Println(count)
}
