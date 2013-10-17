package main

import "fmt"

func linearSearch(arrayzor []int, toFind int) int {
	for i, item := range arrayzor {
		if item == toFind {
			return i
		}
	}
	return -1
}

func main() {
	
	fmt.Println("Linear search:")
	arrayzor := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	index := linearSearch(arrayzor, 10)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("arrayzor[", index, "] = ", arrayzor[index])
	}
}
