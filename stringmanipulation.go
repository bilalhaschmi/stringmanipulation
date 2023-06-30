package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "Hello World"
	result := AlternateCase(input)
	fmt.Println(result) // Output: HeLlO wOrLd
}

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

func ReverseString(input string) string {
	bytes := []byte(input)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	reversed := string(bytes)

	return reversed
}

func IsPalindrome(input string) bool {
	sanitized := sanitizeString(input)

	reversed := ReverseString(sanitized)

	return sanitized == reversed
}

func sanitizeString(input string) string {
	var sanitized strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			sanitized.WriteRune(unicode.ToLower(char))
		}
	}

	return sanitized.String()
}

func CountWords(input string) int {
	words := strings.Fields(input)

	return len(words)
}
