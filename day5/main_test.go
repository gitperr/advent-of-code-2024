package main

import (
	"testing"
)

func TestGetFirstHalfResult(t *testing.T) {
	inputFile := "test_input.txt"
	expectedResult := 143
	result := getFirstHalfResult(inputFile)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}

func TestGetSecondHalfResult(t *testing.T) {
	inputFile := "test_input.txt"
	expectedResult := 123
	result := getSecondHalfResult(inputFile)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}
