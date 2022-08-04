package kata

import (
	"math"
	"sort"
)

func binarySearch(haystack []int, needle int) (index int, found bool) {
	if len(haystack) == 0 {
		return -1, false
	}

	if len(haystack) == 1 {
		return 0, haystack[0] == needle
	}

	// Middle is the middle value. Array of 2: middle is 1, array of 3, middle is 2
	middle := int(math.Ceil(float64(len(haystack)) / 2))

	if haystack[middle] == needle {
		return middle, true
	}

	if haystack[middle] > needle {
		resultIndex, isFound := binarySearch(haystack[:middle], needle)
		return resultIndex, isFound
	}

	if haystack[middle] < needle {
		resultIndex, isFound := binarySearch(haystack[middle+1:], needle)
		return middle + 1 + resultIndex, isFound
	}

	return -1, false
}

func TwoSum(numbers []int, target int) [2]int {
	ret := []int{-2, -2}

	// Sort the incoming numbers
	sort.Ints(numbers)

	for firstIndex, firstVal := range numbers {
		// We want to search for values that are after our current index
		searchSlice := numbers[firstIndex+1:]
		searchValue := target - firstVal

		secondIndex, foundValue := binarySearch(searchSlice, searchValue)
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
