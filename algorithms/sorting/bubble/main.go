package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr1 := []int{5,4,3,2,7,8,5,3}
	bubbleSort(arr1)
	fmt.Println(arr1)
}