package main

import "fmt"

// go until half and check the "backside" of string, O(n/2)
func palindrome2(s string) bool {
	fmt.Println(s)
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		fmt.Println(s[i], "==", s[sLen-i-1], "?")
		if s[i] != s[sLen-i-1] {
			return false
		}
	}
	return true
}

// compare with reverse, O(n)
func palindrome(s string) bool {
	fmt.Println(s, "==", reverse(s), "?")
	return s == reverse(s)
}
func reverse(s string) string {
	var newStr []byte
	for i := len(s)-1; i >= 0; i-- {
		newStr = append(newStr, s[i])
	}
	return string(newStr)
}

func main() {
	fmt.Println(palindrome("mom"))
	fmt.Println(palindrome("maam"))
	fmt.Println(palindrome("madam"))
	fmt.Println(palindrome("madame"))
	fmt.Println(palindrome("racecar"))
	
	fmt.Println(palindrome("abcdedcba"))
	fmt.Println(palindrome("abcdeedcba"))
	fmt.Println(palindrome("abcdeFdcba"))

	fmt.Println(palindrome2("asdfgdsa"))
}