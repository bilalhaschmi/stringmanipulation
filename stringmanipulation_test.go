package main

import "testing"

func TestToUppercase(t *testing.T) {
	inputString := "hello world"
	expected := "HELLO WORLD"

	result := ToUppercase(inputString)

	if result != expected {
		t.Errorf("ToUppercase (%s) = %s, expected %s", inputString, result, expected)
	}

	emptyString := ""
	expectedEmpty := ""

	resultEmpty := ToUppercase(emptyString)

	if resultEmpty != expectedEmpty {
		t.Errorf("ToUppercase(%s) = %s, expected %s", emptyString, resultEmpty, expectedEmpty)
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

	result := UpAndDown(inputString)

	if result != expectedString {
		t.Errorf("UpAndDown(%s) = %s, expectedString %s", inputString, result, expectedString)
	}
}
