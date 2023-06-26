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
