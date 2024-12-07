//https://adventofcode.com/2024/day/5
//command: go run dec6/main.go

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readInput()

	dirs := map[string][]int{
		"UP":    {-1, 0},
		"RIGHT": {0, 1},
		"DOWN":  {1, 0},
		"LEFT":  {0, -1},
	}

	fmt.Println(part1(grid, dirs), part2(grid, dirs))
}

func part1(grid [][]string, dirs map[string][]int) int {
	moves := map[string]bool{}
	pos := findStart(grid)
	newPos := []int{-1, -1}
	dir := "UP"

	for {
		moves[posToKey(pos, "")] = true

		newPos, dir = move(grid, pos, dirs, dir)

		if newPos[0] == pos[0] && newPos[1] == pos[1] {
			break
		}

		pos = newPos
	}

	return len(moves)
}

func part2(grid [][]string, dirs map[string][]int) int {
	loops := 0
	start := findStart(grid)

	for row, rv := range grid {
		for col, _ := range rv {
			if grid[row][col] == "^" || grid[row][col] == "#" {
				continue
			}

			pos := start
			newPos := []int{-1, -1}
			dir := "UP"
			newGrid := copyGrid(grid)
			newGrid[row][col] = "0"
			moves := map[string]bool{}

			for {
				key := posToKey(pos, dir)
				_, ok := moves[key]

				if ok {
					//fmt.Println(moves)
					//fmt.Println("LOOP", pos, dir)
					loops++
					break
				}

				moves[key] = true

				newPos, dir = move(newGrid, pos, dirs, dir)

				if newPos[0] == pos[0] && newPos[1] == pos[1] {
					//fmt.Println("END", pos, dir)
					break
				}

				pos = newPos
			}
		}
	}

	return loops
}

func copyGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))

	for i, row := range grid {
		newGrid[i] = make([]string, len(row))
		copy(newGrid[i], row)
	}

	return newGrid
}

func posToKey(pos []int, dir string) string {
	p1 := strconv.Itoa(pos[0])
	p2 := strconv.Itoa(pos[1])
	var buffer bytes.Buffer
	buffer.WriteString(p1)
	buffer.WriteString(":")
	buffer.WriteString(p2)
	buffer.WriteString(":")
	buffer.WriteString(dir)
	return buffer.String()
}

func move(grid [][]string, pos []int, dirs map[string][]int, dir string) ([]int, string) {
	newPos := []int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}

	if !inBounds(grid, newPos) {
		return pos, dir
	}

	c := grid[newPos[0]][newPos[1]]

	if c == "#" || c == "0" {
		dir = nextdir(dir)
		newPos = []int{pos[0] + dirs[dir][0], pos[1] + dirs[dir][1]}
	}

	return newPos, dir
}

func nextdir(dir string) string {
	switch dir {
	case "UP":
		return "RIGHT"
	case "RIGHT":
		return "DOWN"
	case "DOWN":
		return "LEFT"
	case "LEFT":
		return "UP"
	}
	return ""
}

func findStart(grid [][]string) []int {
	for i, row := range grid {
		for j, col := range row {
			if col == "^" {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}

func inBounds(grid [][]string, pos []int) bool {
	rows := len(grid)
	cols := len(grid[0]) - 1
	return pos[0] >= 0 && pos[0] < rows && pos[1] >= 0 && pos[1] < cols
}

func readInput() (grid [][]string) {
	f, err := os.Open("dec6/input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		grid = append(grid, split(line))
	}

	return grid
}

func split(s string) []string {
	cols := []string{}

	for _, c := range s {
		cols = append(cols, strings.TrimSpace(strings.TrimSuffix(string(c), "\n")))
	}

	return cols
}
