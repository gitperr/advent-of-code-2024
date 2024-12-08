package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	right string = ">"
	left  string = "<"
	up    string = "^"
	down  string = "v"
)

type coordinate struct {
	x int
	y int
}

type guard struct {
	location  coordinate
	direction string
}

func getFirstHalfResult(inputFile string) int {
	var total = 0

	field := loadField(inputFile)
	guard := locateGuardInField(field)
	for {
		fmt.Println("Next obstacle at x,y:")
		obstacle := findNextObstacleOfGuard(field, guard)
		fmt.Println(obstacle)
		if obstacle.x == -99 && obstacle.y == -99 {
			fmt.Println("No other obstacles left")
			break
		}
		//fmt.Println(field)
		fmt.Println(guard)
		field, guard = moveGuardUpUntilObstacle(field, guard, obstacle)
		// TODO:these are not implemented yet, and it's breaking the program
		// moveGuardRightUntilObstacle
		// moveGuardLeftUntilObstacle
		// moveGuardDownUntilObstacle
		fmt.Println(guard)
		//fmt.Println(field)
	}
	return total
}

func turnGuardNinetyDegreesToRight(guard guard) guard {
	switch guard.direction {
	case left:
		guard.direction = up
	case up:
		guard.direction = right
	case right:
		guard.direction = down
	case down:
		guard.direction = left
	}

	return guard
}

func moveGuardRightOnce(guard guard) guard {
	guard.location = coordinate{x: guard.location.x + 1, y: guard.location.y}

	return guard
}

func moveGuardLeftOnce(guard guard) guard {
	guard.location = coordinate{x: guard.location.x - 1, y: guard.location.y}

	return guard
}

func moveGuardUpOnce(guard guard) guard {
	guard.location = coordinate{x: guard.location.x, y: guard.location.y - 1}

	return guard
}

func moveGuardDownOnce(guard guard) guard {
	guard.location = coordinate{x: guard.location.x, y: guard.location.y + 1}

	return guard
}

func markObstacleDone(field [][]string, obstacle coordinate) [][]string {
	field[obstacle.y][obstacle.x] = "O"

	return field
}

func findNextObstacleOfGuard(field [][]string, guard guard) coordinate {
	var x int
	var y int
	var obstacleFound bool
	if guard.direction == up || guard.direction == down {
		for columnIndex, row := range field {
			for rowIndex, char := range row {
				if rowIndex == guard.location.x {
					if char == "#" {
						obstacleFound = true
						x = rowIndex
						y = columnIndex
					}
				}
			}
		}
	}

	if guard.direction == right || guard.direction == left {
		for columnIndex, row := range field {
			for rowIndex, char := range row {
				if columnIndex == guard.location.y {
					if char == "#" {
						obstacleFound = true
						x = rowIndex
						y = columnIndex
					}
				}
			}
		}
	}

	if obstacleFound {
		return coordinate{x: x, y: y}
	} else {
		return coordinate{x: -99, y: -99}
	}
}

func moveGuardUpUntilObstacle(field [][]string, guard guard, obstacle coordinate) ([][]string, guard) {
	for {
		if guard.location.y-1 == obstacle.y {
			guard = turnGuardNinetyDegreesToRight(guard)
			break
		} else {
			field[guard.location.y][guard.location.x] = "X"
			guard = moveGuardUpOnce(guard)
		}
	}

	return field, guard
}

func locateGuardInField(field [][]string) guard {
	guard := guard{}
	for columnIndex, row := range field {
		for rowIndex, char := range row {
			if char == right || char == left || char == up || char == down {
				guard.location = coordinate{x: rowIndex, y: columnIndex}
				guard.direction = char
			}
		}
	}
	return guard
}

func getSecondHalfResult(inputFile string) int {
	var total = 0

	return total
}

func loadField(inputFile string) [][]string {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var field [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var newRow []string
		for _, char := range line {
			newRow = append(newRow, string(char))
		}
		field = append(field, newRow)
	}
	return field
}

func main() {
	inputFile := "test_input.txt"
	fmt.Println(getFirstHalfResult(inputFile))
	//fmt.Println(getSecondHalfResult(inputFile))
}
