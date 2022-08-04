package kata

import "strings"

// The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)

	// foundChars tracks where we find duplicate charactesr at
	foundChars := make(map[rune]int)

	ret := make([]rune, 0, len(word))
	for index, ch := range []rune(word) {
		insert := '('

		// If we find a duplicaet character
		if previousOccurance, ok := foundChars[ch]; ok {
			insert = ')'
			// then we also update the previous occurance to match the insert character
			ret[previousOccurance] = insert
		}

		// this is how we track where it was previously set
		foundChars[ch] = index
		ret = append(ret, insert)
	}
	return string(ret)
}
