package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		sc.Scan()
		s := sc.Text()
		if s == "0" {
			return
		}

		sc.Scan()
		s = sc.Text()
		ns := strings.Split(s, " ")

		m := float64(0)
		var fs []float64
		for _, v := range ns {
			d, _ := strconv.ParseFloat(v, 64)
			m += d
			fs = append(fs, d)
		}
		m = m / float64(len(fs))

		b := float64(0)
		for _, v := range fs {
			b += math.Pow(m-v, 2)
		}
		b = b / float64(len(ns))

		fmt.Println(math.Sqrt(b))
	}
}
