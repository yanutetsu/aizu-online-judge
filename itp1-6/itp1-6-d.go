package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var m int
	fmt.Scan(&m)

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	c := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i] += a[i][j] * b[j]
		}
	}

	for _, v := range c {
		fmt.Println(v)
	}
}
