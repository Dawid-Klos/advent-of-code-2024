package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertStringSliceToInt(s []string) []int {
	var nums []int
	for i := range s {
		num, err := strconv.Atoi(s[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting string to an int: %s \n", err)
		}
		nums = append(nums, num)
	}

	return nums
}

func isLevelSafe(level []int) bool {
	isIncreasing := level[0] < level[1]

	for i := range level {
		if i+1 >= len(level) {
			break
		}
		if level[i] == level[i+1] {
			return false
		}
		if isIncreasing && (level[i] > level[i+1] || level[i+1]-level[i] > 3) {
			return false
		}
		if !isIncreasing && (level[i] < level[i+1] || level[i]-level[i+1] > 3) {
			return false
		}
	}

	return true
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening a file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var result int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		numsAsString := strings.Fields(line)

		nums := convertStringSliceToInt(numsAsString)
		if isLevelSafe(nums) {
			result += 1
		}
	}

	fmt.Println(result)
}
