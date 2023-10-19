package main

import "fmt"

func binary_search(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] < target {
			left = middle + 1
		} else {
			return middle
		}

	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(binary_search(a, 5))
}
