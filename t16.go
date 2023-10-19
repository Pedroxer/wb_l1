package main

import "fmt"

// Худший - O(n^2). Средний O(n*logn)

func partition(left, right int, arr []int) int {
	pivot := arr[right] // опорный элемент. От опорного элемента зависит эффективность алгоритма.
	wall := left
	for i := left; i < right; i++ {
		if arr[i] <= pivot {
			arr[i], arr[wall] = arr[wall], arr[i]
			wall++
		}
	}
	arr[wall], arr[right] = arr[right], arr[wall]
	return wall
}

func quicksort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(start, end, arr)
	quicksort(arr, start, pivot-1)
	quicksort(arr, pivot+1, end)
}

func main() {
	a := []int{3, 1, 5, 4, 2, 6}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
