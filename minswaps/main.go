package main

import (
	"fmt"
)

func main() {

	prob := []int32{1, 3, 5, 2, 4, 6, 7}

	res := minimumSwaps(prob)
	fmt.Println(res)
}

func minimumSwaps(arr []int32) int32 {

	mapped := make(map[int32]int32)

	for k, v := range arr {
		mapped[v-1] = int32(k)
	}
	fmt.Println(mapped)
	checked := make([]bool, len(arr))

	var res int32
	for k, v := range mapped {
		if checked[k] == true || k == v {
			continue
		}
		var cCount int32
		value := k
		for checked[value] == false {
			checked[value] = true
			value = mapped[value]
			cCount++
		}
		fmt.Println(checked)
		res += cCount - int32(1)
	}
	return res
}
