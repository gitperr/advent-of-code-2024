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

func getSecondHalfResult(inputFile string) int {
	var total = 0
	listsToCorrect := findIncorrectLists(inputFile)
	rules := loadRules(inputFile)
	//fmt.Println(listsToCorrect)
	fixedLists := fixIncorrectLists(listsToCorrect, rules)
	//fmt.Println(fixedLists)
	for _, listToAnalyze := range fixedLists {
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

func findIncorrectIndexes(numList []int, rule rule) (bool, int, int) {
	var firstNumIndex int
	var secondNumIndex int
	var incorrectIndexesFound bool = false

	for numIndex, numToCheck := range numList {
		if numToCheck == rule.firstNum {
			firstNumIndex = numIndex
		}

		if numToCheck == rule.secondNum {
			secondNumIndex = numIndex
		}
	}

	if !(firstNumIndex < secondNumIndex) {
		incorrectIndexesFound = true
		return incorrectIndexesFound, firstNumIndex, secondNumIndex
	}

	return incorrectIndexesFound, 0, 0
}

func findIncorrectLists(inputFile string) [][]int {
	var incorrectLists [][]int
	listsToAnalyze := loadListsToAnalyze(inputFile)
	rules := loadRules(inputFile)

	for _, listToAnalyze := range listsToAnalyze {
		checkedList := checkIfIncorrectList(listToAnalyze, rules)
		if checkedList != nil {
			incorrectLists = append(incorrectLists, checkedList)
		}
	}

	return incorrectLists
}

func checkIfIncorrectList(numList []int, rules []rule) []int {
	for _, rule := range rules {
		if !checkBothNumsPresent(numList, rule) {
			continue
		}

		if !compareIndexesOfNumbers(numList, rule) {
			return numList
		}
	}

	return nil
}

func fixIncorrectList(numList []int, rules []rule) ([]int, bool) {
	var incorrectRuleFound bool = false
	for _, rule := range rules {
		if !checkBothNumsPresent(numList, rule) {
			continue
		}

		incorrectIndexFound, indexOne, indexTwo := findIncorrectIndexes(numList, rule)

		if !incorrectIndexFound {
			continue
		}

		incorrectRuleFound = true
		numList[indexOne], numList[indexTwo] = numList[indexTwo], numList[indexOne]
	}

	return numList, incorrectRuleFound
}

func fixListUntilCorrect(numList []int, rules []rule) []int {
	for {
		var incorrectFound bool
		numList, incorrectFound = fixIncorrectList(numList, rules)
		if !incorrectFound {
			break
		}
	}
	return numList
}

func fixIncorrectLists(listsToFix [][]int, rules []rule) [][]int {
	var fixedLists [][]int
	for _, listToFix := range listsToFix {
		fixedList := fixListUntilCorrect(listToFix, rules)
		fixedLists = append(fixedLists, fixedList)
	}

	return fixedLists
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
	//fmt.Println(getFirstHalfResult(inputFile))
	fmt.Println(getSecondHalfResult(inputFile))
}
