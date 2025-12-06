package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countZeroCrossings(position int, direction byte, distance int) int {
	if direction == 'L' {
		endPosition := position - distance
		if endPosition <= 0 && position > 0 {
			return 1 + (-endPosition)/100
		}
		return (-endPosition) / 100
	}

	endPosition := position + distance
	return endPosition / 100
}

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	position := 50
	zeroCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())

		if len(instruction) < 2 {
			continue
		}

		direction := instruction[0]
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			continue
		}

		zeroCount += countZeroCrossings(position, direction, distance)

		if direction == 'L' {
			position = (position - distance) % 100
			if position < 0 {
				position += 100
			}
		} else {
			position = (position + distance) % 100
		}
	}

	fmt.Println("Password (times dial points at 0):", zeroCount)
}
