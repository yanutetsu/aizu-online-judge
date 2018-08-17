package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextChar(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextInt(sc *bufio.Scanner) int {
	i, _ := strconv.Atoi(nextChar(sc))
	return i
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for {
		h := nextInt(sc)
		w := nextInt(sc)

		if h == 0 && w == 0 {
			return
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Print("#")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
