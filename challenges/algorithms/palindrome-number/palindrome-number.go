package main

import "fmt"

func main() {
	Test()
}

func Test() {
	testCases := []struct {
		desc	string
		input	int
		want	bool
	}{
		{
			desc: "even len is pal",
			input: 2332,
			want: true,
		},{
			desc: "even len is not pal",
			input: 233244,
			want: false,
		},{
			desc: "odd len is pal",
			input: 23332,
			want: true,
		},{
			desc: "odd len is not pal",
			input: 23322,
			want: false,
		},{
			desc: "1 digit",
			input: 5,
			want: true,
		},{
			desc: "1 digit 0",
			input: 0,
			want: true,
		},
	}
	for _, tC := range testCases {
		retval := isPalindromeNum(tC.input)
		if retval != tC.want {
			fmt.Println(tC.desc)
			fmt.Println("got", retval, "want", tC.want)
		} else {
			fmt.Println("test passed")
		}
	}
}

func isPalindromeNum(num int) bool {
	return num == reverseNumber(num)
}

func reverseNumber(num int) int {
	// take last digit from num
	// reversed *= 10
	// reversed += last digit from num
	// repeat until num is 0

	reversed := 0
	for num > 0 {
		lastDigit := num % 10
		num /= 10

		reversed *= 10
		reversed += lastDigit
	}
	return reversed
}
