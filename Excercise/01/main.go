package main

import "fmt"

func main() {
	h, even := half(99)
	fmt.Println(h, "-", even)
}

func half(num int) (float64, bool) {

	if num%2 == 0 {
		return float64(num) / 2, true
	}
	return float64(num) / 2, false
}
