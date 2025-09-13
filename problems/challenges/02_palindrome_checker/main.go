package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(text string) (bool, string) {
	temp := ""
	// for i := 0; i < len(text); i++ {
	// 	if text[i] >= 'A' && text[i] <= 'Z' {
	// 		temp += string(text[i] + 32)
	// 	} else if text[i] >= 'a' && text[i] <= 'z' {
	// 		temp += string(text[i])
	// 	}
	// }
	for _, ch := range text {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			temp += string(ch)
		} else {
			temp += ""
		}
	}
	// another way using strings package
	// temp = strings.ReplaceAll(temp, " ", "")
	temp = strings.ToLower(temp)
	i := 0
	j := len(temp) - 1
	for i < j {
		if temp[i] != temp[j] {
			return false, temp
		}
		i++
		j--
	}
	return true, temp
}

func main() {
	fmt.Println(isPalindrome("madam"))
	fmt.Println(isPalindrome("hello"))
	fmt.Println(isPalindrome("Taco Cat"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("No 'x' 121 x on"))
}
