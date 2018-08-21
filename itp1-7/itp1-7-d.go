package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	a := make([][]int, n)
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		c[i] = make([]int, l)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Scan(&b[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	for _, cr := range c {
		for j, cc := range cr {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(cc)
		}
		fmt.Println()
	}
}
