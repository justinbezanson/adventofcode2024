//https://adventofcode.com/2024/day/5
//command: go run dec5/main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readInput()

	fmt.Println(part1(rules, updates), part2(rules, updates))
}

func part1(rules [][]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		valid := true
	nextUpdate:
		for i, v := range update {
			for _, rule := range rules {
				ruleIndex := slices.Index(update, rule[1])
				if rule[0] == v && ruleIndex > -1 && ruleIndex < i {
					valid = false
					break nextUpdate
				}
			}
		}

		if valid {
			middleIndex := (len(update) - 1) / 2
			sum += update[middleIndex]
		}
	}

	return sum
}

func part2(rules [][]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		valid := true
	nextUpdate:
		for i, v := range update {
			for _, rule := range rules {
				ruleIndex := slices.Index(update, rule[1])
				if rule[0] == v && ruleIndex > -1 && ruleIndex < i {
					valid = false
					break nextUpdate
				}
			}
		}

		if !valid {
			newUpdate := reorderUpdate(update, rules)
			middleIndex := (len(newUpdate) - 1) / 2
			sum += newUpdate[middleIndex]
		}
	}

	return sum
}

func isValid(update []int, rules [][]int) bool {
	for i, v := range update {
		for _, rule := range rules {
			ruleIndex := slices.Index(update, rule[1])
			if rule[0] == v && ruleIndex > -1 && ruleIndex < i {
				return false
			}
		}
	}
	return true
}

func reorderUpdate(update []int, rules [][]int) []int {
	for i, v := range update {
		for _, rule := range rules {
			ruleIndex := slices.Index(update, rule[1])
			if rule[0] == v && ruleIndex > -1 && ruleIndex < i {
				update[i], update[ruleIndex] = update[ruleIndex], update[i]
			}
		}
	}

	if !isValid(update, rules) {
		return reorderUpdate(update, rules)
	}

	return update
}

func readInput() (rules [][]int, updates [][]int) {
	f, err := os.Open("dec5/input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Index(line, "|") > -1 {
			rules = append(rules, split(line, "|"))
		} else if strings.Index(line, ",") > -1 {
			updates = append(updates, split(line, ","))
		}
	}

	return rules, updates
}

func split(s string, sep string) []int {
	slice := strings.Split(s, sep)
	ints := []int{}

	for _, v := range slice {
		i, err := strconv.Atoi(strings.TrimSuffix(v, "\n"))
		if err != nil {
			return nil
		}

		ints = append(ints, i)
	}

	return ints
}
