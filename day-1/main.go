package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening a file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		nums := strings.Fields(line)
		if len(nums) < 2 {
			fmt.Fprintf(os.Stderr, "Invalid line format: %s\n", line)
		}

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing number: %v\n", err)
			os.Exit(1)
		}

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing number: %v\n", err)
			os.Exit(1)
		}

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	var result int
	for i := range leftList {
		result += abs(leftList[i] - rightList[i])
	}

	fmt.Println("Result: ", result)

}
