package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	a := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		a[i] = make([]int, c+1)
	}

	for i := 0; i < r; i++ {
		a[i] = make([]int, c+1)
		for j := 0; j < c; j++ {
			fmt.Scan(&a[i][j])
			a[i][c] += a[i][j]
			a[r][j] += a[i][j]
			a[r][c] += a[i][j]
		}
	}

	for _, ar := range a {
		for j, ac := range ar {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(ac)
		}
		fmt.Println()
	}
}
