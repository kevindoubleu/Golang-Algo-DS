package main

import (
	"fmt"
	"strings"
)

func main() {
	// approach 1
	// 
	// use string builder for result
	// iterate the string
	// if found more than 1 same consecutive char, combo++
	//  else, its different, then
	//   write the combo
	//   restart the combo
	//   write the next char if exist

	Test()
}

func Test() {
	testCases := []struct {
		desc	string
		input	string
		want	string
	}{
		{
			desc: "regular string",
			input: "abbcccddddqwe",
			want: "ab2c3d4qwe",
		},{
			desc: "more than 9 combo",
			input: "aaaaaaaaaab",
			want: "a10b",
		},{
			desc: "empty str",
			input: "",
			want: "",
		},{
			desc: "1 char string",
			input: "a",
			want: "a",
		},{
			desc: "string of 1 repeating char",
			input: "aaaaa",
			want: "a5",
		},
	}
	for _, tC := range testCases {
		retval := compress(tC.input)
		if retval != tC.want {
			fmt.Println("got", retval, "want", tC.want)
		} else {
			fmt.Println("test passed")
		}
	}
}

func compress(str string) string {
	if len(str) == 0 { return str }
	strLen := len(str)

	combo := 1

	result := strings.Builder{}
	result.WriteByte(str[0])

	for i := 1; i < strLen; i++ {
		if str[i] == str[i-1] {
			combo++
			continue
		}

		if combo > 1 {
			result.WriteString(fmt.Sprint(combo))
		}

		result.WriteByte(str[i])
		combo = 1
	}

	if combo > 1 {
		result.WriteString(fmt.Sprint(combo))
	}

	return result.String()
}
