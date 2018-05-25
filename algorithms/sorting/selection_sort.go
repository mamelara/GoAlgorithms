package main

import (
	"fmt"
)

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func SelectionSort(s []int) {

	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		swap(s, i, min)
	}
}

func main() {

	arr := []int{1, 13, 20, 2, 25}

	fmt.Println("Unsorted array ", arr)
	SelectionSort(arr)
	fmt.Println("Sorted array ", arr)
}
