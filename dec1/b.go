//https://adventofcode.com/2024/day/1#part2
//command: go run dec1/b.go input.txt

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	left, right, err := inputToArray()
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0

	for _, v := range left {
		sum += countMap(right, v) * v
	}

	fmt.Println(sum)
}

func inputToArray() (left []int, right map[int]int, err error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)
	rightSlice := []int{}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}

		cols := strings.Fields(line)
		li, err := strconv.Atoi(cols[0])
		if err != nil {
			return nil, nil, err
		}
		ri, err := strconv.Atoi(cols[1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, li)
		rightSlice = append(rightSlice, ri)
	}

	defer f.Close()

	right = createMap(rightSlice)

	return left, right, nil
}

func countMap(m map[int]int, c int) int {
	v, ok := m[c]
	if ok {
		return v
	}
	return 0
}

func createMap(s []int) map[int]int {
	countMap := make(map[int]int)
	for _, v := range s {
		countMap[v]++
	}
	return countMap
}
