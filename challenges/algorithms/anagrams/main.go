package main

import (
	"fmt"
)

func mapByteIntEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for char, count := range a {
		// fmt.Println(char, count)
		if b[char] != count {
			return false
		}
	}
	return true
	// return reflect.DeepEqual(a, b)
}

func makeCharset(s string) map[byte]int {
	charset := make(map[byte]int)
	for _, c := range []byte(s) {
		if charset[c] == 0 {
			charset[c] = 1
		} else {
			charset[c] += 1
		}
	}
	return charset
}

func anagram(a, b string) bool {
	charsetA := makeCharset(a)
	charsetB := makeCharset(b)

	return mapByteIntEqual(charsetA, charsetB)
}

func main() {
	fmt.Println(anagram("hello", "olleh"))
	fmt.Println(anagram("hello", "olleh."))
}