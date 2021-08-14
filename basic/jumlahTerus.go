package main

import (
	"strconv"
	"strings"
)

func JumlahTerus(input string) int {
	stringes := strings.Split(input, "")
	var result int = 0
	for _, string_ := range stringes {
		num, err := strconv.Atoi(string_)
		if err == nil {
			result += num
		}
	}

	return result
}
