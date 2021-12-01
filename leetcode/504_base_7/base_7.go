package main

import "strconv"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var isNegative = num < 0
	if isNegative {
		num = num * -1
	}

	result := ""
	for num > 0 {
		result = strconv.Itoa(num%7) + result
		num = num / 7
	}

	if isNegative {
		return "-" + result
	} else {
		return result
	}
}
