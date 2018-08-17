package main

import (
	"fmt"
)

func checkNum(i int) {
	n := i
	if n%3 == 0 {
		fmt.Printf(" %d", i)
		endCheckNum(i)
		return
	}
	includeThree(i, n)
}

func includeThree(i, n int) {
	if n%10 == 3 {
		fmt.Printf(" %d", i)
		endCheckNum(i)
		return
	}
	n /= 10
	if n != 0 {
		includeThree(i, n)
		return
	}
	endCheckNum(i)
}

func endCheckNum(i int) {
	i++
	if i <= num {
		checkNum(i)
		return
	}
	fmt.Println()
}

var num int

func main() {
	fmt.Scan(&num)

	checkNum(1)
}
