package helper

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Left  string
	Right string
}

func ParseFileToArrayList(separator string, pathToList string) [][]int {
	bytes, err := os.ReadFile(pathToList)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	var result [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, separator)

		// convert parts to int
		var intParts []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)

			intParts = append(intParts, num)
		}

		// append the list to result
		// result is a list of lists. we dont want any
		result = append(result, intParts)
	}

	return result
}

func GetListFromFile(separator string, pathToList string) []Pair {
	bytes, err := os.ReadFile(pathToList)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	var result []Pair

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, separator)

		if len(parts) != 2 {
			fmt.Println("Unexpected line:", line)
			continue
		}

		result = append(result, Pair{Left: parts[0], Right: parts[1]})
	}

	return result
}
