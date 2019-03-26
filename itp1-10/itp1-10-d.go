package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var y, p, p1, p2, p3, p4 float64
	fmt.Scan(&n)

	x := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&y)
		p = math.Abs(x[i] - y)
		p1 = p1 + p
		p2 = p2 + math.Pow(p, 2)
		p3 = p3 + math.Pow(p, 3)
		p4 = math.Max(p, p4)
	}
	fmt.Println(p1)
	fmt.Println(math.Sqrt(p2))
	fmt.Println(math.Cbrt(p3))
	fmt.Println(p4)
}
