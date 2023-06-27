package main

import "strings"

func ToUppercase(input string) string {
	return strings.ToUpper(input)
}

func TrimIt(input string) string {
	return strings.TrimSpace(input)
}

func UpAndDown(input string) string {
	result := ""
	for i, char := range input {
		if i%2 == 0 {
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
	}
	return result
}
