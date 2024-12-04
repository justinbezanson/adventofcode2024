//https://adventofcode.com/2024/day/3
//command: go run dec3/main.go

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	contents := readInput()

	fmt.Println(part1(contents), part2(contents))
}

func part2(contents string) int {
	sum := 0
	proceed := true
	r := regexp.MustCompile(`mul\(\d+,\d+\)|don\'t\(\)|do\(\)`)
	matches := r.FindAllString(contents, -1)

	for _, match := range matches {
		if match == "don't()" {
			proceed = false
			continue
		}

		if match == "do()" {
			proceed = true
			continue
		}

		if proceed {
			nums := getInts(match)
			sum += (nums[0] * nums[1])
		}
	}

	return sum
}

func part1(contents string) int {
	sum := 0
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(contents, -1)

	for _, match := range matches {
		nums := getInts(match)
		sum += (nums[0] * nums[1])
	}

	return sum
}

func getInts(s string) []int {
	s = strings.ReplaceAll(s, "mul(", "")
	s = strings.ReplaceAll(s, ")", "")
	parts := strings.Split(s, ",")
	nums := []int{parseInt(parts[0]), parseInt(parts[1])}

	return nums
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func readInput() string {
	contents, err := os.ReadFile("dec3/input.txt")

	if err != nil {
		fmt.Println("Could not read input file: ", err)
	}

	return string(contents)
}
