package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func findSmallestNumberInSlice(y []int) (int, []int) {
	var min = y[0]
	var index = 0
	for i, v := range y {
		if min > v {
			min = v
			index = i
		}
	}

	y = append(y[:index], y[index+1:]...)

	return min, y
}

func findOccurrencesOfNumberInList(num int, y []int) int {
	var occurrences = 0
	for _, v := range y {
		if num == v {
			occurrences = occurrences + 1
		}
	}

	return occurrences
}

func answerFirstHalf(firstList []int, secondList []int) int {
	var totalDistance = 0
	for len(firstList) > 0 {
		min, newSlice := findSmallestNumberInSlice(firstList)
		firstList = newSlice

		secondMin, secondSlice := findSmallestNumberInSlice(secondList)
		secondList = secondSlice

		totalDistance = totalDistance + distance(min, secondMin)
	}
	return totalDistance
}

func answerSecondHalf(firstList []int, secondList []int) int {
	var similarityScore = 0

	for _, v := range firstList {
		similarityScore = similarityScore + (v * findOccurrencesOfNumberInList(v, secondList))
	}

	return similarityScore
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	var firstList []int
	var secondList []int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Error: Each line must contain exactly two numbers")
			return
		}

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting first number to int:", err)
			return
		}

		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting second number to int:", err)
			return
		}

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)
	}

	// After answering the first half I commented this out because it modifies the slices (removes items)
	// fmt.Println(answerFirstHalf(firstList, secondList)) // Answer --> 1834060

	fmt.Println(answerSecondHalf(firstList, secondList)) // Answer --> 21607792

}
