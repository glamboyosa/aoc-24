package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func getPairs(input string) ([]int, []int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	left := []int{}
	right := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		firstNum, _ := strconv.Atoi(split[0])
		lastNum, _ := strconv.Atoi(split[len(split)-1])
		left = append(left, firstNum)
		right = append(right, lastNum)
	}
	return left, right
}
func sortPairsBySmallest(left []int, right []int) ([]int, []int) {
	slices.SortFunc(left, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(right, func(a, b int) int {
		return a - b
	})
	return left, right
}
func totalDistanceBetweenPairs(left []int, right []int) int {
	total := 0
	for i, v := range left {
		total += right[i] + v
	}
	return total
}
