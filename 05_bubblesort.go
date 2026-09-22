package main

import "fmt"

func main() {
	numbers := []int{9, 4, 7, 1, 5}
	fmt.Println("Before:", numbers)
	for pass := 0; pass < len(numbers)-1; pass++ {
		for i := 0; i < len(numbers)-1-pass; i++ {
			if numbers[i] < numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
		fmt.Println("Pass", pass+1, numbers)
	}
	fmt.Println("Sorted:", numbers)

}
