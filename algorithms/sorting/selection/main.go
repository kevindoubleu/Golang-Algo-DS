package main

import "fmt"

func selectionSort(arr []int) []int {
	var idxOfMin int
	// for every item in the array
	for i := range arr {
		idxOfMin = i
		// find the smallest item from i until the end
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[idxOfMin] {
				idxOfMin = j
			}
		}
		// then switch it with i
		if idxOfMin != i {
			arr[idxOfMin], arr[i] = arr[i], arr[idxOfMin]
		}
	}
	return arr
}

func main() {
	arr1 := []int{5,4,3,2,7,8,5,3}
	fmt.Println(selectionSort(arr1))
}