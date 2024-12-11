package main

import (
	"bufio"
	"strconv"
	"strings"
)

func getSafetyReport(input string) ([]int, []int) {
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
