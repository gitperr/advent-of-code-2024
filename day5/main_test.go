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
