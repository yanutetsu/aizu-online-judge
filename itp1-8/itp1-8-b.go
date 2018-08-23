package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		sc.Scan()
		s := sc.Text()

		if si, _ := strconv.Atoi(s); si == 0 {
			return
		}

		t := 0
		for _, r := range s {
			ri := int(r - '0')
			t += ri
		}
		fmt.Println(t)
	}
}
