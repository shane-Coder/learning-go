package stringutils

import "strings"

func Reverse(s string) string {
	// var temp string
	// l := len(s)
	// for l > 0 {
	// 	temp += string(s[l-1])
	// 	l--
	// }
	// s = temp
	// return s

	// another way
	// Convert string to a slice of runes to handle multi-byte characters
	runes := []rune(s)
	// The swapping logic is more efficient
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ToUpper(s string) string {
	// logic 1 -> ASCII value of 'a' is 97 and 'A' is 65
	temp := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			temp += string(c - 32)
		} else {
			temp += string(c)
		}
	}
	s = temp
	// logic 2 -> using strings package
	// s = strings.ToUpper(s)
	return s
}

func ToLower(s string) string {
	var temp string
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			temp += string(ch + 32)
		} else {
			temp += string(ch)
		}
	}
	return temp

	// return strings.ToLower(s) // another way using strings package
}

func IsPalindrome(s string) (bool, string) {
	// temp := strings.TrimSpace(ToLower(s)) // another way using strings package
	temp := strings.ReplaceAll(s, " ", "")
	temp = ToLower(temp)
	s = temp
	pal := Reverse(temp)
	pal = ToLower(pal)
	if pal == s {
		return true, pal
	} else {
		return false, pal
	}
}
