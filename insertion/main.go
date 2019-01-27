package main

import (
	"fmt"
	"math/rand"
)

// insertion sort algorithm
// best case O(n)
// worst case O(n^2)
// average case O(n^2)
func main() {
	unsorted_slice := slice_generation(10)
	fmt.Println(unsorted_slice)
	for i := 1; i < len(unsorted_slice); i++ {

		j := i

		for j > 0 && unsorted_slice[j-1] > unsorted_slice[j] {
			unsorted_slice[j-1], unsorted_slice[j] = unsorted_slice[j], unsorted_slice[j-1]
			j--
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
