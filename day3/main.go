package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calc(mul string) int {
	var modifiedStr = ""
	var total = 1
	if len(mul) > 4 {
		modifiedStr = mul[4 : len(mul)-1]
		//fmt.Println("Original string:", mul)
		//fmt.Println("Modified string:", modifiedStr)
	} else {
		//fmt.Println("String is too short to modify.")
	}

	mulNums := strings.Split(modifiedStr, ",")

	for _, num := range mulNums {
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("Error occurred while converting str to int")
		}
		total = total * number
	}
	return total
}

func getFirstHalfResult(listOfMuls []string) int {
	var total = 0
	for _, mul := range listOfMuls {
		total += calc(mul)
	}

	return total
}

func getSecondHalfResult(listOfMuls []string) int {
	var total = 0
	var do = true
	for _, mul := range listOfMuls {
		if mul == "don't()" {
			do = false
		}

		if mul == "do()" {
			do = true
		}

		if do {
			if strings.Contains(mul, "mul") {
				total += calc(mul)
			}
		}
	}

	return total

}

func main() {
	filePath := "input.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	pattern := `mul\([0-9]+,[0-9]+\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(string(content), -1)

	fmt.Println(getFirstHalfResult(matches))

	pattern = `mul\([0-9]+,[0-9]+\)|do(n't)?\(\)`
	re = regexp.MustCompile(pattern)

	matches = re.FindAllString(string(content), -1)

	fmt.Println(getSecondHalfResult(matches))
}
