package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	firstNum  int
	secondNum int
}

func getFirstHalfResult(inputFile string) int {
	var total = 0
	listsToAnalyze := loadListsToAnalyze(inputFile)
	rules := loadRules(inputFile)

	for _, listToAnalyze := range listsToAnalyze {
		total = total + analyzeList(listToAnalyze, rules)
	}
	return total
}

func loadRules(filePath string) []rule {
	var listOfRules []rule

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			splitLine := strings.Split(line, "|")
			firstNum, _ := strconv.Atoi(splitLine[0])
			secondNum, _ := strconv.Atoi(splitLine[1])
			var newRule = rule{
				firstNum:  firstNum,
				secondNum: secondNum,
			}
			listOfRules = append(listOfRules, newRule)
		}
	}

	return listOfRules
}

func loadListsToAnalyze(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var listsToAnalyze [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			splitLine := strings.Split(line, ",")
			var newList []int
			for _, numStr := range splitLine {
				numInt, _ := strconv.Atoi(numStr)
				newList = append(newList, numInt)
			}
			listsToAnalyze = append(listsToAnalyze, append([]int(nil), newList...))
		}
	}

	return listsToAnalyze
}

func getMiddleNumberInList(numList []int) int {
	return numList[(len(numList)-1)/2]
}

func checkBothNumsPresent(numList []int, rule rule) bool {
	var firstNumPresent bool
	var secondNumPresent bool

	for _, numToCheck := range numList {
		if numToCheck == rule.firstNum {
			firstNumPresent = true
		}

		if numToCheck == rule.secondNum {
			secondNumPresent = true
		}
	}

	if firstNumPresent && secondNumPresent {
		return true
	}

	return false
}

func compareIndexesOfNumbers(numList []int, rule rule) bool {
	var firstNumIndex int
	var secondNumIndex int
	for numIndex, numToCheck := range numList {
		if numToCheck == rule.firstNum {
			firstNumIndex = numIndex
		}

		if numToCheck == rule.secondNum {
			secondNumIndex = numIndex
		}

	}

	return firstNumIndex < secondNumIndex
}

func analyzeList(numList []int, rules []rule) int {
	var middleNumber = 0
	for _, rule := range rules {
		if !checkBothNumsPresent(numList, rule) {
			continue
		}

		if !compareIndexesOfNumbers(numList, rule) {
			return middleNumber
		}
	}

	return getMiddleNumberInList(numList)
}

func main() {
	inputFile := "input.txt"
	fmt.Println(getFirstHalfResult(inputFile))
}
