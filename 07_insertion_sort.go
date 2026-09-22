package main

import "fmt"

func main() {
	numbers := []int{7, 3, 8, 2, 5}
	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = key
		fmt.Println("Step", i, numbers)
	}
	fmt.Println("Sorted:", numbers)
}
