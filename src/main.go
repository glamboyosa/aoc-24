package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	dayFlag := flag.String("day", "", "which day to run (e.g., day01)")
	flag.Parse()
	fmt.Println("Day", *dayFlag)

	switch *dayFlag {
	case "day01":
		fmt.Println("Day 01")
		input := readInput("day01")
		left, right := getPairs(input)
		left, right = sortPairsBySmallest(left, right)
		totalDistance := totalDistanceBetweenPairs(left, right)
		fmt.Println("Total distance between pairs", totalDistance)

	case "day02":
		fmt.Println("Day 02")
		input := readInput("day02")
		numberSafeReports := getSafetyReport(input)
		// Add your day 2 logic here
		fmt.Println("Number of safe reports", numberSafeReports)

	default:
		if strings.TrimSpace(*dayFlag) == "" {
			fmt.Println("Please specify a day with -day flag")
		} else {
			fmt.Printf("Unknown day: %s\n", *dayFlag)
		}
	}
}
