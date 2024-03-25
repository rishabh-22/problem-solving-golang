package problems

import (
	"strconv"
	"strings"
)

func reverseString(a string) string {
	temp := []rune(a)
	var res []string
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, string(temp[i]))
	}
	return strings.Join(res, "")
}

func IsPalindrome(x int) bool {
	a := strconv.Itoa(x)
	rev := reverseString(a)
	if rev == a {
		return true
	}
	return false
}
