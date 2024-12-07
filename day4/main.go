package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isLastElementInList(charIndex, listLength int) bool {
	return charIndex+1 == listLength
}

func isFirstElementInList(charIndex int) bool {
	return charIndex == 0
}

func checkIfValidXmasString(s string) bool {
	//fmt.Printf("String given to check is %s\n", s)
	if strings.ToLower(s) == "xmas" || strings.ToLower(s) == "samx" {
		//fmt.Printf("MATCH, returning true\n")
		return true
	}

	return false
}

func analyzeToTheLeft(startIndex int, row []string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[startIndex] + row[startIndex-1] + row[startIndex-2] + row[startIndex-3]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeToTheRight(startIndex int, row []string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[startIndex] + row[startIndex+1] + row[startIndex+2] + row[startIndex+3]
	return checkIfValidXmasString(potentialXmasString)
}

func isSafeToAnalyzeToLeft(startIndex int) bool {
	return startIndex-3 >= 0
}

func isSafeToAnalyzeToRight(startIndex, lengthOfRow int) bool {
	return startIndex+3 <= lengthOfRow-1
}

func isSafeToAnalyzeUpwards(rowIndex int) bool {
	return rowIndex-3 >= 0
}

func isSafeToAnalyzeDownwards(rowIndex, lengthOf2DArray int) bool {
	return rowIndex+3 <= lengthOf2DArray-1
}

func analyzeVerticallyDownwards(startIndex int, array2D [][]string) bool {
	return true
}

func analyzeHorizontally(array2D [][]string) int {
	var total = 0
	for _, row := range array2D {
		for charIndex, _ := range row {
			if isLastElementInList(charIndex, len(row)) {
				if analyzeToTheLeft(charIndex, row) {
					total = total + 1
					continue
				}
			}

			if isFirstElementInList(charIndex) {
				if analyzeToTheRight(charIndex, row) {
					total = total + 1
				}
				continue
			}

			if isSafeToAnalyzeToLeft(charIndex) {
				if analyzeToTheLeft(charIndex, row) {
					total = total + 1
				}
			}

			if isSafeToAnalyzeToRight(charIndex, len(row)) {
				if analyzeToTheRight(charIndex, row) {
					total = total + 1
				}
			}
		}
	}
	return total
}

func analyzeVertically(array2D [][]string) int {
	var total = 0
	for _, row := range array2D {
		for charIndex, _ := range row {
			if isLastElementInList(charIndex, len(row)) {
				if analyzeToTheLeft(charIndex, row) {
					total = total + 1
					continue
				}
			}	
}

// Main function
func main() {
	// Open the file
	filePath := "test_input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize the 2D array
	var array2D [][]string

	// Read the file character by character, line by line
	for scanner.Scan() {
		line := scanner.Text()
		charArray := []string{} // Create a slice for this line
		for _, char := range line {
			charArray = append(charArray, string(char))
		}
		array2D = append(array2D, charArray) // Add the line's characters to the 2D array
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	// Print the 2D array to check the results
	fmt.Println("2D Array:")
	for _, row := range array2D {
		fmt.Println(row)
	}

	fmt.Println(analyzeHorizontally(array2D))
}
