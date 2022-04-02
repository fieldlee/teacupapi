package utils

import (
	"regexp"
	"strconv"
)

type Validator struct {
	Min   int
	Max   int
	Field string
	Value string
	Flags string
}

func checkEmail(str string) bool {
	ma, err := regexp.MatchString("^[A-Za-z\\d]+([-_.][A-Za-z\\d]+)*@([A-Za-z\\d]+[-.])+[A-Za-z\\d]{2,4}$", str)
	if err != nil {
		return false
	}
	return ma
}

func checkBool(str string) bool {

	_, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return true
}

func checkFloat(str string) bool {

	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

func checkLength(str string, min, max int) bool {

	if min == 0 && max == 0 {
		return true
	}

	n := len(str)
	if n < min || n > max {
		return false
	}

	return true
}
