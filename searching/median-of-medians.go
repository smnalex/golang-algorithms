package main

import (
	"fmt"
	"sort"
)

//
// medianOfMedians takes a list to be searched,
// the number i-thm smallest element (k), and
// the size of the subarrays (r) - r should be >= 3
//
func medianOfMedians(elementList []int, k, r int) int {

	// our base case, len < 10
	if len(elementList) < 10 {
		sort.Ints(elementList)
		return elementList[k - 1]
	}

	// m is the number of subarrays we need
	m := (len(elementList) + r - 1)/r

	medians := make([]int, m)

	// fill our median array
	for i := 0; i < m; i++ {
		
		v := (i * r) + r
		var arr []int

		// sort our big arrays into sub arrays
		if v >= len(elementList) {
			arr = make([]int, len(elementList[(i * r):]))
			copy(arr, elementList[(i * r):])
		} else {
			arr = make([]int, r)
			copy(arr, elementList[(i * r):v])
		}

		// sort our subarray
		sort.Ints(arr)

		// get the median of the subarray
		medians[i] = arr[len(arr)/2]
	}

	// get the median of our medians
	pivot := medianOfMedians(medians, (len(medians) + 1)/2, r)

	var left, right []int

	// split our list by our pivot, where left < pivot and right > pivot
	for i := range elementList {
		if elementList[i] < pivot {
			left = append(left, elementList[i])
		} else if elementList[i] > pivot {
			right = append(right, elementList[i])
		}
	}

	// recurse!
	switch {
	case k == (len(left) + 1):
		return pivot
	case k <= len(left):
		return medianOfMedians(left, k, r)
	default:
		return medianOfMedians(right, k - len(left) - 1, r)
	}
}

func main() {
	arrayzor := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 20}
	arr := make([]int, len(arrayzor))
	copy(arr, arrayzor)
	sort.Ints(arr)

	// find the ith smallest element
	for _, j := range []int{5, 10, 7, 8} {
		// get the ith smallest element
		i := medianOfMedians(arrayzor, j, 5)
		fmt.Println(j, "smallest element = ", i)
		v := arr[j - 1]
		fmt.Println("arr[", j - 1, "] = ", v)
		if i != v {
			fmt.Println("Oops! Algorithm is wrong")
		}
	}
}