package main

import (
	"fmt"
	"strings"
)

func reverseSentence(sentence string) string {
	// splitting into words
	words := strings.Split(sentence, " ")
	var newWords []string
	for _, w := range words {
		newWords = append(newWords, reverseWord(w))
	}
	return strings.Join(newWords, " ")
}

func reverseWord(word string) string {
	var newStr []byte
	for i := len(word)-1; i >= 0; i-- {
		newStr = append(newStr, word[i])
	}

	// if theres punctuation
	// put it at the end where it originally was
	if newStr[0] == '!' || newStr[0] == '.' || newStr[0] == '?' {
		newStr = append(newStr, newStr[0])
		newStr = newStr[1:]
	}
	// doesnt handle symbols other than ! ? .
	// doesnt handle symbols inside words

	return string(newStr)
}

func main() {
	// input
	// sentence: string
	// punctuations: ! ? .

	// output
	// sentence: string

	// split the sentence into words
	// 		and punctuation at the end?
	// reverse every word
	// 		handle punctuations here
	// combine the words back

	// works
	fmt.Println(reverseSentence("Hello World!"))
	// still works
	fmt.Println(reverseSentence("Hello, World!"))
	// problematic
	fmt.Println(reverseSentence("Hello,World!"))

	// what i would improve
	// the splitting would use regex [^a-zA-Z0-9_ ]
	// the punctuation would also use regex to support more symbols
}