package main

import "testing"

func TestToUppercase(t *testing.T) {
	allLowerCaseString := "hello world"
	expected := "HELLO WORLD"

	result := ToUppercase(allLowerCaseString)

	if result != expected {
		t.Errorf("ToUppercase (%s) = %s, expected %s", allLowerCaseString, result, expected)
	}

	emptyString := ""
	expectedEmpty := ""

	resultEmpty := ToUppercase(emptyString)

	if resultEmpty != expectedEmpty {
		t.Errorf("ToUppercase(%s) = %s, expected %s", emptyString, resultEmpty, expectedEmpty)
	}

	mixedCaseString := "Hello World"
	expectedMixedCase := "HELLO WORLD"

	resultMixedCase := ToUppercase(mixedCaseString)

	if resultMixedCase != expectedMixedCase {
		t.Errorf("ToUppercase(%s) = %s, expectedMixedCase %s", mixedCaseString, resultMixedCase, expectedMixedCase)
	}

	uppercaseString := "HELLO WORLD"
	expectedUppercaseString := "HELLO WORLD"

	resultUppercaseString := "HELLO WORLD"

	if resultUppercaseString != expectedUppercaseString {
		t.Errorf("ToUppercase(%s) = %s, expectedUppercaseString %s", uppercaseString, resultUppercaseString, expectedUppercaseString)
	}
}

func TestTrim(t *testing.T) {
	inputString := "   hello   "
	expectedString := "hello"

	result := TrimIt(inputString)

	if result != expectedString {
		t.Errorf("Trim(%s) = %s, expectedString %s", inputString, result, expectedString)
	}
}

func TestUpAndDown(t *testing.T) {
	inputString := "hello world"
	expectedString := "HeLlO wOrLd"

	result := AlternateCase(inputString)

	if result != expectedString {
		t.Errorf("UpAndDown(%s) = %s, expectedString %s", inputString, result, expectedString)
	}
}

func TestReverseString(t *testing.T) {
	inputString := "Golang is fun"
	expectedString := "nuf si gnaloG"

	result := ReverseString(inputString)

	if result != expectedString {
		t.Errorf("ReverseString(%s) = %s, expectedString %s", inputString, result, expectedString)
	}

}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"12321", true},
		{"12345", false},
	}

	for _, testCase := range testCases {
		result := IsPalindrome(testCase.input)

		if result != testCase.expected {
			t.Errorf("IsPalindrome(%s) = %t, expected %t", testCase.input, result, testCase.expected)
		}
	}
}
