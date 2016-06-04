package main

import "fmt"

func main() {

	printing := func(num int) (float64, bool) {

		if num%2 == 0 {
			return float64(num) / 2, true
		}
		return float64(num) / 2, false
	}
	h, even := printing(99)
	fmt.Println(h, "-", even)
}
