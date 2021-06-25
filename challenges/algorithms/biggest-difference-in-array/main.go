package main

import "fmt"

func biggestDifference(array []int) int {
	if len(array) <= 1 {
		return 0
	}

	lastIdx := len(array)-1

	biggest := array[lastIdx]
	smallest := array[lastIdx]
	maxDiff := 0

	for i := lastIdx-1; i >= 0; i-- {
		if array[i] > biggest {
			biggest = array[i]
			smallest = array[i]
		}

		if array[i] < smallest {
			smallest = array[i]
		}

		if biggest - smallest > maxDiff {
			maxDiff = biggest - smallest
		}
	}

	return maxDiff
}

func main() {
	// 1st approach
	// brute force approach
	// count differences between each element
	// double for loops and just like selection sort
	// O(n^2)

	// 2nd approach
	// use 3 helper variables: biggest, smallest, biggest_diff
	// we traverse from the right
	// for every item
	// 		check with biggest
	// 			if we update biggest, smallest must also get updated
	// 			because smallest must come before biggest
	// 		check with smallest
	// 		check current biggest - smallest with current biggest_diff
	// O(n)

	// guarantee that there will be a pair?
	// 		no
	// arr with no pair like [2, 1]?
	// 		returns 0
	// empty array?
	// 		returns nil
	// array with 1 element?
	// 		returns 0, bcs x - x = 0

	// 7
	fmt.Println(biggestDifference([]int{ 2, 7, 9, 5, 1, 3, 5 }))

	// 6
	fmt.Println(biggestDifference([]int{ 10, 3, 6, 8, 9, 4, 3 }))
}