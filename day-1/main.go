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

func separateNumbersIntoSlices(scanner bufio.Scanner) ([]int, []int) {
	var leftList, rightList []int

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

	return leftList, rightList
}

func partOneSolution(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	var result int
	for i := range leftList {
		result += abs(leftList[i] - rightList[i])
	}

	return result
}

func partTwoSolution(leftList, rightList []int) int {
	rightListMap := make(map[int]int)

	for i := range rightList {
		_, ok := rightListMap[rightList[i]]
		if ok {
			rightListMap[rightList[i]] += 1
			continue
		}

		rightListMap[rightList[i]] = 1
	}

	var result int
	for i := range leftList {
		occurence, ok := rightListMap[leftList[i]]
		if ok {
			result += leftList[i] * occurence
		}
	}

	return result
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening a file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	leftList, rightList := separateNumbersIntoSlices(*scanner)

	firstPart := partOneSolution(leftList, rightList)
	fmt.Println("First part solution: ", firstPart)

	secondPart := partTwoSolution(leftList, rightList)
	fmt.Println("Second part solution: ", secondPart)
}
