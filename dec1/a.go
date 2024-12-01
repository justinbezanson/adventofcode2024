//https://adventofcode.com/2024/day/1
//command: go run dec1/a.go input.txt

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	distance := 0
	linesLeft, linesRight, err := inputToArray()

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range linesLeft {
		distance += absdiff(linesLeft[i], linesRight[i])
	}

	fmt.Println(distance)
}

func absdiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func inputToArray() (linesLeft []int, linesRight []int, err error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return
	}

	r := bufio.NewReader(f)

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

		linesLeft = append(linesLeft, li)
		linesRight = append(linesRight, ri)
	}

	sort.Slice(linesLeft, func(i, j int) bool {
		return linesLeft[i] < linesLeft[j]
	})

	sort.Slice(linesRight, func(i, j int) bool {
		return linesRight[i] < linesRight[j]
	})

	defer f.Close()

	return linesLeft, linesRight, nil
}
