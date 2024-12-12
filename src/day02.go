package main

import (
	"bufio"
	"strconv"
	"strings"
)

func Peek[T comparable](arr []T, index int) (T, bool) {
	var zero T
	if index < 0 || index >= len(arr) {
		return zero, false
	}
	return arr[index], true
}

func abs(x int) (int, bool) {
	if x < 0 {
		return -x, true
	}
	return x, false
}

func getSafetyReport(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	isIncreasing := true
	safeReports := map[int][]string{}
	numberSafeReports := 0
	for lineNumber, line := range lines {

		reports := strings.Split(line, " ")

		for i, report := range reports {
			nextIndex := i + 1
			nextValue, ok := Peek(reports, nextIndex)
			if !ok {
				continue
			}

			reportInt, _ := strconv.Atoi(report)
			newNextValue, _ := strconv.Atoi(nextValue)

			// Check for equal values first
			if reportInt == newNextValue {
				safeReports[lineNumber] = append(safeReports[lineNumber], "no")
				continue
			}

			if i == 0 && nextIndex == 1 {
				value := reportInt - newNextValue
				if value < 0 {
					isIncreasing = true
				} else {
					isIncreasing = false
				}
				absValue, _ := abs(value)
				if (isIncreasing && value < 0) || (!isIncreasing && value > 0) {
					if absValue >= 1 && absValue <= 3 {
						safeReports[lineNumber] = append(safeReports[lineNumber], "yes")
					} else {
						safeReports[lineNumber] = append(safeReports[lineNumber], "no")

					}
				}
			}

			value := reportInt - newNextValue
			absValue, ok := abs(value)
			if isIncreasing != ok {
				safeReports[lineNumber] = append(safeReports[lineNumber], "no")
				continue
			}
			if (isIncreasing && value < 0) || (!isIncreasing && value > 0) {
				if absValue >= 1 && absValue <= 3 {
					safeReports[lineNumber] = append(safeReports[lineNumber], "yes")
				} else {
					safeReports[lineNumber] = append(safeReports[lineNumber], "no")

				}
			}
		}
		allYes := true
		for _, s := range safeReports[lineNumber] {
			if s != "yes" {
				allYes = false
				break
			}
		}
		if allYes {
			numberSafeReports++
		}
	}
	return numberSafeReports
}
