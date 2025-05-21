package main

import (
	"ordenacaobusca/bubbleSort"
	"ordenacaobusca/insertionSort"
	"fmt"
)
func main() {
	array :=[]int{8,5,1,6,9,5,7}

	var choice int
	fmt.Println("Choose the sorting algorithm you want to use:")
	fmt.Println("1. Bubble Sort")
	fmt.Println("2. Insertion Sort")
	_, err := fmt.Scanln(&choice)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
	}

	switch choice {
		case 1:
			bubbleSort.BubbleSort(array)

		case 2:
			insertionSort.InsertionSort(array)

		default:
			fmt.Println("Invalid choice. Please select 1 or 2.")
	}
}