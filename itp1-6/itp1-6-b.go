package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([][]int, 4)
	a[0] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	a[1] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	a[2] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	a[3] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	suite2num := map[string]int{
		"S": 0,
		"H": 1,
		"C": 2,
		"D": 3,
	}

	num2suite := map[int]string{
		0: "S",
		1: "H",
		2: "C",
		3: "D",
	}

	for i := 0; i < n; i++ {
		var suite string
		fmt.Scan(&suite)

		var rank int
		fmt.Scan(&rank)

		a[suite2num[suite]][rank-1] = 1
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			if a[i][j] == 0 {
				fmt.Printf("%s %d\n", num2suite[i], j+1)
			}
		}
	}
}
