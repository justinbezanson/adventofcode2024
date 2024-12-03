//https://adventofcode.com/2024/day/2
//command: go run dec2/main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getReports()

	answer1 := part1(reports)
	answer2 := part2(reports)

	fmt.Println(answer1, answer2)
}

func part1(reports [][]string) int {
	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		}
	}

	return safe
}

func part2(reports [][]string) int {
	safe := 0

	for _, report := range reports {
		if isSafe(report) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				dampened := append([]string(nil), report[:i]...)
				dampened = append(dampened, report[i+1:]...)
				if isSafe(dampened) {
					safe++
					break
				}
			}
		}
	}

	return safe
}

func isSafe(report []string) bool {
	safe := true
	dir := 0
	diff := 0

	for i := 1; i < len(report); i++ {
		prev, _ := strconv.Atoi(report[i-1])
		curr, _ := strconv.Atoi(report[i])

		if dir == 0 {
			if curr > prev {
				dir = 1
			} else {
				dir = -1
			}
		}

		if dir == 1 {
			diff = curr - prev
		} else {
			diff = prev - curr
		}

		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}

	return safe
}

func getReports() [][]string {
	input := readInputToLines()
	reports := [][]string{}

	for _, line := range input {
		report := strings.Split(line, " ")
		reports = append(reports, report)
	}

	return reports
}

func readInputToLines() (lines []string) {
	f, err := os.Open("dec2/input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, strings.TrimSuffix(line, "\n"))
	}

	return lines
}
