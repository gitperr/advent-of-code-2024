package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestGetFirstHalfResult(t *testing.T) {
	filePath := "test_input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array2D [][]string

	for scanner.Scan() {
		line := scanner.Text()
		charArray := []string{}
		for _, char := range line {
			charArray = append(charArray, string(char))
		}
		array2D = append(array2D, charArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	expectedResult := 18
	result := getFirstHalfResult(array2D)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}

func TestGetSecondHalfResult(t *testing.T) {
	filePath := "test_input2.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var array2D [][]string

	for scanner.Scan() {
		line := scanner.Text()
		charArray := []string{}
		for _, char := range line {
			charArray = append(charArray, string(char))
		}
		array2D = append(array2D, charArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}
	expectedResult := 9
	result := getSecondHalfResult(array2D)

	if result != expectedResult {
		t.Errorf("Expected %d, but got %d", expectedResult, result)
	}
}
