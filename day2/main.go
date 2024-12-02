package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileLinesAs2DArray(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		strFields := strings.Fields(line)

		var intFields []int
		for _, str := range strFields {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, fmt.Errorf("failed to convert %q to integer: %w", str, err)
			}
			intFields = append(intFields, num)
		}

		lines = append(lines, intFields)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func distance(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func analyzeLineWithRemoval(line []int) bool {
	if analyzeLine(line) {
		return true
	}

	for i := range line {
		modifiedLine := make([]int, 0, len(line)-1)
		modifiedLine = append(modifiedLine, line[:i]...)
		modifiedLine = append(modifiedLine, line[i+1:]...)

		if analyzeLine(modifiedLine) {
			return true
		}
	}

	return false
}

func analyzeLine(line []int) bool {
	var decreasing = false
	var increasing = false
	for i, num := range line {
		if i+1 >= len(line) {
			break
		}
		nextNum := line[i+1]
		if num > nextNum && !increasing {
			decreasing = true
		}

		if num < nextNum && !increasing {
			increasing = true
		}
		if distance(num, nextNum) > 3 {
			return false
		} else if distance(num, nextNum) < 1 {
			return false
		}

		if decreasing && num < nextNum {
			return false
		}

		if increasing && num > nextNum {
			return false
		}
	}

	return true
}

func main() {
	var firstHalfSafeReports = 0
	var secondHalfSafeReports = 0
	filename := "input.txt"

	lines, err := readFileLinesAs2DArray(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	for _, line := range lines {
		// first half
		if analyzeLine(line) {
			firstHalfSafeReports = firstHalfSafeReports + 1
		}

		// second half
		if analyzeLineWithRemoval(line) {
			secondHalfSafeReports = secondHalfSafeReports + 1
		}
	}
	fmt.Printf("Undampened safe report count: %v\n", firstHalfSafeReports) // -> 279
	fmt.Printf("Dampened safe report count: %v\n", secondHalfSafeReports)  // -> 343
}
