package main

import (
	"errors"
	"fmt"
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
	plusesCount := 0
	plusEncounterIndex := 0
	minusesCount := 0

	input = strings.Replace(input, " ", "", -1)
	input_len := len(input)

	if input_len == 0 {
		return "", errorEmptyInput
	}

	for i := 0; i < input_len; i++ {
		if input[i] == 43 {
			plusesCount = plusesCount + 1
			plusEncounterIndex = i
		}
		if input[i] == 45 {
			minusesCount = minusesCount + 1
		}
	}

	if plusesCount > 1 || plusEncounterIndex == 0 || minusesCount > 2 {
		return "", errorNotTwoOperands
	}

	nums := strings.Split(input, "+")
	result := 0

	for i := 0; i < len(nums); i++ {
		n, merr := strconv.Atoi(nums[i])
		if merr == nil {
			result = result + n
		} else {
			return "", errorNotTwoOperands
		}

	}

	return strconv.Itoa(result), nil
}

func main() {
	s := "-334+-455"
	fmt.Println(StringSum(s))
}
