package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToUppercase(input string) string {
	return strings.ToUpper(input)
}

func TrimIt(input string) string {
	return strings.TrimSpace(input)
}

func AlternateCase(input string) string {
	result := ""
	previousIsUpper := false

	for _, char := range input {
		if unicode.IsSpace(char) {
			result += string(char)
		} else {
			if previousIsUpper {
				result += strings.ToLower(string(char))
			} else {
				result += strings.ToUpper(string(char))
			}

			previousIsUpper = !previousIsUpper
		}
	}
	return result
}

func main() {
	input := "Hello World"
	result := AlternateCase(input)
	fmt.Println(result) // Output: HeLlO wOrLd
}
