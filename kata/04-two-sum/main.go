package kata

import (
	"sort"

	"golang.org/x/exp/slices"
)

func TwoSum(numbers []int, target int) [2]int {
	ret := []int{-1, -1}

	// Sort the incoming numbers
	slices.Sort(numbers)

	for firstIndex, firstVal := range numbers {
		// We want to search for values that are after our current index
		searchSlice := numbers[firstIndex+1:]
		searchValue := target - firstVal

		secondIndex, foundValue := slices.BinarySearch(searchSlice, searchValue)
		if foundValue {
			ret[0] = firstIndex
			// The index returned from the searc his relative to our searchSlice's base index
			ret[1] = firstIndex + 1 + secondIndex
			break
		}
	}

	sort.Ints(ret)
	return [2]int{ret[0], ret[1]}
}
