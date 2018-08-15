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

	w, _ := nextInt()
	h, _ := nextInt()
	x, _ := nextInt()
	y, _ := nextInt()
	r, _ := nextInt()

	result := "No"
	if (0+r) <= x && x <= (w-r) && (0+r) <= y && y <= (h-r) {
		result = "Yes"
	}

	fmt.Printf("%s\n", result)
}
