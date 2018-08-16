package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()

	s := sc.Text()
	r, _ := strconv.ParseFloat(s, 64)

	fmt.Printf("%f %f\n", r*r*math.Pi, r*2*math.Pi)
}
