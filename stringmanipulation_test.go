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
