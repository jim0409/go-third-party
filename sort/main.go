package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting a slice of strings by length
	strs := []string{"United States", "India", "France", "United Kingdom", "Spain"}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("Sorted strings by length: ", strs)

	// Stable sort
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Println("[Stable] Sorted strings by length: ", strs)

	// Sorting a slice of strings in the reverse order of length
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[j]) < len(strs[i])
	})
	fmt.Println("[Stable] Sorted strings by reverse order of length: ", strs)

	a := []int{5, 3, 4, 7, 8, 9}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		fmt.Println(v)
	}

	doubleArr := [][]int{{2, 6}, {0, 1}, {1, 4}, {3, 2}, {4, 5}}
	sort.Slice(doubleArr, func(i int, j int) bool {
		if doubleArr[i][0] == doubleArr[j][0] {
			return doubleArr[i][1] < doubleArr[j][1]
		}
		return doubleArr[i][0] < doubleArr[j][0]
	})

	fmt.Println(doubleArr)

}
