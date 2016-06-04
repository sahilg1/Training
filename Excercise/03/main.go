package main

import "fmt"

func main() {
	largest := vary(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(largest)
}

func vary(num ...int) int {
	var largest int
	for _, x := range num {
		if x > largest {
			largest = x
		}
	}
	return largest
}
