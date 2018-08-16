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

	n, _ := nextInt()

	max := 0
	min := 0
	sum := 0
	for i := 0; i < n; i++ {
		a, _ := nextInt()
		if a <= min || i == 0 {
			min = a
		}
		if max <= a || i == 0 {
			max = a
		}
		sum += a
	}

	fmt.Printf("%d %d %d\n", min, max, sum)
}
