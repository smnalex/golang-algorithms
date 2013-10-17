package main

import "fmt"

func swap(array []int, i, j int) {
	
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

func selectionSort(array []int) {
	
	for i := 0; i < len(array) - 1; i++ {
		min := i
		for j := i + 1; j < len(array) - 1; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		swap(array, i, min)
	}
}

func main() {

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", array)
	selectionSort(array)
	fmt.Println("Sorted array: ", array)
}