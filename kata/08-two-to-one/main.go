package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	runeList := []rune(s1)
	runeList = append(runeList, []rune(s2)...)

	keys := make(map[rune]bool)
	result := []string{}
	for _, eachValue := range runeList {
		if _, exists := keys[eachValue]; !exists {
			result = append(result, string(eachValue))
			keys[eachValue] = true
		}
	}

	sort.Strings(result)
	return strings.Join(result, "")
}
