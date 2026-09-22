package main

import "fmt"

func main() {
	numbers := []int{7, 3, 8, 2, 5}
	for i := 0; i < len(numbers)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
		fmt.Println("Pass", i+1, numbers)
	}
	fmt.Println("Sorted:", numbers)
}
