package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextChar(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextInt(sc *bufio.Scanner) int {
	i, _ := strconv.Atoi(nextChar(sc))
	return i
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for {
		a := nextInt(sc)
		op := nextChar(sc)
		b := nextInt(sc)

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
