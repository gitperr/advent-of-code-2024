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

func checkIfValidMasString(s string) bool {
	if strings.ToLower(s) == "mas" || strings.ToLower(s) == "sam" {
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

// ------

func isSafeToAnalyzeToLeftMas(startIndex int) bool {
	return startIndex-1 >= 0
}

func isSafeToAnalyzeToRightMas(startIndex, lengthOfRow int) bool {
	return startIndex+1 <= lengthOfRow-1
}

func isSafeToAnalyzeUpwardsMas(rowIndex int) bool {
	return rowIndex-1 >= 0
}

func isSafeToAnalyzeDownwardsMas(rowIndex, lengthOf2DArray int) bool {
	return rowIndex+1 <= lengthOf2DArray-1
}

func analyzeVerticallyDownwards(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex+1][rowStartIndex] + array2D[rowIndex+2][rowStartIndex] + array2D[rowIndex+3][rowStartIndex]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeVerticallyUpwards(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex-1][rowStartIndex] + array2D[rowIndex-2][rowStartIndex] + array2D[rowIndex-3][rowStartIndex]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeDownRight(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex+1][rowStartIndex+1] + array2D[rowIndex+2][rowStartIndex+2] + array2D[rowIndex+3][rowStartIndex+3]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeDownLeft(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex+1][rowStartIndex-1] + array2D[rowIndex+2][rowStartIndex-2] + array2D[rowIndex+3][rowStartIndex-3]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeUpRight(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex-1][rowStartIndex+1] + array2D[rowIndex-2][rowStartIndex+2] + array2D[rowIndex-3][rowStartIndex+3]
	return checkIfValidXmasString(potentialXmasString)
}

func analyzeUpLeft(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialXmasString = ""
	potentialXmasString += row[rowStartIndex] + array2D[rowIndex-1][rowStartIndex-1] + array2D[rowIndex-2][rowStartIndex-2] + array2D[rowIndex-3][rowStartIndex-3]
	return checkIfValidXmasString(potentialXmasString)
}

func checkForXmas(row []string, rowIndex int, rowStartIndex int, array2D [][]string) bool {
	var potentialMasStringOne = ""
	var potentialMasStringTwo = ""
	potentialMasStringOne += array2D[rowIndex-1][rowStartIndex-1] + row[rowStartIndex] + array2D[rowIndex+1][rowStartIndex+1]
	potentialMasStringTwo += array2D[rowIndex-1][rowStartIndex+1] + row[rowStartIndex] + array2D[rowIndex+1][rowStartIndex-1]
	if checkIfValidMasString(potentialMasStringOne) && checkIfValidMasString(potentialMasStringTwo) {
		return true
	}
	return false
}

func getSecondHalfResult(array2D [][]string) int {
	var total = 0
	for rowIndex, row := range array2D {
		for charIndex, char := range row {
			if char == "A" {
				//fmt.Println("Char is A, safety check next")
				if isSafeToAnalyzeDownwardsMas(rowIndex, len(array2D)) && isSafeToAnalyzeToLeftMas(charIndex) && isSafeToAnalyzeUpwardsMas(rowIndex) && isSafeToAnalyzeToRightMas(charIndex, len(row)) {
					//fmt.Println("Safe to analyze")
					if checkForXmas(row, rowIndex, charIndex, array2D) {
						total = total + 1
					}
				}
			}
		}
	}
	return total
}

func analyzeDiagonally(array2D [][]string) int {
	var total = 0
	for rowIndex, row := range array2D {
		for charIndex, char := range row {
			if char != "X" {
				continue
			}
			if isSafeToAnalyzeDownwards(rowIndex, len(array2D)) && isSafeToAnalyzeToLeft(charIndex) {
				if analyzeDownLeft(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}

			if isSafeToAnalyzeDownwards(rowIndex, len(array2D)) && isSafeToAnalyzeToRight(charIndex, len(row)) {
				if analyzeDownRight(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}

			if isSafeToAnalyzeUpwards(rowIndex) && isSafeToAnalyzeToLeft(charIndex) {
				if analyzeUpLeft(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}

			if isSafeToAnalyzeUpwards(rowIndex) && isSafeToAnalyzeToRight(charIndex, len(row)) {
				if analyzeUpRight(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}
		}
	}

	return total
}

func analyzeHorizontally(array2D [][]string) int {
	var total = 0
	for _, row := range array2D {
		for charIndex, char := range row {
			if char != "X" {
				continue
			}
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
	for rowIndex, row := range array2D {
		for charIndex, char := range row {
			if char != "X" {
				continue
			}
			if isSafeToAnalyzeDownwards(rowIndex, len(array2D)) {
				if analyzeVerticallyDownwards(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}

			if isSafeToAnalyzeUpwards(rowIndex) {
				if analyzeVerticallyUpwards(row, rowIndex, charIndex, array2D) {
					total = total + 1
				}
			}
		}
	}
	return total
}

func getFirstHalfResult(array2D [][]string) int {
	var total = 0
	total = total + analyzeHorizontally(array2D) + analyzeVertically(array2D) + analyzeDiagonally(array2D)
	return total
}

func main() {
	filePath := "input.txt"
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

	fmt.Println(getFirstHalfResult(array2D))
	fmt.Println(getSecondHalfResult(array2D))
}
