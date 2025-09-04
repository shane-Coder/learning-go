package stringutils

func Reverse(s string) string {
	var temp string
	l := len(s)
	for l > 0 {
		temp += string(s[l-1])
		l--
	}
	s = temp
	return s
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
