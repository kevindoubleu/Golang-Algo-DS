package main

import "fmt"

func insertionSort(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		key := arr[i]
		j := i-1
		for j >= 0 && arr[j] > key {
			// we can't use swaps here bcs j can go to -1
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func insertionSort2(arr []int) {
	for i := range arr {
		// put on very left
		if arr[0] > arr[i] {
			insertBefore(arr, 0, i)
		} else {
			// insert somewhere in the middle between current and start
			for j := 1; j < i; j++ {
				// the equals can be on the left or right of &&
				// its to handle duplicate numbers
				if arr[j-1] <= arr[i] && arr[i] < arr[j] {
					insertBefore(arr, j, i)
				}
			}
		}
	}
}

// moves item at right idx to left idx and shifts left+1 until right forward by 1
func insertBefore(arr []int, left, right int) {
	temp := arr[right]
	copy(arr[left+1:right+1], arr[left:right])
	arr[left] = temp
}

func main() {
	arr1 := []int{5,4,3,2,7,8,5,3}
	insertionSort(arr1)
	fmt.Println(arr1)

	arr2 := []int{5,4,3,2,7,8,5,3,-10}
	insertionSort2(arr2)
	fmt.Println(arr2)
}