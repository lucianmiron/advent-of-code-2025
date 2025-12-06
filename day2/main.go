package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeatedPattern(number int64) bool {
	str := strconv.FormatInt(number, 10)
	length := len(str)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := str[:patternLen]
		isRepeating := true

		for i := patternLen; i < length; i += patternLen {
			if str[i:i+patternLen] != pattern {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}

func findInvalidIDsInRange(start, end int64) []int64 {
	var invalidIDs []int64

	for num := start; num <= end; num++ {
		if isRepeatedPattern(num) {
			invalidIDs = append(invalidIDs, num)
		}
	}

	return invalidIDs
}

func parseRanges(input string) [][2]int64 {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")
	ranges := make([][2]int64, 0, len(parts))

	for _, part := range parts {
		rangeParts := strings.Split(strings.TrimSpace(part), "-")
		if len(rangeParts) != 2 {
			continue
		}

		start, err1 := strconv.ParseInt(rangeParts[0], 10, 64)
		end, err2 := strconv.ParseInt(rangeParts[1], 10, 64)

		if err1 != nil || err2 != nil {
			continue
		}

		ranges = append(ranges, [2]int64{start, end})
	}

	return ranges
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	ranges := parseRanges(string(data))

	var totalSum int64
	var allInvalidIDs []int64

	for _, r := range ranges {
		invalidIDs := findInvalidIDsInRange(r[0], r[1])
		allInvalidIDs = append(allInvalidIDs, invalidIDs...)
		for _, id := range invalidIDs {
			totalSum += id
		}
	}

	fmt.Printf("Found %d invalid IDs\n", len(allInvalidIDs))
	fmt.Printf("Sum of all invalid IDs: %d\n", totalSum)
}
