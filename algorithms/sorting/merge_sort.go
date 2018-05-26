package main

import "fmt"

func sort(firstHalf []int, secondHalf []int, mergedArray []int) []int {

	var i, j, k = 0, 0, 0

	first := make([]int, len(firstHalf))
	second := make([]int, len(secondHalf))

	copy(first, firstHalf)
	copy(second, secondHalf)

	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			mergedArray[k] = first[i]
			i++
		} else {
			mergedArray[k] = second[j]
			j++
		}
		k++
	}

	for i < len(firstHalf) {
		mergedArray[k] = first[i]
		k++
		i++
	}

	for j < len(secondHalf) {
		mergedArray[k] = second[j]
		j++
		k++
	}

	return mergedArray

}

func merge(array []int) []int {

	if len(array) == 1 {
		return array
	}

	var mid int = len(array) / 2

	firstHalf := merge(array[:mid])
	secondHalf := merge(array[mid:])
	array = sort(firstHalf, secondHalf, array)
	return array
}

func MergeSort(array []int) []int {
	return merge(array)
}

func main() {

	array := []int{3, 2, 10, 1, 5, 4, 6, 7, 9, 8}
	array = MergeSort(array)
	fmt.Println(array)

}
