package main

import (
	"fmt"

	"github.com/bartosz-skejcik/advent_of_code_2024/internal/helper"
)

func is_safe(nums []int) bool {
	inc := nums[1] > nums[0]
	if inc {
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			if !(diff >= 1 && diff <= 3) {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			if !(diff <= -1 && diff >= -3) {
				return false
			}
		}
		return true
	}
}

func is_really_safe(nums []int) bool {
	if is_safe(nums) {
		return true
	}

	for i := range nums {
		// is_safe(nums[:i] + nums[i+1:])
		new_nums := append([]int{}, nums[:i]...)
		new_nums = append(new_nums, nums[i+1:]...)
		if is_safe(new_nums) {
			return true
		}
	}

	return false
}

func part1() {
	list := helper.ParseFileToArrayList(" ", "input.txt")

	var safeReports = 0
	for _, line := range list {
		if is_safe(line) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func part2() {
	list := helper.ParseFileToArrayList(" ", "input.txt")

	var safeReports = 0
	for _, line := range list {
		if is_really_safe(line) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func main() {
	fmt.Println("Part 1")
	part1()
	fmt.Println("Part 2")
	part2()
}
