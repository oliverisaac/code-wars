package kata

import "sort"

func TwoSum(numbers []int, target int) [2]int {
	ret := []int{-1, -1}

	// Sort the incoming numbers
	sort.Ints(numbers)

	for firstIndex, firstVal := range numbers {
		for secondIndex, secondVal := range numbers {
			// We do not want to use the same index twice
			if secondIndex == firstIndex {
				continue
			}

			if secondVal+firstVal == target {
				ret[0] = firstIndex
				ret[1] = secondIndex
			}
		}
	}

	sort.Ints(ret)
	return [2]int{ret[0], ret[1]}
}
