package main

import "strings"

func clean(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}

func reverseText(input string) string {
	byte_str := []rune(input)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func cleanText(input string) string {
	return clean([]byte(strings.TrimSpace(input)))
}
func Palindrome(input string) bool {
	cleanCode := cleanText(strings.ToLower(input))
	return cleanCode == reverseText(cleanCode)
}
