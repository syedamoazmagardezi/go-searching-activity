package main

import "fmt"

func main() {
	scores := []int{76, 44, 30, 50, 100}

	fmt.Println("All scores:", scores)
	fmt.Println("First:", scores[0])
	fmt.Println("Third:", scores[2])
	fmt.Println("\nAll values using loop :")
	for i := 0; i < len(scores); i++ {
		fmt.Println("Index", i, "=", scores[i])
	}

	fmt.Println("\nValues > 75 :")
	for i := 0; i < len(scores); i++ {
		if scores[i] > 75 {
			fmt.Println("Index", i, "=", scores[i])
		}
	}
	fmt.Println("\nSmallest is Index 2 = 30")
}
