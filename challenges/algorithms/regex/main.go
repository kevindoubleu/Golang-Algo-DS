package main

import (
	"fmt"
	"regexp"
)

func extract(s string) string {
	re := regexp.MustCompile(`tokopedia.com/discovery/(.*)/`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return "no matches"
	}
	return matches[1]
}

func main() {
	fmt.Println(extract("http://tokopedia.com/discovery/get-this-string/new"))
	fmt.Println(extract("http://tokopedia.com/discovery/hohoho/new"))
	fmt.Println(extract("http://google.com/discovery/hohoho/new"))
}