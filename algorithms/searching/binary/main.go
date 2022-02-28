package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{23, 43, 65, 7, 21, 3, 652, 2}
	sort.Ints(arr) // binary search only work on sorted arrays
	fmt.Println(arr)

	findThese := []int{1, 2, 3, 65, 23, 43, 65, 7, 21, 3, 652, 2}
	for _, target := range findThese {
		fmt.Println("searching for", target, "got", search(arr, target))
	}
}

func search(arr []int, target int) int {
	left, right := 0, len(arr)-1
	mid := 0

	for left <= right {
		mid = (right + left) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			// target is on the right
			left = mid+1
		} else if arr[mid] > target {
			// target is on the left
			right = mid-1
		}
	}
	return -1
}
