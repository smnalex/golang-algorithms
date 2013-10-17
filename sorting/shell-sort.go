package main

import (
	"fmt"
	"algoutils"
)

func shellSort(array []int) {
	h := 1
	for h < len(array) {
		h = 3 * h + 1
	}
	for h >= 1 {
		for i := h; i < len(array); i++ {
			for j := i; j >= h && array[j] < array[j - h]; j = j - h {
				algoutils.Swap(array, j, j - h)
			}
		}
		h = h/3
	}
}

func main() {
  
  array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
  fmt.Println("Unsorted array: ", array)
  shellSort(array)
  fmt.Println("Sorted array: ", array)
}