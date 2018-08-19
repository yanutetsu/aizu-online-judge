package main

import "fmt"

func main() {
	a := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			a[i][j] = make([]int, 10)
		}
	}

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scan(&b, &f, &r, &v)

		a[b-1][f-1][r-1] += v
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", a[i][j][k])
			}
			fmt.Println()
		}
		if i < 3 {
			fmt.Println("####################")
		}
	}
}
