package main

import (
	"fmt"
	"math/rand"
)

// bubble sort algorithm
// best case O(n)
// worst case O(n^2)
// average case O(n^2)
func main() {

	unsorted_slice := slice_generation(10)
	fmt.Println(unsorted_slice)

	fully_sorted := false

	for !fully_sorted {

		fully_sorted = true

		for i := 1; i < len(unsorted_slice); i++ {

			if unsorted_slice[i-1] > unsorted_slice[i] {
				unsorted_slice[i-1], unsorted_slice[i] = unsorted_slice[i], unsorted_slice[i-1]

				fully_sorted = false
			}
		}
	}

	fmt.Println(unsorted_slice)
}

func slice_generation(slice_count int) []int {
	var new_slice []int
	for i := 0; i < slice_count; i++ {
		new_slice = append(new_slice, rand.Intn(100000))
	}
	return new_slice
}
