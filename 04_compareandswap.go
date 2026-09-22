package main

import "fmt"

func main() {
	p := 25
	q := 30
	fmt.Println("Before:", p, q)
	p, q = q, p

	fmt.Println("After :", p, q)
}
