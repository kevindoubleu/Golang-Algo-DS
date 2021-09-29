package main

import "fmt"

// wrapper function
func mergeSort(arr []int) {
	copy(arr, _mergeSort(arr))
}

// merge sort is a divide and conquer algo
// this is the divide and conquer part
func _mergeSort(arr []int) []int {
	fmt.Println("mergesort called with", arr)
	arrLen := len(arr)
	if arrLen > 2 {
		// the actual divide part of divide and conuer
		// we "clone" ourselves and work on halves of the original arr
		arr = _merge(_mergeSort(arr[:arrLen/2]), _mergeSort(arr[arrLen/2:]))
	} else if arrLen == 2 {
		// we keep "cloning" until we only have 2 items
		// then we sort them simply
		// the actual conquer part of divide and conquer
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
	}
	// if we only have 1 item then just return it bcs it's sorted

	// return back the sorted array of 2 elements to be merged by _merge
	return arr
}

// this is the combining part (after divide and conquer part)
func _merge(left, right []int) []int {
	// here we know that the arrays left and right are both sorted
	var result []int
	// until 1 or the arrays is exhausted
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// here, 1 of the arrays or both have been exhausted
	// since we know that left and right is sorted
	// we can just put the remaining items into the results, no worries
	if len(left) > 0 {
		result = append(result, left...)
	} else if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}

func main() {
	arr1 := []int{5,4,3,2,7,8,5,3,-10}
	mergeSort(arr1)
	fmt.Println(arr1)
}