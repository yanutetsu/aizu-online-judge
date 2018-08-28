package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	s := sc.Text()

	ss := make([]rune, len(s))

	for i, v := range s {
		ss[i] = v
	}

	sc.Scan()
	p := sc.Text()

	pp := make([]rune, len(p))

	for i, v := range p {
		pp[i] = v
	}

	same := false
	for si, sv := range ss {
		if sv == pp[0] {
			same = true
			for pi, pv := range pp {
				sl := si + pi
				if len(ss)-1 < sl {
					sl = sl - len(ss)
				}
				if ss[sl] != pv {
					same = false
					break
				}
			}
			if same {
				break
			} else {
				continue
			}
		}
	}

	if same {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
