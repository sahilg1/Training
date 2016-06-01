package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %d \t %x \t %q \n", i, i, i, i)
	}
}
