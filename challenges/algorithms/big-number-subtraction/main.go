package main

import (
	"fmt"
	"strconv"
)

// top bottom subtraction like in elementary school
func bigNumSub(a, b string) string {
	// assume that input is always valid string of numbers?
	// 		no
	if !validNumberString(a) || !validNumberString(b) {
		fmt.Println("invalid number(s)")
		return "invalid"
	}

	// is the first given number always bigger than the second?
	// 		no
	a, b = biggerNumberStringFirst(a, b)

	// prepare a string to hold result that'll be reversed
	// start from the back
	// check for flag for top digit decrement
	// subtract top num with bottom num
	// 		if enough then just get the result
	// 		if not then borrow from the front, top digit gets added by 10
	// 			a flag gets set to decrement top digit in next loop

	var result []byte
	borrow := false

	fmt.Printf("%20s\n", a)
	fmt.Printf("%20s\n", b)
	fmt.Println("--------------------- -")

	a = string(reverse([]byte(a)))
	b = string(reverse([]byte(b)))

	for i := 0; i < len(b); i++ {
		top, _ := strconv.Atoi(string(a[i]))
		bot, _ := strconv.Atoi(string(b[i]))
		
		if borrow {
			top -= 1
			borrow = false
		}

		if top < bot {
			fmt.Println("borrowing")
			top += 10
			borrow = true
		}

		currentResult := strconv.Itoa(top-bot)
		fmt.Println(top, "-", bot, "=", currentResult)
		result = append(result, byte(currentResult[0]))
	}

	// append the remaining digits from the bigger number
	result = append(result, a[len(b):]...)

	return string(reverse(result))
}

func validNumberString(s string) bool {
	for _, c := range s {
		_, err := strconv.Atoi(string(c))
		if err != nil {
			return false
		}
	}
	return true
}

// returns the bigger number string as a, and smaller as b
// assumes x and y are valid number strings
func biggerNumberStringFirst(x, y string) (a, b string) {
	if len(x) < len(y) {
		return y, x
	} else if len(x) > len(y) {
		return x, y
	}
 
	// len of both number is the same
	for i := 0; i < len(x); i++ {
		xDigit, _ := strconv.Atoi(string(x[i]))
		yDigit, _ := strconv.Atoi(string(y[i]))
		if xDigit > yDigit {
			return x, y
		}
	}
	return y, x
}

func reverse(s []byte) []byte {
	var newStr []byte
	sLen := len(s)
	for i := sLen-1; i >= 0; i-- {
		newStr = append(newStr, s[i])
	}
	return newStr
}

func main() {
	// 1111
	fmt.Println(bigNumSub("5432", "4321"))

	// 11442188888888889989
	fmt.Println(bigNumSub("11443333311111111100", "1144422222221111"))

	// 122356441111911120
	fmt.Println(bigNumSub("122387876566565674", "31435454654554"))

	// invalid
	fmt.Println(bigNumSub("1", "zero"))

	// reversed position
	fmt.Println(bigNumSub("0", "123"))
}