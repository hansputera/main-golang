package main

import "strings"

func findItem(s []string, k string) int {
	for i, v := range s {
		if v == k {
			return i
		}
	}
	return -1
}

func HurufJagaJarak(input string) string {
	var temp []string = []string{}
	stringes := strings.Split(input, "")
	for _, v := range stringes {
		index := findItem(temp, v)
		if index == -1 {
			temp = append(temp, v)
		}
	}
	return strings.Join(temp, "")
}
