package main

import "testing"

func TestGetFirstHalfResult(t *testing.T) {
	input := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	expectedResult := 161
	result := getFirstHalfResult(input)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}

func TestGetSecondHalfResult(t *testing.T) {
	input := []string{"mul(2,4)", "mul(8,5)"}
	expectedResult := 48
	result := getSecondHalfResult(input)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}
