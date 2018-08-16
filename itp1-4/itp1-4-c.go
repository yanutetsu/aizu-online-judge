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

	n := nextInt()

	for i := 0; i < n; i++ {
		a, _ := nextInt()
		sc.Scan()
		op := sc.Text()
		b, _ := nextInt()

		switch op {
		case "?":
			return
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		}
	}
}
