package main

import "fmt"

const str = "Hey I am new here"

func main() {
	const p = 42
	q := "Hey"
	fmt.Printf("%T - %T \n %T \n", q, p, str)

}
