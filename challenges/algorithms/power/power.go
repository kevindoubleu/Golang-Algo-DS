package main

import "fmt"

func main() {
	fmt.Println("testing iterative")
	Test(powIter)
	fmt.Println("testing recursive")
	Test(powRec)
}

func Test(functionBeingTested func(int, int) int) {
	type input struct {
		num int
		pow int
	}

	testCases := []struct {
		desc	string
		input	input
		want	int
	}{
		{
			desc: "positive",
			input: input{5, 3},
			want: 125,
		},{
			desc: "negative to negative",
			input: input{-5, 3},
			want: -125,
		},{
			desc: "negative to positive",
			input: input{-5, 2},
			want: 25,
		},{
			desc: "power of 0",
			input: input{5, 0},
			want: 1,
		},
	}
	for _, tC := range testCases {
		retval := functionBeingTested(tC.input.num, tC.input.pow)
		if retval != tC.want {
			fmt.Println(tC.desc)
			fmt.Println("got", retval, "want", tC.want)
		} else {
			fmt.Println("test passed")
		}
	}
}

func powRec(num int, pow int) int {
	if pow == 0 {
		return 1
	}
	if pow == 1 {
		return num
	}
	return num * powRec(num, pow-1)
}

func powIter(num int, pow int) int {
	result := 1
	for i := 0; i < pow; i++ {
		result *= num
	}
	return result
}
