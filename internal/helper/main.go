package helper

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	Left  string
	Right string
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
