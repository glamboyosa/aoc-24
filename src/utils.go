package main

import (
	"os"
	"strings"
)

func readInput(aocDay string) string {
	data, err := os.ReadFile("./inputs/" + aocDay + ".txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
