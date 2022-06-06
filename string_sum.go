package string_sum

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	fNum := 0
	fNumLen := 0
	fSign := 1
	sNum := 0
	sNumLen := 0
	beginFirst := false

	input = strings.Replace(input, " ", "", -1)
	input_len := len(input)

	if input_len == 0 {
		return "", errorEmptyInput
	}

	for i := 0; i < input_len; i++ {
		if input[input_len-i-1] == 43 {
			if beginFirst == true {
				return "", errorNotTwoOperands
			}
			beginFirst = true
			continue
		}
		if !beginFirst && input[input_len-i-1] >= 48 && input[input_len-i-1] <= 57 {
			sNum = sNum + int(input[input_len-i-1]%48)*int(math.Pow(10, float64(sNumLen)))
			sNumLen = sNumLen + 1
		}
		if beginFirst && input[input_len-i-1] == 45 {
			fSign = -1
		}
		if beginFirst && input[input_len-i-1] >= 48 && input[input_len-i-1] <= 57 {
			fNum = fNum + int(input[input_len-i-1]%48)*int(math.Pow(10, float64(fNumLen)))
			fNumLen = fNumLen + 1
		}
	}

	if beginFirst == false {
		return "", err
	}

	result := fSign*fNum + sNum

	return strconv.Itoa(result), nil
}
