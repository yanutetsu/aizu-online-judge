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
	sc.Split(bufio.ScanWords)

	a, _ := nextInt()
	b, _ := nextInt()

	o := "=="
	if a < b {
		o = "<"
	} else if a > b {
		o = ">"
	}

	fmt.Printf("a %s b\n", o)
}
