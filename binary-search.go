package main

import "fmt"

func sortedBinarySearch(sortedArray []int, toFind int) int {

  var low, high int
  low = sortedArray[0]
  high = sortedArray[len(sortedArray) - 1]

  if toFind > high || toFind < low {
    return -1
  }

  for low <= high {
    mid := low + (high - low)/2
    switch {
      case toFind < sortedArray[mid]:
        high = mid - 1
      
      case toFind > sortedArray[mid]:
        low = mid + 1
      
      case toFind == sortedArray[mid]:
        return mid
    }
  }
  return -1
}

func main() {
  
  // Sorted binary search test
  fmt.Println("Sorted binary search:")
  sortedArray := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
  index := sortedBinarySearch(sortedArray, 10)
  if index == -1 {
    fmt.Println("Number not found")
  } else {
    fmt.Println("Index: ", index)
    fmt.Println("sortedArray[", index, "] = ", sortedArray[index])
  }

  // Unsorted binary search test
  fmt.Println("Unsorted binary search:")
  unsortedArray := []int{0, 4, 2, 8, 6, 12, 10, 16, 14, 18}
  index = unsortedBinarySearch(unsortedArray, 10)
  if index == -1 {
    fmt.Println("Number not found")
  } else {
    fmt.Println("Index: ", index)
    fmt.Println("unsortedArray[", index, "] = ", unsortedArray[index])
  }
}