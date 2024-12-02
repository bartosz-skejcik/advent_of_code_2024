package main

import (
	"fmt"
	"math"

	"strconv"

	"github.com/bartosz-skejcik/advent_of_code_2024/internal/helper"
)

func sortFromLowToHigh(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
	return list
}

func getLeftAndRightList() ([]int, []int) {
	// var left_list = []int{3, 4, 2, 1, 3, 3}
	// var right_list = []int{4, 3, 5, 3, 9, 3}
	var left_list = []int{}
	var right_list = []int{}

	result := helper.GetListFromFile("   ", "input.txt")

	for _, pair := range result {
		leftItem, _ := strconv.Atoi(pair.Left)
		rightItem, _ := strconv.Atoi(pair.Right)

		left_list = append(left_list, leftItem)
		right_list = append(right_list, rightItem)
	}

	return left_list, right_list
}

func numOfOccurrences(item int, list []int) int {
	var sum int
	for _, val := range list {
		if val == item {
			sum += 1
		}
	}

	return sum * item
}

func calcSimilarityScore(left_side []int, right_side []int) int {
	var score int

	for _, val := range left_side {
		occurrances := numOfOccurrences(val, right_side)

		score += occurrances
	}

	return score
}

func main() {
	left_list, right_list := getLeftAndRightList()

	left_list = sortFromLowToHigh(left_list)
	right_list = sortFromLowToHigh(right_list)

	var differences []int

	for i, value := range left_list {
		var right_element = right_list[i]

		differences = append(differences, int(math.Abs(float64(value-right_element))))
	}

	var sum int
	for _, value := range differences {
		sum += value
	}

	score := calcSimilarityScore(left_list, right_list)
	fmt.Printf("Score is: %d\n", score)

	fmt.Println("Sum of differences:", sum)
}
