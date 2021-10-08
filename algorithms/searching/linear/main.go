package main

import "fmt"

func main() {
	arr := []int{23, 43, 65, 7, 21, 3, 652, 2}

	findThese := []int{1, 2, 3, 65}
	for _, target := range findThese {
		fmt.Println("searching for", target, "got", search(arr, target))
	}
}

// returns index of item or -1 if not found
func search(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
