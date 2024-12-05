//https://adventofcode.com/2024/day/4
//command: go run dec4/main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	grid := readInputToGrid()

	fmt.Println(part1(grid), part2(grid))
}

func part1(grid [][]string) int {
	sum := 0
	word := "XMAS"
	rows := len(grid)
	cols := len(grid[0])

	dirMap := [][]int{
		{0, 1},   // (Horizontal forward)
		{0, -1},  // (Horizontal backward)
		{1, 0},   // Vertical (down)
		{-1, 0},  // Vertical (up)
		{1, 1},   // Diagonal (down-right)
		{-1, -1}, // Diagonal (up-left)
		{1, -1},  // Diagonal (down-left)
		{-1, 1},  // Diagonal (up-right)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range dirMap {
				dirRow := dir[0]
				dirCol := dir[1]

				if isMatch(grid, word, row, col, dirRow, dirCol) {
					sum++
				}
			}
		}
	}

	return sum
}

func part2(grid [][]string) int {
	sum := 0
	rows := len(grid)
	cols := len(grid[0])
	matches := []string{"SAMSAM", "MASMAS", "MASSAM", "SAMMAS"}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			x := grid[i-1][j-1] + grid[i][j] + grid[i+1][j+1] + grid[i-1][j+1] + grid[i][j] + grid[i+1][j-1]
			fmt.Println(x)
			if slices.Contains(matches, x) {
				sum++
			}
		}
	}

	return sum
}

func isMatch(grid [][]string, word string, startRow int, startCol int, dirRow int, dirCol int) bool {
	for i, c := range word {
		newRow := startRow + i*dirRow
		newCol := startCol + i*dirCol

		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) || grid[newRow][newCol] != string(c) {
			return false
		}
	}

	return true
}

func readInputToGrid() (grid [][]string) {
	f, err := os.Open("dec4/input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSuffix(line, "\n")
		cols := []string{}

		for _, c := range line {
			cols = append(cols, string(c))
		}

		grid = append(grid, cols)
	}

	return grid
}
