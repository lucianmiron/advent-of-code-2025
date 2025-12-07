package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxJoltage(bank string) int {
	maxVal := 0
	length := len(bank)

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			firstDigit := int(bank[i] - '0')
			secondDigit := int(bank[j] - '0')
			joltage := firstDigit*10 + secondDigit

			if joltage > maxVal {
				maxVal = joltage
			}
		}
	}

	return maxVal
}

func maxJoltage12(bank string) int64 {
	numDigits := 12
	length := len(bank)
	result := int64(0)
	startPos := 0

	for digit := 0; digit < numDigits; digit++ {
		digitsNeeded := numDigits - digit - 1
		endPos := length - digitsNeeded

		bestDigit := byte('0')
		bestPos := startPos
		for pos := startPos; pos < endPos; pos++ {
			if bank[pos] > bestDigit {
				bestDigit = bank[pos]
				bestPos = pos
			}
		}

		result = result*10 + int64(bestDigit-'0')
		startPos = bestPos + 1
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0
	totalSum12 := int64(0)

	for scanner.Scan() {
		bank := scanner.Text()
		if bank == "" {
			continue
		}
		joltage := maxJoltage(bank)
		totalSum += joltage
		joltage12 := maxJoltage12(bank)
		totalSum12 += joltage12
	}

	fmt.Printf("Part 1 - Total joltage: %d\n", totalSum)
	fmt.Printf("Part 2 - Total joltage: %d\n", totalSum12)
}
