package kata

import "sort"

func TwoSum(numbers []int, target int) [2]int {
	ret := []int{-1, -1}

	// Sort the incoming numbers
	sort.Ints(numbers)

	sort.Ints(ret)
	return [2]int{ret[0], ret[1]}
}
