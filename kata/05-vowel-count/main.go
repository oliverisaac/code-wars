package kata

import "regexp"

func GetCount(str string) (count int) {
	regex, _ := regexp.Compile("[aeiou]")
	return len(str) - len(regex.ReplaceAllString(str, ""))
}
