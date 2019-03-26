package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	_, _ = fmt.Scan(&x1, &y1, &x2, &y2)
	a := x2 - x1
	b := y2 - y1
	c := a*a + b*b
	fmt.Println(math.Sqrt(c))
}
