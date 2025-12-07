package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	rows := len(grid)
	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalRemoved := 0

	for {
		var toRemove [][2]int

		for r := range rows {
			for c := range cols {
				if grid[r][c] != '@' {
					continue
				}

				adjacentRolls := 0
				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
						adjacentRolls++
					}
				}

				if adjacentRolls < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
		}
		totalRemoved += len(toRemove)
	}

	fmt.Printf("Total rolls removed: %d\n", totalRemoved)
}
