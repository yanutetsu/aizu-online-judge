package main

import "fmt"

func main() {
	var m, f, r int
	var a string
	for {
		fmt.Scan(&m, &f, &r)

		if m == -1 && f == -1 && r == -1 {
			return
		}

		if m == -1 || f == -1 {
			a = "F"
		} else {
			sum := m + f
			if 80 <= sum {
				a = "A"
			} else if 65 <= sum && sum < 80 {
				a = "B"
			} else if 50 <= sum && sum < 65 {
				a = "C"
			} else if 30 <= sum && sum < 50 {
				if 50 <= r {
					a = "C"
				} else {
					a = "D"
				}
			} else if sum < 30 {
				a = "F"
			}
		}

		fmt.Println(a)
	}
}
